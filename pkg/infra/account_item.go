package infra

import (
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

func (rt accountItemRouter) HandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rt.c.GET(writer{w}, reader{r})
	}
}

type accountItemRouter struct {
	c *controller.AccountItemController
}

func NewAccountItemRouter(c *controller.AccountItemController) *accountItemRouter {
	return &accountItemRouter{c}
}
