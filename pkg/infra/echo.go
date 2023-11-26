package infra

import (
	"github.com/inahym196/accountant/pkg/infra/router"
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	echo "github.com/labstack/echo/v4"
)

type server struct {
	AccountItemRouter router.Router
}

func NewServer(c controller.AccountItemController) server {
	return server{router.NewAccountItemRouter(c)}
}

func (s server) Run(addr string) {
	e := echo.New()
	accountItem := e.Group("/account_item")
	accountItem.GET("", router.RouterFunc(s.AccountItemRouter.Get))
	accountItem.POST("", router.RouterFunc(s.AccountItemRouter.Post))
	e.Start(addr)
}
