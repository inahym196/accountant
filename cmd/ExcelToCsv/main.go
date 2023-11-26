package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/inahym196/accountant/pkg/util"
)

func ExcelToCsv(w io.Writer, path string, sheetIndex int) error {
	excel, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	SheetName := excel.GetSheetName(sheetIndex)
	rows, err := excel.Rows(SheetName)
	if err != nil {
		return err
	}
	csvw := csv.NewWriter(w)
	defer csvw.Flush()

	var subject string
	cnt := 0
	for rows.Next() {
		cnt++
		cols := rows.Columns()
		if strings.Contains(cols[0], "科目一覧") {
			subject = strings.Split(cols[0], "　")[0]
			rows.Next()
			cnt++
			continue
		}
		if cols[0] == "" || cols[13] == "true" {
			continue
		}
		rowStr := []string{fmt.Sprintf("%04d", cnt), subject, cols[1], cols[8], cols[11], cols[12]}
		if err := csvw.Write(rowStr); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	file, _ := os.Create(filepath.Join(util.ProjectRoot(), "./full.csv"))
	if err := ExcelToCsv(file, filepath.Join(util.ProjectRoot(), "./AccountList.xlsx"), 3); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
