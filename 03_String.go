package main

import "fmt"

func main() {

	// *********************************
	// ********** 文 字 列 *************
	// *********************************

	fmt.Println(len("ハローワルド"))
	fmt.Println(len("ハローワルド"))
	// 문자열 길이 정해 주는 곳
	// 文字列の長さを返却(空白も含む)。

	fmt.Println("ハローワルド"[1])
	// 文字列の中に0番目ではなく1番目のバイト（ロ）で表現。

	fmt.Println("ハロー" + "ワルド")
	// 文字列と文字列の連結（文字列を足し算することは無意味）。

}
