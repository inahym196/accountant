package main

import (
	"path/filepath"

	"github.com/inahym196/accountant/pkg/infra"
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	"github.com/inahym196/accountant/pkg/interface/database"
	"github.com/inahym196/accountant/pkg/usecase"
	"github.com/inahym196/accountant/pkg/util"
)

func newServer(path string) infra.Server {
	db_conn := infra.NewSQLiteConnector(path).GetConn()
	repo := database.NewAccountItemRepository(db_conn)
	i := usecase.NewAccountItemInteractor(repo)
	c := controller.NewAccountItemController(i)
	return infra.NewServer(c)
}

func main() {
	server := newServer(filepath.Join(util.ProjectRoot(), "./test.sqlite3"))
	server.Run("localhost:8080")
}
