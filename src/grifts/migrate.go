package grifts

import (
	"database/sql"
	"github.com/eldad87/go-boilerplate/src/config"
	sqlLogger "github.com/eldad87/go-boilerplate/src/pkg/go-sql-driver/logger"
	promZap "github.com/eldad87/go-boilerplate/src/pkg/uber/zap"
	databaseDriver "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/packr"
	"github.com/luna-duclos/instrumentedsql"
	instrumentedsqlOpenTracing "github.com/luna-duclos/instrumentedsql/opentracing"
	. "github.com/markbates/grift/grift"
	"github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var _ = Namespace("db", func() {
	Desc("migrate", "Migrates the databases")
	Set("migrate", func(c *Context) error {
		/*
		 * PreRequisite: Config
		 * **************************** */
		conf, err := config.GetConfig(os.Getenv("BUILD_ENV"), nil)
		if err != nil {
			panic(err) // Nothing we can do
		}

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

		/*
		 * PreRequisite: DataBase
		 * **************************** */
		// Logger
		databaseDriver.SetLogger(sqlLogger.NewLogger(logger))
		// Opentracing https://ceshihao.github.io/2018/11/29/tracing-db-operations/
		sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(databaseDriver.MySQLDriver{}, instrumentedsql.WithTracer(instrumentedsqlOpenTracing.NewTracer(false))))
		db, err := sql.Open("instrumented-mysql", conf.GetString("database.dsn"))
		if err != nil {
			logger.Sugar().Fatal("Database failed to listen: %v. Due to error: %v", conf.GetString("database.dsn"), err)
		}

		if err := db.Ping(); err != nil {
			logger.Sugar().Errorf("Database failed to Ping: %v. Due to error: %v", conf.GetString("database.dsn"), err)
		}

		migrations := &migrate.PackrMigrationSource{
			Box: packr.NewBox("../../src/migration"),
		}
		appliedMigrations, err := migrate.Exec(db, conf.GetString("database.driver"), migrations, migrate.Up)
		if err != nil {
			logger.Error("Error applying migration: ", zap.Error(err))
		}
		logger.Debug("Applied migrations: ", zap.Int("total", appliedMigrations))

		return nil
	})
})
