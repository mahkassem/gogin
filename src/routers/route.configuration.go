package routers

var routes = map[string]RouteConfig{
	"user": {
		Path:        "/users",
		Description: "User routes",
		Routes: []Route{
			{
				Path:    "/",
				Method:  "GET",
				Handler: "GetAllUsers",
			},
			{
				Path:    "/:id",
				Method:  "GET",
				Handler: "GetUserById",
			},
			{
				Path:    "/",
				Method:  "POST",
				Handler: "CreateUser",
			},
			{
				Path:    "/:id",
				Method:  "PUT",
				Handler: "UpdateUser",
			},
			{
				Path:    "/:id",
				Method:  "DELETE",
				Handler: "DeleteUser",
			},
		},
	},
	"ping": {
		Path:        "/ping",
		Description: "Ping routes",
		Routes: []Route{
			{
				Path:    "/",
				Method:  "GET",
				Handler: "Test",
			},
		},
	},
}

type Route struct {
	Path    string `json:"path"`
	Method  string `json:"method"`
	Handler string `json:"handler"`
}

type RouteConfig struct {
	Path        string  `json:"path"`
	Description string  `json:"description"`
	Routes      []Route `json:"routes"`
}
