package pkg

import(
	"fmt"
	"strings"
)

//  文字を枠線付きで表示する関数
func DisplayCaption(s string, blamk_space int)  {
	length_s := len(s)
	// blamk_space := 1
	fmt.Println("*" + strings.Repeat("-", length_s+blamk_space*2) + "*")
	fmt.Println("|" + strings.Repeat(" ", blamk_space) + s + strings.Repeat(" ", blamk_space) +"|")
	fmt.Println("*" + strings.Repeat("-", length_s+blamk_space*2) + "*")
}

//  文字を||つきで表示する関数
func DisplaySentence(s string)  {
	fmt.Println("| " + s + " |")
}

//  アスキーアートを表示する関数
func DisplayAA(s string)  {
	fmt.Println("   " + s)
	fmt.Println("  ／ーーー＼")
	fmt.Println(" /・　 ・　 ＼/|")
	fmt.Println("／ーーー＼　　 |")
	fmt.Println("｜ ――― ｜　　｜|")
	fmt.Println("＼＿＿＿／　　 |")
	fmt.Println("　　＼　　　　／")
	fmt.Println("　　　ーーーー)")

} 
