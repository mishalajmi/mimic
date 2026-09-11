package mock

type Definition struct {
	Name   string  `yaml:"name"`
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Name     string   `yaml:"name"`
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

type Request struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

type Response struct {
	Status  int               `yaml:"status"`
	Headers map[string]string `yaml:"headers"`
	Body    any               `yaml:"body"`
}
