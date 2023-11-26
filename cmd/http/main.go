package main

import (
	"path/filepath"

	"github.com/inahym196/accountant/pkg/infra"
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	"github.com/inahym196/accountant/pkg/interface/database"
	"github.com/inahym196/accountant/pkg/usecase"
	"github.com/inahym196/accountant/pkg/util"
)

func setupAccountItem(path string) controller.AccountItemController {
	db_conn := infra.NewSQLiteConnector(path).GetConn()
	repo := database.NewAccountItemRepository(db_conn)
	i := usecase.NewAccountItemInteractor(repo)
	return controller.NewAccountItemController(i)
}

func main() {
	accountItem := setupAccountItem(filepath.Join(util.ProjectRoot(), "./test.sqlite3"))
	server := infra.NewServer(accountItem)
	server.Run("localhost:8080")
}
