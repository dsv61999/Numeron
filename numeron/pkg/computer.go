package pkg

import(
	"math/rand"  //  乱数の生成に必要
	"time"  // 
	"strconv"  //  stringとintの方変換に必要
	// "fmt"
)

//  コンピュータをセットアップする関数
func (c_p *Computer_Player) SetComputer() {
	// c_p.Computer_Name = "Mike"
	rand.Seed(time.Now().UnixNano())

	//  重複がないようにコンピュータ数字を決める処理
	for{
		new_digit := rand.Intn(10)
		do_continue := false
		for _, j := range c_p.Computer_Num {
			if j == strconv.Itoa(new_digit) {
				do_continue = true
				break
			}
		}
		if do_continue{
			continue
		}
		c_p.Computer_Num = append(c_p.Computer_Num, strconv.Itoa(new_digit))
		if len(c_p.Computer_Num) == c_p.Digits_num{
			break
		}
	}
	//  表示処理
	DisplayCaption("Hello. My name is " + c_p.Computer_Name + ". " + "Let's play Numeron !!", 1)

	//  デバッグ用
	// num := ""
	// for i := 0; i < c_p.Digits_num; i++ {
	// 	num += c_p.Computer_Num[i]
	// }	
	// DisplayCaption("You win!!! " + "My number is " + num + "!!!", 2)
}

//  プレイヤーの予測した数字がコンピュータの数字にどれだけEatとBiteしているかを計算する関数
func (c_p *Computer_Player) Count_Comp_Eat_Bite() {
	//  Eatをカウントする処理
	c_p.Computer_Eaten_Count = 0
	for i := 0; i < c_p.Digits_num; i++ {
		if c_p.Computer_Num[i] == c_p.Player_Predict_Num[i]{
			c_p.Computer_Eaten_Count++
		}
	}
	//  Biteをカウントする処理
	c_p.Computer_Bitten_Count = 0
	for i := 0; i < c_p.Digits_num; i++ {
		for j := 0; j < len(c_p.Player_Predict_Num); j++ {
			if c_p.Computer_Num[i] == c_p.Player_Predict_Num[j]{
				c_p.Computer_Bitten_Count++
			}
		}
	}
	c_p.Computer_Bitten_Count = c_p.Computer_Bitten_Count - c_p.Computer_Eaten_Count

	//  表示処理
	DisplaySentence(strconv.Itoa(c_p.Computer_Eaten_Count) + "Eat. " + strconv.Itoa(c_p.Computer_Bitten_Count) + "Bite.")

	//  正解した時としていないときの処理
	if c_p.Computer_Eaten_Count == c_p.Digits_num{
		num := ""
		for i := 0; i < c_p.Digits_num; i++ {
			num += c_p.Computer_Num[i]
		}
		DisplayCaption("You win!!! " + "My number is " + num + "!!!", 2)
		if c_p.Turn == 1{
			DisplaySentence("It took " + strconv.Itoa(c_p.Turn) + " turn.")
		}else{
			DisplaySentence("It took " + strconv.Itoa(c_p.Turn) + " turns.")
		}
	}else{
		c_p.InputNum()
		c_p.Count_Comp_Eat_Bite()
	}
}


