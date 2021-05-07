package main

import (
	"fmt"
)


var input string;
	// 変数inputをstring資料型で宣言する

func main()  {

	// ******************************
	// *********** マップ ***********
	// ******************************

	fmt.Scanf("%s", &input);
		// %sの位置に変数inputの値が入る。

	map_01()
		// map_01を呼び出す

	map_02()
		// map_02を呼び出す

	map_03()
		// map_03を呼び出す

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

	elements["Li"] = "Lithium"
		// Liを押したらLithiumを出す

	elements["Be"] = "Beryllium"
		// Beを押したらBerylliumを出す

	elements["B"] = "Boron"
		// Bを押したらBoronを出す

	elements["C"] = "Carbon"
		// Cを押したらCarbonを出す

	elements["He"] = "Helium"
		// Heを押したらHeliumを出す

	elements["F"] = "Fluorine"
		// Fを押したらFluorineを出す

	elements["Ne"] = "Neon"
		// Nを押したらNeonを出す


	fmt.Println(elements[input])
	// elementの”input”を出す

	fmt.Println(elements["un"])
		// 多分、これしたら何も出てこないと思う

	name, ok := elements["H"]
		// name、okというelementsを宣言してHを登録する

	fmt.Println(name, ok)
		// name、okを出力する

	elements_1 := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}
		// 上のコードをもっと簡単に作れる

	fmt.Println(elements_1)
		// elements_1を出力する
		
}

func map_03()  {
	// このfunctionは普通の状態をセーブできるように作ったコードだ

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B": map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C": map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N": map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O": map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F": map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne": map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}
	if el, ok := elements[input]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

