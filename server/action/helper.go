package action

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func getIntParam(ctx echo.Context, param string) (int, error) {
	p := ctx.Param(param)

	if p == "" {
		return -1, nil
	}

	ret, err := strconv.Atoi(p)

	if err != nil {
		return -1, err
	}

	return ret, nil
}
