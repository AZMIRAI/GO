package main

import "fmt"

func main()  {

	// ******************************
	// ********** 練 習 *************
	// ******************************


	fmt.Print("数字を入力してください : ");
		// 文字列を出力する

	var input float64;
		// 変数inputをfloat64資料型で宣言する

	fmt.Scanf("%f", &input);
		// %fの位置に変数inputの値が入る。

	output := input * 2;
		// 変数outputを宣言して、値は変数input値の2倍に指定する。

	fmt.Println(output);
		// 変数outputを出力する。

}
// 6월 6일.. 
