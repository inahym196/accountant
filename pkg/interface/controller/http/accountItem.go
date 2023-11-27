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
	PathParam() map[string]string
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
	subject, name := ctx.PathParam()["subject"], ctx.PathParam()["name"]

	if name == "" || subject == "" {
		ctx.Text("please specify only one name, and subject")
		return
	}
	dto, err := c.u.FindByTitle(subject, name)
	if err != nil {
		ctx.Text(err.Error())
		return
	}
	ctx.Text(fmt.Sprintf("%s %s %s %s %s", dto.Subject, dto.Name, dto.JPName, dto.PeriodType, dto.Balance))
}

func (c accountItemController) GetAll(ctx Context) {
	dtos, err := c.u.GetAll()
	if err != nil {
		ctx.Text(err.Error())
	}
	var texts string
	for _, dto := range dtos {
		texts += dto.Subject + " " + dto.Name + " " + dto.JPName + " " + dto.PeriodType + " " + dto.Balance + "\n"
	}
	ctx.Text(texts)
}

func (c accountItemController) Save(ctx Context) {
	data := ctx.PostForm()
	subject, name, jpname, period, balance := data["subject"], data["name"], data["jpname"], data["period"], data["balance"]
	if len(subject) != 1 || len(name) != 1 || len(jpname) != 1 || len(period) != 1 || len(balance) != 1 {
		ctx.Text("please specify only one subject,name,jpname,period,balance")
		return
	}
	dto := usecase.AccountItemDTO{
		Subject:    subject[0],
		Name:       name[0],
		JPName:     jpname[0],
		PeriodType: period[0],
		Balance:    balance[0],
	}
	err := c.u.Save(dto)
	if err != nil {
		ctx.Text(err.Error())
	}
}
