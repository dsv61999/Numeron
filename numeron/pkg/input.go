package pkg

import(
	"fmt"
	"strconv" //  stringとintの方変換に必要
	"strings"
)

//  何桁でプレイするかの入力を処理する関数
func  (c_p *Computer_Player)InputDigitNum() {
	DisplaySentence("How many digits do you play?")
	var num int
	fmt.Scan(&num)
	//  0(文字もintでは0になる)と10桁以上を弾く処理
	if num == 0 || num > 10{
		DisplaySentence("Please enter an integer 1~10.")
		c_p.InputDigitNum()
	}
	DisplaySentence("Do you play with " + strconv.Itoa(num) + " digits? (Please press y or n)")
	var ok string
	fmt.Scan(&ok)
	if ok=="y" || ok=="yes"{
		c_p.Digits_num = num
	}else{
		c_p.InputDigitNum()
	}
}

//  対戦するコンピュータを選ぶ入力を処理する関数
func (c_p *Computer_Player) InputComputerSelect()  {
	DisplaySentence("Please select Computer.")
	fmt.Println("1..." + "Mike")
	fmt.Println("2..." + "Amy")
	fmt.Println("3..." + "Bob")
	fmt.Println("4..." + "Emma")
	var select_computer string
	fmt.Scan(&select_computer)
	switch select_computer {
		case "1":
			c_p.Computer_Name = "Mike"
		case "2":
			c_p.Computer_Name = "Amy"
		case "3":
			c_p.Computer_Name = "Bob"
		case "4":
			c_p.Computer_Name = "Emma"
		default:
			c_p.Computer_Name = "Mike"
		}
}

//  プレイヤーの予測した数字の入力を処理する関数
func (c_p *Computer_Player) InputNum()  {
	DisplaySentence("Please predict " + c_p.Computer_Name + "'s number")
	var num string
	fmt.Scan(&num)
	c_p.Player_Predict_Num = strings.Split(num, "")
	if(len(c_p.Player_Predict_Num) != c_p.Digits_num){
		DisplaySentence("Please predict " + strconv.Itoa(c_p.Digits_num) + " digits.")
		c_p.InputNum()
	}
	c_p.Turn++
}
