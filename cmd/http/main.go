package main

import (
	"path/filepath"

	"github.com/inahym196/accountant/pkg/infra"
	"github.com/inahym196/accountant/pkg/infra/echo"
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	"github.com/inahym196/accountant/pkg/interface/database"
	"github.com/inahym196/accountant/pkg/usecase"
	"github.com/inahym196/accountant/pkg/util"
)

func main() {
	db_conn := infra.NewSQLiteConnector(filepath.Join(util.ProjectRoot(), "./test.sqlite3")).GetConn()
	repo := database.NewAccountItemRepository(db_conn)
	i := usecase.NewAccountItemInteractor(repo)
	c := controller.NewAccountItemController(i)
	echo.NewServer(c).Run("localhost:8080")
}
