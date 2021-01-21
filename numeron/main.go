package main

import (
	// "fmt"
	// "github.com/360EntSecGroup-Skylar/excelize/v2"
	"numeron/pkg"
)

func main()  {
	pkg.DisplayCaption("NumeronApp", 2)
	pkg.DisplayAA("Hello")

	var c_p pkg.Computer_Player
	c_p.Turn = 0
	c_p.InputDigitNum()
	c_p.InputComputerSelect()
	c_p.SetComputer()
	c_p.InputNum()
	c_p.Count_Comp_Eat_Bite()

}