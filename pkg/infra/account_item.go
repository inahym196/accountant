package infra

import (
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

func (s accountItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.c.GET(writer{w}, reader{r})
	}
}

type accountItemHandler struct {
	c controller.AccountItemController
}

func NewAccountItemHandler(c controller.AccountItemController) accountItemHandler {
	return accountItemHandler{c}
}
