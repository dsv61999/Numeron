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
func (player *Player) InputNum(computer_name string, digits_num int)  {
	DisplayCaption("Please predict " + computer_name + "'s number", 1)
	var num string
	fmt.Scan(&num)
	player.Predict_num = strings.Split(num, "")
	// fmt.Println(player.Predict_num[0])
	// fmt.Println(player.Predict_num[1])
	// fmt.Println(player.Predict_num[2])
}
