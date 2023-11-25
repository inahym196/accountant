package infra

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
	accountItemRouter
}

func NewServer(c controller.AccountItemController) server { return server{NewAccountItemRouter(c)} }

func (s server) Run(addr string) {
	e := echo.New()
	accountItem := e.Group("/account_item")
	accountItem.GET("", s.accountItemRouter.Get)
	accountItem.POST("", s.accountItemRouter.Post)
	e.Start(addr)
}

type accountItemRouter struct {
	c controller.AccountItemController
}

func NewAccountItemRouter(c controller.AccountItemController) accountItemRouter {
	return accountItemRouter{c}
}

func (a accountItemRouter) Get(c echo.Context) error {
	ctx := context{c}
	a.c.Get(ctx)
	return nil
}

func (a accountItemRouter) Post(c echo.Context) error {
	ctx := context{c}
	a.c.Save(ctx)
	return nil
}
