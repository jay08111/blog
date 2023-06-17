package action

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (m *WebAction) GetAllPosts(ctx echo.Context) error {

	logrus.Print("hey")

	return nil
}

func (m *WebAction) GetRecentPosts(ctx echo.Context) error {
	logrus.Print("hey")
	return nil
}

func (m *WebAction) GetSinglePost(ctx echo.Context) error {
	logrus.Print("hey")
	return nil
}
