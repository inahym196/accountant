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

type GetHandler struct {
	rt router
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.rt.c.GET(writer{w}, reader{r})
}

type router struct {
	c   *controller.AccountItemController
	mux *http.ServeMux
}

type Router interface {
	Run(addr string)
}

func NewRouter(c *controller.AccountItemController) Router { return router{c, http.NewServeMux()} }

func (rt router) Run(addr string) {
	rt.mux.Handle("/", GetHandler{rt})
	http.ListenAndServe(addr, rt.mux)
}
