package controller

import (
	"fmt"

	"github.com/inahym196/accountant/pkg/usecase"
)

type Context interface {
	Writer
	Reader
}

type Writer interface {
	Text(text string)
}
type Reader interface {
	PostForm() map[string][]string
	Query() map[string][]string
}

type AccountItemController interface {
	Get(ctx Context)
	GetAll(ctx Context)
	Save(ctx Context)
}

type accountItemController struct {
	u usecase.AccountItemUseCase
}

func NewAccountItemController(u usecase.AccountItemUseCase) AccountItemController {
	return accountItemController{u}
}

func (c accountItemController) Get(ctx Context) {
	title := ctx.Query()["title"]
	if len(title) != 1 {
		ctx.Text("please specify only one title")
		return
	}
	dto, err := c.u.FindByTitle(title[0])
	if err != nil {
		ctx.Text(err.Error())
		return
	}
	ctx.Text(fmt.Sprintf("%s %s %s %s", dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element))
}

func (c accountItemController) GetAll(ctx Context) {
	dtos, err := c.u.GetAll()
	if err != nil {
		ctx.Text(err.Error())
	}
	var texts string
	for _, dto := range *dtos {
		texts += dto.Title + " " + dto.JapaneseTitle + " " + dto.PeriodType + " " + dto.Element + "\n"
	}
	ctx.Text(texts)
}

func (c accountItemController) Save(ctx Context) {
	data := ctx.PostForm()
	title, jp_title, period, element := data["title"], data["jp_title"], data["period"], data["element"]
	if len(title) != 1 || len(jp_title) != 1 || len(period) != 1 || len(element) != 1 {
		ctx.Text("please specify only one title,jp_title,period,element")
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
		ctx.Text(err.Error())
	}
}
