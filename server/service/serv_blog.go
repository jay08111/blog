package service

import (
	"net/http"
	"server/domain"
	resp "server/responder"

	"github.com/labstack/echo/v4"
)

func GetAllBlogPosts(ctx echo.Context) error {
	getResult, err := domain.PostsSelectAll()
	if err != nil {
		return resp.JsonResponder.Response(ctx, http.StatusServiceUnavailable, err)
	}
	return resp.JsonResponder.Response(ctx, http.StatusOK, getResult)
}
