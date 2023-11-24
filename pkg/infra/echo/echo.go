package echo

import (
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
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

type server struct {
	c controller.AccountItemController
}

func NewServer(c controller.AccountItemController) server {
	return server{c}
}

func (s server) Run(addr string) {
	e := echo.New()
	accountItem := e.Group("/account_item")
	accountItem.GET("", s.GetAccountItem)
	accountItem.POST("", s.PostAccountItem)
	e.Start(addr)
}

func (s server) GetAccountItem(c echo.Context) error {
	ctx := context{c}
	s.c.Get(ctx)
	return nil
}

func (s server) PostAccountItem(c echo.Context) error {
	ctx := context{c}
	s.c.Save(ctx)
	return nil
}
