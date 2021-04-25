package main

import "fmt"

func main(){

	// ******************************
	// ********** SWITCH 文 *********
	// ******************************

	switch_01();
		// switch_01 呼出し
}
func switch_01()  {
	i := 3;
	switch i {
		// スイッチを書きたい名前を書いて｛｝を開く

	case 0: fmt.Println("丸")
		// caseを使って希望の文字や数を書いて:を書くその次に自分の書きたい文字列を入れる。

	case 1: fmt.Println("一")
		// 上と同じ

	case 2: fmt.Println("二")
		// 上と同じ

	case 3: fmt.Println("三")
		// 上と同じ

	case 4: fmt.Println("四")
		// 上と同じ

	case 5: fmt.Println("五")
		// 上と同じ

	default: fmt.Println("知らない数字")
		// 何も該当しないとき、出力をする。

	}

}