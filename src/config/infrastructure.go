package config

type Route struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Handler string `json:"handler"`
}

type RouteConfig struct {
	Path        string           `json:"path"`
	Description string           `json:"description"`
	Routes      map[string]Route `json:"routes"`
}

type Config struct {
	Routes map[string]RouteConfig `json:"routes"`
}

type Router interface {
	SetupRouter()
}

type Response struct {
	Message string
	Data    interface{}
	Error   *Error
}

type Error struct {
	error
	Message  string
	Code     int
	MetaData interface{}
}
