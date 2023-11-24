package infra

import (
	"fmt"
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

type context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c context) Text(text string) {
	c.w.WriteHeader(http.StatusOK)
	fmt.Fprintln(c.w, string(text))
}

func (c context) Query() map[string][]string {
	return c.r.URL.Query()
}

func (c context) PostForm() map[string][]string {
	c.r.ParseForm()
	return c.r.PostForm
}

type Server struct {
	c controller.AccountItemController
}

func NewServer(c controller.AccountItemController) Server { return Server{c} }
func (s Server) Run(addr string) {
	http.HandleFunc("/account_item", s.AccountItemHandleFunc)
	http.ListenAndServe(addr, nil)
}

func (s Server) AccountItemHandleFunc(w http.ResponseWriter, r *http.Request) {
	ctx := context{w, r}
	switch r.Method {
	case http.MethodGet:
		s.c.Get(ctx)
	case http.MethodPost:
		s.c.Save(ctx)
	}
}
