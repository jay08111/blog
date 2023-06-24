package action

import (
	"net/http"
	"server/domain"

	resp "server/responder"

	"github.com/labstack/echo/v4"
)

func (m *WebAction) GetAllPosts(ctx echo.Context) error {
	getResult, err := domain.RetAllPosts(domain.MainDB)

	if err != nil {
		return resp.JsonResponder.Response(ctx, http.StatusServiceUnavailable, err)
	}

	return resp.JsonResponder.Response(ctx, http.StatusOK, getResult)
}

func (m *WebAction) GetRecentPosts(ctx echo.Context) error {
	getRecentPosts, err := domain.RetRecentPosts(ctx, domain.MainDB)

	if err != nil {
		return resp.JsonResponder.Response(ctx, http.StatusServiceUnavailable, err)
	}

	return resp.JsonResponder.Response(ctx, http.StatusOK, getRecentPosts)

}

func (m *WebAction) GetSinglePost(ctx echo.Context) error {
	id, err := getIntParam(ctx, "id")

	if err != nil {
		return resp.JsonResponder.Response(ctx, http.StatusBadRequest, err)
	}

	getSinglePost, err := domain.RetSinglePost(ctx, domain.MainDB, id)

	if err != nil {
		return resp.JsonResponder.Response(ctx, http.StatusServiceUnavailable, err)
	}

	return resp.JsonResponder.Response(ctx, http.StatusOK, getSinglePost)

}
