package infra

import (
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

func (s accountItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.c.Get(writer{w}, reader{r})
	case http.MethodPost:
		s.c.Save(writer{w}, reader{r})
	}
}

type accountItemHandler struct {
	c controller.AccountItemController
}

func NewAccountItemHandler(c controller.AccountItemController) accountItemHandler {
	return accountItemHandler{c}
}
