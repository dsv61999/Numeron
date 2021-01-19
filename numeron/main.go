package main

import (
	// "fmt"
	// "github.com/360EntSecGroup-Skylar/excelize/v2"
	"numeron/pkg"
)

func main()  {
	pkg.DisplayCaption("NumeronApp", 2)
	num_digits := pkg.InputDigitNum()

	var c_p pkg.Computer_Player
	c_p.SetComputer(num_digits)
	c_p.InputNum(num_digits)

	// computer := pkg.SetComputer(num_digits)
	// f := excelize.NewFile()
	// // Create a new sheet.
	// index := f.NewSheet("Sheet2")
	// // Set value of a cell.
	// f.SetCellValue("Sheet2", "A2", "Hello world.")
	// f.SetCellValue("Sheet1", "B2", 100)
	// // Set active sheet of the workbook.
	// f.SetActiveSheet(index)
	// // Save spreadsheet by the given path.
	// if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }
}