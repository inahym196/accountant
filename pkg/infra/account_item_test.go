package infra

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	controller "github.com/inahym196/accountant/pkg/interface/controller/http"
	"github.com/inahym196/accountant/pkg/interface/database"
	"github.com/inahym196/accountant/pkg/usecase"
)

func TestAccountItemHandler(t *testing.T) {
	db_conn := NewSQLiteConnector("../../test.sqlite3")
	repo := database.NewAccountItemRepository(db_conn.Conn)
	i := usecase.NewAccountItemInteractor(repo)
	c := controller.NewAccountItemController(i)
	s := NewAccountItemHandler(c)
	ts := httptest.NewServer(s)
	defer ts.Close()

	cli := &http.Client{}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, ts.URL+"?title=test", strings.NewReader(""))
	if err != nil {
		t.Error(err)
	}

	res, err := cli.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid responce: %v", res)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()
	want := "test test instant assets\n"
	if string(body) != want {
		t.Errorf("invalid body: %v, want: %v", string(body), want)
	}
}
