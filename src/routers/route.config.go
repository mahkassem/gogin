package routers

var routes = map[string]Route{
	"user": {
		Path: "/users",
		Middlewares: []string{
			"AuthenticationMiddleware",
			"TestMiddleware",
		},
		Routes: []Route{
			{
				Path:            "/",
				Handler:         "GetAllUsers",
				SkipMiddlewares: []string{"*"},
				Routes: []Route{
					{
						Path:    "test/:id",
						Method:  "GET",
						Handler: "GetUserById",
						Middlewares: []string{
							"Test2Middleware",
						},
					},
				},
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
		Path: "/ping",
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
	Path            string
	Method          string
	Handler         string
	Middlewares     []string
	SkipMiddlewares []string
	Routes          []Route
}
