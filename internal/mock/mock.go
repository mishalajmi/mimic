package mock

type HTTPMethod string

var (
	Get    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

type Definition struct {
	Routes []Route `yaml:"routes"`
}

type Route struct {
	Method   HTTPMethod `yaml:"method"`
	Path     string     `yaml:"path"`
	Response Response   `yaml:"response"`
}

type Response struct {
	Status int    `yaml:"status"`
	Body   string `yaml:"body"`
}
