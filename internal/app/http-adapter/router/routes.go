package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mkorobovv/aichat/internal/app/http-adapter/controller"
)

func (r *Router) AppendRoutes(ctr *controller.Controller) {
	apiV1Subrouter := r.router.PathPrefix(apiV1Prefix).Subrouter()

	routes := []Route{
		{
			Name:    "/send",
			Path:    "/send",
			Method:  http.MethodPost,
			Handler: http.HandlerFunc(ctr.SendMessage),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}

func (r *Router) appendRoutesToRouter(subrouter *mux.Router, routes []Route) {
	for _, route := range routes {
		subrouter.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}
