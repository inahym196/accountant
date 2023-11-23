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

type AccountItemController interface {
	Get(w Writer, r Reader)
	Save(w Writer, r Reader) error
}

type accountItemController struct {
	u usecase.AccountItemUseCase
}

func NewAccountItemController(i usecase.AccountItemUseCase) AccountItemController {
	return accountItemController{i}
}

func (c accountItemController) Get(w Writer, r Reader) {
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

func (c accountItemController) Save(w Writer, r Reader) error {
	q := r.Query()
	title, jp_title, period, element := q["title"], q["jp_title"], q["period"], q["element"]
	if len(title) != 1 || len(jp_title) != 1 || len(period) != 1 || len(element) != 1 {
		w.Text("please specify only one title,jp_title,period,element")
		return nil
	}
	dto := usecase.AccountItemDTO{
		Title:         title[0],
		JapaneseTitle: jp_title[0],
		PeriodType:    period[0],
		Element:       element[0],
	}
	return c.u.Save(dto)
}
