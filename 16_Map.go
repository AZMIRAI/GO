package main

import (
	"fmt"
)

func main()  {

	// ******************************
	// *********** マップ ***********
	// ******************************

	map_01()
		// map_01を呼び出す

	map_02()
		// map_01を呼び出す
}

func map_01()  {
	x:=make(map[string]int)
		// ｘというマップを初期化する

	x["key"] = 10
		// ｘというマップに”key”って設定して０１を入れる

	fmt.Println(x["key"])
		// １０値を呼び出す
}

func map_02()  {

	elements:=make(map[string]string)
		// elementsというマップを作る～

	elements["H"] = "Hydrogen"
		// Hを押したらHydrogenを出す

	elements["He"] = "Helium"
		// Heを押したらHeliumを出す

	var input string;
		// 変数inputをstring資料型で宣言する

	fmt.Scanf("%s", &input);
		// %sの位置に変数inputの値が入る。

	fmt.Println(elements[input])
		// elementの”input”を出す
}
