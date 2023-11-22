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

func (rt router) Get(w http.ResponseWriter, r *http.Request) {
	rt.c.GET(writer{w}, reader{r})
}

type router struct {
	c *controller.AccountItemController
}

func NewRouter(c *controller.AccountItemController) router { return router{c} }

func (rt router) Run(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rt.Get)
	http.ListenAndServe(addr, mux)
}
