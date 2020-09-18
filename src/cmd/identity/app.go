package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/TheZeroSlave/zapsentry"
	metricCollector "github.com/afex/hystrix-go/hystrix/metric_collector"
	"github.com/eldad87/go-boilerplate/src/config"
	promZap "github.com/eldad87/go-boilerplate/src/pkg/uber/zap"

	sqlLogger "github.com/eldad87/go-boilerplate/src/pkg/go-sql-driver/logger"
	databaseDriver "github.com/go-sql-driver/mysql"

	sqlmwInterceptor "github.com/eldad87/go-boilerplate/src/pkg/ngrok/sqlmw"
	"github.com/ngrok/sqlmw"

	"github.com/heptiolabs/healthcheck"
	"github.com/ibm-developer/generator-ibm-core-golang-gin/generators/app/templates/plugins"
	jaegerZap "github.com/jaegertracing/jaeger-client-go/log/zap"
	jaegerprom "github.com/jaegertracing/jaeger-lib/metrics/prometheus"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	muxMiddleware "github.com/eldad87/go-boilerplate/src/pkg/mux/middleware"
	mux "github.com/gorilla/mux"
)

func main() {
	/*
	 * PreRequisite: Config
	 * **************************** */
	conf, err := config.GetConfig(os.Getenv("BUILD_ENV"), nil)
	if err != nil {
		panic(err) // Nothing we can do
	}
	if conf.GetString("environment") != "production" {
		conf.Debug()
	}

	/*
	 * PreRequisite: Prometheus
	 * **************************** */
	collector := plugins.InitializePrometheusCollector(plugins.PrometheusCollectorConfig{
		Namespace: conf.GetString("app.name"),
	})
	http.Handle(conf.GetString("prometheus.route"), promhttp.Handler())

	/*
	 * PreRequisite: Hystrix
	 * **************************** */
	// Expose CB Prometheus metrics
	metricCollector.Registry.Register(collector.NewPrometheusCollector)

	/*
	 * PreRequisite: Health Check + Expose status Prometheus metrics gauge
	 * **************************** */
	healthChecker := healthcheck.NewMetricsHandler(prometheus.DefaultRegisterer, "health_check")
	healthChecker.AddLivenessCheck("Goroutine Threshold", healthcheck.GoroutineCountCheck(conf.GetInt("health_check.goroutine_threshold")))

	// Expose to HTTP
	http.HandleFunc(conf.GetString("health_check.route.group")+conf.GetString("health_check.route.live"), healthChecker.LiveEndpoint)
	http.HandleFunc(conf.GetString("health_check.route.group")+conf.GetString("health_check.route.ready"), healthChecker.ReadyEndpoint)

	/*
	 * PreRequisite: Logger
	 * **************************** */
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level.UnmarshalText([]byte(conf.GetString("log.level")))
	zapConfig.Development = conf.GetString("environment") != "production"
	// Expose log level Prometheus metrics
	hook := promZap.MustNewPrometheusHook([]zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel,
		zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel, zapcore.DebugLevel})
	logger, _ := zapConfig.Build(zap.Hooks(hook))

	// Sentry
	if conf.GetString("sentry.dsn") != "" {
		atom := zap.NewAtomicLevel()
		err := atom.UnmarshalText([]byte(conf.GetString("sentry.log_level")))
		if err != nil {
			logger.Fatal("Failed to parse Zap-Sentry log level", zap.String("sentry.log_level", conf.GetString("sentry.log_level")))
		}

		cfg := zapsentry.Configuration{
			Level: atom.Level(), //when to send message to sentry
			Tags: map[string]string{
				"component": conf.GetString("app.name"),
			},
		}
		core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(conf.GetString("sentry.dsn")))
		//in case of err it will return noop core. so we can safely attach it
		if err != nil {
			logger.Fatal("failed to init sentry / zap")
		}
		logger = zapsentry.AttachCoreToLogger(core, logger)
	}
	defer logger.Sync()

	/*
	 * PreRequisite: Jaeger
	 * **************************** */
	// Add jaeger metrics and reporting to prometheus route
	logAdapt := jaegerZap.NewLogger(logger)
	factory := jaegerprom.New() // By default uses prometheus.DefaultRegisterer
	metrics := jaeger.NewMetrics(factory, map[string]string{"lib": "jaeger"})

	// Add tracing to application
	transport, err := jaeger.NewUDPTransport(conf.GetString("opentracing.jaeger.host")+":"+conf.GetString("opentracing.jaeger.port"), 0)
	if err != nil {
		healthChecker.AddReadinessCheck("jaeger", func() error { return err }) // Permanent, take us down.
		logger.Sugar().Errorf("%+v", err)
	}

	reporter := jaeger.NewCompositeReporter(
		jaeger.NewLoggingReporter(logAdapt),
		jaeger.NewRemoteReporter(transport,
			jaeger.ReporterOptions.Metrics(metrics),
			jaeger.ReporterOptions.Logger(logAdapt),
		),
	)
	defer reporter.Close()

	sampler := jaeger.NewConstSampler(true)
	tracer, closer := jaeger.NewTracer(conf.GetString("app.name"),
		sampler,
		reporter,
		jaeger.TracerOptions.Metrics(metrics),
	)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	/*
	 * PreRequisite: DataBase
	 * **************************** */
	// Logger
	databaseDriver.SetLogger(sqlLogger.NewLogger(logger))
	// Tracer
	mysqlInterceptor := sqlmwInterceptor.Interceptor{Tracer: tracer}
	sql.Register("instrumented-mysql", sqlmw.Driver(databaseDriver.MySQLDriver{}, mysqlInterceptor))
	db, err := sql.Open("instrumented-mysql", conf.GetString("database.dsn"))
	if err != nil {
		logger.Sugar().Fatal("Database failed to listen: %v. Due to error: %v", conf.GetString("database.dsn"), err)
	}

	if err := db.Ping(); err != nil {
		logger.Sugar().Errorf("Database failed to Ping: %v. Due to error: %v", conf.GetString("database.dsn"), err)
	}
	// Our app is not ready if we can't connect to our database (`var db *sql.DB`) in <1s.
	healthChecker.AddReadinessCheck(conf.GetString("database.driver"), healthcheck.DatabasePingCheck(db, 1*time.Second))

	/*
	 * HTTP
	 * **************************** */
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := mux.NewRouter()
	r.Use(muxMiddleware.Prometheus, muxMiddleware.ContextReqId, muxMiddleware.Opentracing)

	// https://github.com/grpc-ecosystem/grpc-gateway/issues/348
	muxHandlerFunc := nethttp.MiddlewareFunc(
		tracer,
		r.ServeHTTP,
		nethttp.OperationNameFunc(func(r *http.Request) string {
			return fmt.Sprintf("HTTP-gRPC %s %s", r.Method, r.URL.String())
		}),
	)

	http.HandleFunc(conf.GetString("app.identity.http_route_prefix.http_route_prefix")+"/", muxHandlerFunc)

	/*
	 * Start listening for incoming HTTP requests
	 * **************************** */
	logger.Info("Starting..")
	http.ListenAndServe(":"+conf.GetString("app.port"), nil)
}
