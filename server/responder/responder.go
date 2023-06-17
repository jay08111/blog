package responder

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type WebResponder struct {
}

var WebRes *WebResponder

type JSONResponder struct {
}

type TextResponder struct {
}

var JsonResponder *JSONResponder
var TxtResponder *TextResponder

func NewJsonResponder() *JSONResponder {
	return &JSONResponder{}
}

func NewTextResponder() *TextResponder {
	return &TextResponder{}
}

func (m *JSONResponder) Response(ctx echo.Context, code int, val interface{}) error {
	switch t := val.(type) {
	case error:
		logrus.Error(ctx, t)
		return ctx.String(code, t.Error())
	default:
		return ctx.JSON(code, t)
	}
}

func (m *TextResponder) responder(ctx echo.Context, code int, val string) error {
	return ctx.String(code, val)
}

func init() {
	JsonResponder = NewJsonResponder()
	TxtResponder = NewTextResponder()

	WebRes = &WebResponder{}
}
