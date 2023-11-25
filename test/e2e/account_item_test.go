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

func TestAccountItemHandler(t *testing.T) {

	// Arrange
	db_conn := infra.NewSQLiteConnector(filepath.Join(util.ProjectRoot(), "./test.sqlite3")).GetConn()
	repo := database.NewAccountItemRepository(db_conn)
	i := usecase.NewAccountItemInteractor(repo)
	c := controller.NewAccountItemController(i)
	s := infra.NewServer(c)
	req := httptest.NewRequest(http.MethodGet, "/account_item?title=test", strings.NewReader(""))
	rec := httptest.NewRecorder()

	// Action
	s.GetAccountItem(echo.New().NewContext(req, rec))

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
