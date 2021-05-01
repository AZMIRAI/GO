package main

import "fmt"

func main()  {
	slice_01();
		// slice_01を呼び出す

	slice_02();
		// slice_02を呼び出す

	slice_03();
		// slice_03を呼び出す

}

func slice_01()  {
	var x []float64;
		// sliceを作る

	x = make([]float64, 5);
		// sliceに５つの空間を作る！

	fmt.Println(x);
		// ｘというsliceを出力する

}

func slice_02()  {
	slice1 := []int{1,2,3};
		// slice1の関数を作る

	slice2 := append(slice1, 4,5);
		// slice2の関数を作る、関数を作るときは関数で呼び出して作れる

	fmt.Println(slice1, slice2);
		// slice1、slice2を出力する

}

func slice_03()  {
	slice1 := []int{1,2,3,4,6666};
		// slice1というsliceを作る

	slice2 := make([]int,7);
		// slice2というsliceを作る(空間は７)

	copy(slice2,slice1);
		// slice1をslice2にコピーする

	fmt.Println(slice1,slice2);
		// slice1,slice2を出力する

}