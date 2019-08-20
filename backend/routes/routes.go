package routes

import (
	"net/http"

	"../auth"
	"../controllers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

//format: method, pattern, controller(handler), need middleware
func init() {
	register("POST", "/user/register", controllers.Register, nil)
	register("POST", "/user/login", controllers.Login, nil)
	register("GET", "/user/info", controllers.UserInfo, auth.TokenMiddleware)
	register("POST", "/user/logout", controllers.Logout, auth.TokenMiddleware)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).Path(route.Pattern)

		if route.Middleware != nil {
			//If the route need middleware handle
			r.Handler(route.Middleware(route.Handler))
		} else {
			//normal route
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
