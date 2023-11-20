package controller

import (
	"fmt"
	"net/http"

	"github.com/inahym196/accountant/pkg/usecase"
)

type AccountItemController struct {
	u usecase.IAccountItemUseCase
}

func NewAccountItemController(i usecase.IAccountItemUseCase) *AccountItemController {
	return &AccountItemController{i}
}

func (h AccountItemController) Get(w http.ResponseWriter, r *http.Request) {
	if title := r.URL.Query().Get("title"); title != "" {
		dto, err := h.u.FindByTitle(title)
		if err != nil {
			fmt.Fprintf(w, "error")
			return
		}
		fmt.Fprintf(w, "%s %s %s %s", dto.Title, dto.JapaneseTitle, dto.PeriodType, dto.Element)
	} else {
		fmt.Fprintln(w, "failed to parse")
	}
}
