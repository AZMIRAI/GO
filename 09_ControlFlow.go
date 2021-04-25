package main

import "fmt"

func main()  {

	// ******************************
	// ******* コントロール文 ********
	// ******************************

	practice_01();
		// practice_01 呼出し

	practice_02();
		// practice_02 呼出し

	for_01();
		// for_01 呼出し

	for_02();
		// for_02 呼出し

}

func practice_01()  {
	fmt.Println("1");
	fmt.Println("2");
	fmt.Println("3");
	fmt.Println("4");
	fmt.Println("5");
	fmt.Println("6");
	fmt.Println("7");
	fmt.Println("8");
	fmt.Println("9");
	fmt.Println("10");
		// 1から10まで1行ずつ出力

}

func practice_02()  {
	fmt.Println(`1
2
3
4
5
6
7
8
9
10
`)
	// 1から10まで1行ずつ出力

}

func for_01()  {
	i := 1;
		// 変数iを宣言

	for i < 11 {
		// もしiが11より小さかったら

		fmt.Println(i);
		i ++;
			// iを出力してi値を上げよう

	}
}

func for_02()  {
	for i:=1; i<11; i++{
		// Javaと同じみたいに宣言したらいい しかし（）がいらん

		fmt.Println(i);
			// iを出力

	}
}
