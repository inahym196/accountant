package main

import (
	"path/filepath"

	"github.com/inahym196/accountant/pkg/infra"
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	"github.com/inahym196/accountant/pkg/interface/database"
	"github.com/inahym196/accountant/pkg/usecase"
	"github.com/inahym196/accountant/pkg/util"
)

func main() {
	db_conn := infra.NewSQLiteConnector(filepath.Join(util.ProjectRoot(), "./test.sqlite3"))
	repo := database.NewAccountItemRepository(db_conn.Conn)
	i := usecase.NewAccountItemInteractor(repo)
	c := controller.NewAccountItemController(i)
	infra.NewServer(c).Run("localhost:8080")
}
