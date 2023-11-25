package infra

import (
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	echo "github.com/labstack/echo/v4"
)

type AccountItemRouter interface {
	Get(echo.Context) error
	Post(echo.Context) error
}

type accountItemRouter struct {
	c controller.AccountItemController
}

func NewAccountItemRouter(c controller.AccountItemController) AccountItemRouter {
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
