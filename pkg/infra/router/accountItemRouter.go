package router

import (
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

type accountItemRouter struct {
	c controller.AccountItemController
}

func NewAccountItemRouter(c controller.AccountItemController) Router {
	return accountItemRouter{c}
}

func (a accountItemRouter) Get(ctx context) error {
	a.c.Get(ctx)
	return nil
}

func (a accountItemRouter) Post(ctx context) error {
	a.c.Save(ctx)
	return nil
}
