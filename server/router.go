package main

import (
	"net/http"
	"server/action"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Route struct {
	Method     string
	Path       string
	Action     echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

var BlogPostRouteMap = []*Route{
	{
		Method: http.MethodGet,
		Path:   "/v1/web/posts",
		Action: action.Web.GetAllPosts,
	},
	{
		Method: http.MethodGet,
		Path:   "/v1/web/single/:id",
		Action: action.Web.GetSinglePost,
	},
	{
		Method: http.MethodGet,
		Path:   "/v1/web/recent_posts",
		Action: action.Web.GetRecentPosts,
	},
}

type Router struct {
	Server *echo.Echo
}

func (m *Router) Init() {
	m.Server = echo.New()

	m.addRoute(
		BlogPostRouteMap,
	)

	m.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
	}))
}

func (m *Router) addRoute(routeMaps ...[]*Route) {
	for _, routeMap := range routeMaps {
		for i := range routeMap {
			m.Server.Add(
				routeMap[i].Method,
				routeMap[i].Path,
				routeMap[i].Action,
			)
		}
	}
}
