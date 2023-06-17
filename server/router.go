package main

import (
	"net/http"
	"server/action"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Route struct {
	Method string
	Path   string
	Action echo.HandlerFunc
}

var GetBlogPostRouteMap = []*Route{
	{
		Method: http.MethodGet,
		Path:   "/v1/web/posts",
		Action: action.Post.GetAllPosts,
	},
	{
		Method: http.MethodGet,
		Path:   "/v1/web/single/:id",
		Action: action.Post.GetSinglePost,
	},
	{
		Method: http.MethodGet,
		Path:   "/v1/web/recent_posts",
		Action: action.Post.GetRecentPosts,
	},
}

type Router struct {
	Server *echo.Echo
}

func (m *Router) Init() {
	m.Server = echo.New()

	m.addRoute(
		GetBlogPostRouteMap,
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
