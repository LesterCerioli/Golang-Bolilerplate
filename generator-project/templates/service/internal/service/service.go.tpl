package service

type ExampleService struct{}

func NewExampleService {
    return &ExampleService{}
}

func (s *ExampleService) DoSomething() string {
    return "Service is doing somethi8ng!"
}