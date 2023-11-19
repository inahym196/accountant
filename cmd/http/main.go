package main

import (
	"fmt"
	"net/http"

	"github.com/inahym196/accountant/pkg/infra"
	"github.com/inahym196/accountant/pkg/infra/sqlite"
	my_handler "github.com/inahym196/accountant/pkg/interface/handler"
	"github.com/inahym196/accountant/pkg/usecase"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	fmt.Fprintf(w, "%s", body)
}

func main() {
	db_conn := infra.NewSQLiteConnector("./test.sqlite3")
	repo := sqlite.NewAccountItemRepository(db_conn.Conn)
	i := usecase.NewAccountItemInteractor(repo)
	handler := my_handler.NewAccountItemHandler(i)
	http.HandleFunc("/", handler.Get)
	http.ListenAndServe("localhost:8080", nil)
}
