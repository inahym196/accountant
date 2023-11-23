package controller

import (
	"fmt"

	"github.com/inahym196/accountant/pkg/usecase"
)

type Reader interface {
	Query() map[string][]string
}
type Writer interface {
	JSON(json []byte)
	Text(text string)
	SetStatus(code int)
}

type AccountItemController struct {
	u usecase.AccountItemUseCase
}

func NewAccountItemController(i usecase.AccountItemUseCase) *AccountItemController {
	return &AccountItemController{i}
}

func (c AccountItemController) GET(w Writer, r Reader) {
	title := r.Query()["title"]
	if len(title) != 1 {
		w.Text("please specify only one title")
		return
	}
	dto, err := c.u.FindByTitle(title[0])
	if err != nil {
		w.Text("error")
		return
	}
	w.Text(fmt.Sprintf("%s %s %s %s", dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element))
}
