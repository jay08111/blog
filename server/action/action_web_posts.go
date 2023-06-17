package action

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (m *PostAction) GetAllPosts(ctx echo.Context) error {
	logrus.Print("hey")

	return nil
}

func (m *PostAction) GetRecentPosts(ctx echo.Context) error {
	logrus.Print("hey")
	return nil
}

func (m *PostAction) GetSinglePost(ctx echo.Context) error {
	logrus.Print("hey")
	return nil
}
