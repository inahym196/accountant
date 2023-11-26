package router

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type context struct {
	e echo.Context
}

func (ctx context) Text(text string) {
	ctx.e.String(http.StatusOK, text)
}

func (ctx context) Query() map[string][]string {
	return ctx.e.QueryParams()
}

func (ctx context) PostForm() map[string][]string {
	params, err := ctx.e.FormParams()
	if err != nil {
		return nil
	}
	return params
}

type Router interface {
	Get(ctx context) error
	Post(ctx context) error
}

func RouterFunc(f func(ctx context) error) func(echo.Context) error {
	return func(e echo.Context) error { return f(context{e}) }
}

func NewContext(c echo.Context) context {
	return context{c}
}
