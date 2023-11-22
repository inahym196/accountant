package infra

import (
	"fmt"
	"net/http"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
)

type writer struct {
	w http.ResponseWriter
}

func (w writer) JSON(json []byte) {
	fmt.Fprintln(w.w, string(json))
}
func (w writer) Text(text string) {
	fmt.Fprintln(w.w, string(text))
}
func (w writer) SetStatus(code int) {
	w.w.WriteHeader(code)
}

type reader struct {
	r *http.Request
}

func (r reader) Query() map[string][]string {
	return r.r.URL.Query()
}

type Router struct{ ai *accountItemRouter }

func NewRouter(c *controller.AccountItemController) Router { return Router{NewAccountItemRouter(c)} }
func (rt Router) Run(addr string) {
	http.HandleFunc("/account_item", rt.ai.HandleFunc)
	http.ListenAndServe(addr, nil)
}
