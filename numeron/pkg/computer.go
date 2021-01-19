package pkg

import(
	"math/rand"  //  乱数の生成に必要
	"time"  // 
	"strconv"  //  stringとintの方変換に必要
	// "fmt"
)

//  コンピュータの情報を保持する構造体
type Computer_Player struct{
	Computer_Name string
	Computer_Num []string
	Player_Name string
	Player_Num []string
	Player_Predict_Num []string
	// Predict_num []string
	// Digits_num int
}

//  コンピュータをセットアップする関数
func (c_p *Computer_Player) SetComputer(num_digits int) {
	// var computer Computer
	c_p.Computer_Name = "Mike"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_digits; i++ {
		// c_p.Computer_Num = append(c_p.Computer_Num, strconv.Itoa(rand.Intn(10)))  //  0〜9の乱数を追加
		c_p.Computer_Num = append(c_p.Computer_Num, strconv.Itoa(i)) //  デバッグ用
	}
	DisplaySentence("Hello. My name is " + c_p.Computer_Name + ". " + "Let's play Numeron !!")
}

//  プレイヤーの予測した数字がコンピュータの数字にどれだけEatとByteしているかを計算する関数
// func (c_p *Computer_Player)  {
	
// }
