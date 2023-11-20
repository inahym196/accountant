package main

import (
	"net/http"

	"github.com/inahym196/accountant/pkg/interface/database"
	ai_handler "github.com/inahym196/accountant/pkg/interface/handler"
	"github.com/inahym196/accountant/pkg/usecase"
)

func main() {
	db_conn := database.NewSQLiteConnector("./test.sqlite3")
	repo := database.NewAccountItemRepository(db_conn.Conn)
	i := usecase.NewAccountItemInteractor(repo)
	handler := ai_handler.NewAccountItemHandler(i)
	http.HandleFunc("/", handler.Get)
	http.ListenAndServe("localhost:8080", nil)
}
