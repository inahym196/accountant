package e2e_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/inahym196/accountant/pkg/infra"
	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	"github.com/inahym196/accountant/pkg/interface/database"
	"github.com/inahym196/accountant/pkg/usecase"
	"github.com/inahym196/accountant/pkg/util"
	"github.com/labstack/echo/v4"
)

func setupAccountItem(path string) controller.AccountItemController {
	db_conn := infra.NewSQLiteConnector(path).GetConn()
	repo := database.NewAccountItemRepository(db_conn)
	i := usecase.NewAccountItemInteractor(repo)
	return controller.NewAccountItemController(i)
}

func TestAccountItemController(t *testing.T) {

	// Arrange
	accountItem := setupAccountItem(filepath.Join(util.ProjectRoot(), "./test.sqlite3"))
	req := httptest.NewRequest(http.MethodGet, "/?title=test", strings.NewReader(""))
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	// Action
	infra.RouterFunc(accountItem.Get, nil)(ctx)

	// Assert
	if rec.Code != http.StatusOK {
		t.Errorf("invalid responce: %v", rec.Code)
	}
	body, err := io.ReadAll(rec.Body)
	if err != nil {
		t.Error(err)
	}
	want := "test test instant assets"
	if string(body) != want {
		t.Errorf("invalid body: %v, want: %v", string(body), want)
	}
}
