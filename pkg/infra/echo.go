package infra

import (
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	echo "github.com/labstack/echo/v4"
)

type context struct{ e echo.Context }

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

func RouterFunc(f func(ctx controller.Context), err error) func(e echo.Context) error {
	return func(e echo.Context) error {
		f(context{e})
		return err
	}
}

type server struct {
	accountItem controller.AccountItemController
}

func NewServer(accountItem controller.AccountItemController) server { return server{accountItem} }

func (s server) Run(addr string) {
	e := echo.New()
	accountItem := e.Group("/account_item")
	accountItem.GET("", RouterFunc(s.accountItem.Get, nil))
	accountItem.POST("", RouterFunc(s.accountItem.Save, nil))
	e.Start(addr)
}
