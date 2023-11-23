package controller

import (
	"fmt"

	"github.com/inahym196/accountant/pkg/usecase"
)

type Reader interface {
	PostForm() map[string][]string
	Query() map[string][]string
}
type Writer interface {
	JSON(json []byte)
	Text(text string)
	SetStatus(code int)
}

type AccountItemController interface {
	Get(w Writer, r Reader)
	GetAll(w Writer, r Reader)
	Save(w Writer, r Reader)
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
		w.Text(err.Error())
		return
	}
	w.Text(fmt.Sprintf("%s %s %s %s", dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element))
}

func (c accountItemController) GetAll(w Writer, r Reader) {
	dtos, err := c.u.GetAll()
	if err != nil {
		w.Text(err.Error())
	}
	var texts string
	for _, dto := range *dtos {
		texts += dto.Title + " " + dto.JapaneseTitle + " " + dto.PeriodType + " " + dto.Element + "\n"
	}
	w.Text(texts)
}

func (c accountItemController) Save(w Writer, r Reader) {
	data := r.PostForm()
	title, jp_title, period, element := data["title"], data["jp_title"], data["period"], data["element"]
	if len(title) != 1 || len(jp_title) != 1 || len(period) != 1 || len(element) != 1 {
		w.Text("please specify only one title,jp_title,period,element")
		return
	}
	dto := usecase.AccountItemDTO{
		Title:         title[0],
		JapaneseTitle: jp_title[0],
		PeriodType:    period[0],
		Element:       element[0],
	}
	err := c.u.Save(dto)
	if err != nil {
		w.Text(err.Error())
	}
}
