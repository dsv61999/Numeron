package pkg

import(
	"fmt"
	"strconv" //  stringとintの方変換に必要
	"strings"
)

//  何桁でプレイするかの入力を処理する関数
func  InputDigitNum() int{
	DisplaySentence("How many digits do you play?")
	var num int
	fmt.Scan(&num)
	DisplaySentence("Do you play with " + strconv.Itoa(num) + " digits? (Please press y or n)")
	var ok string
	fmt.Scan(&ok)
	if ok=="y" || ok=="yes"{
		return num
		// computer.digits_num = num
	}else{
		return InputDigitNum()
	}
}

//  プレイヤーの予測した数字の入力を処理する関数
func (c_p *Computer_Player) InputNum(digits_num int)  {
	DisplayCaption("Please predict " + c_p.Computer_Name + "'s number", 1)
	var num string
	fmt.Scan(&num)
	c_p.Player_Num = strings.Split(num, "")
	fmt.Println(c_p.Player_Num[0])
	fmt.Println(c_p.Player_Num[1])
	fmt.Println(c_p.Player_Num[2])
}
