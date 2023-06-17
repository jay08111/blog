package action

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"server/service"
)

func (m *WebAction) GetAllPosts(ctx echo.Context) error {
	return service.GetAllBlogPosts(ctx)
}

func (m *WebAction) GetRecentPosts(ctx echo.Context) error {
	logrus.Print("hey")
	return nil
}

func (m *WebAction) GetSinglePost(ctx echo.Context) error {
	logrus.Print("hey")
	return nil
}
