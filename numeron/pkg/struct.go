package pkg

//  コンピュータのゲーム内の情報を保持する構造体
type Computer_Player struct{
	Computer_Name string  //  コンピュータの名前
	Computer_Num []string  //  コンピュータの設定した数字
	Computer_Eaten_Count int  //  コンピュータがEatされた数
	Computer_Bitten_Count int  //  コンピュータがBiteされた数
	Player_Name string
	Player_Num []string
	Player_Predict_Num []string
	Digits_num int  //  プレイしている桁数
	Turn int  //  ターン数
}


