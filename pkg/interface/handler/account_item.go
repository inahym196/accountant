package handler

import (
	"fmt"
	"net/http"

	"github.com/inahym196/accountant/pkg/usecase"
)

type AccountItemHandler struct {
	u usecase.IAccountItemUseCase
}

func NewAccountItemHandler(i usecase.IAccountItemUseCase) *AccountItemHandler {
	return &AccountItemHandler{i}
}

func (h AccountItemHandler) Get(w http.ResponseWriter, r *http.Request) {
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
