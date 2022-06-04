package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()

	for i := 1; i <= 1000; i++ {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i), i)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i), "samakata chloe")
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
