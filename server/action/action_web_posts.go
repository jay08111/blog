package action

import (
	"net/http"
	"server/domain"

	resp "server/responder"

	"github.com/labstack/echo/v4"
)

func (m *WebAction) GetAllPosts(ctx echo.Context) error {
	getResult, err := domain.PostsSelectAll()

	if err != nil {
		return resp.JsonResponder.Response(ctx, http.StatusServiceUnavailable, err)
	}

	return resp.JsonResponder.Response(ctx, http.StatusOK, getResult)
}

func (m *WebAction) GetRecentPosts(ctx echo.Context) error {
	return nil
}

func (m *WebAction) GetSinglePost(ctx echo.Context) error {
	return nil
}
