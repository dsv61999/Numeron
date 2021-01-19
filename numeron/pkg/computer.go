package pkg

import(
	"math/rand"  //  乱数の生成に必要
	"time"  // 
	"strconv"  //  stringとintの方変換に必要
	// "fmt"
)

//  コンピュータの情報を保持する構造体
type Computer struct{
	Name string
	Num []string
	// Predict_num []string
	// Digits_num int
}

//  コンピュータをセットアップする関数
func (computer *Computer) SetComputer(num_digits int) {
	// var computer Computer
	computer.Name = "Mike"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_digits; i++ {
		// computer.Num = append(computer.Num, strconv.Itoa(rand.Intn(10)))  //  0〜9の乱数を追加
		computer.Num = append(computer.Num, strconv.Itoa(i)) //  デバッグ用
	}
	DisplaySentence("Hello. My name is " + computer.Name + ". " + "Let's play Numeron !!")
}

//  プレイヤーの予測した数字がコンピュータの数字にどれだけEatとByteしているかを計算する関数
func (computer *Computer)  {
	
}
