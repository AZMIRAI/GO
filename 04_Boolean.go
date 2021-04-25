package main

import "fmt"

func main()  {

	// ******************************
	// ********** ブール値 *************
	// ******************************

	fmt.Println(true && true);
		// 真＆真＝真

	fmt.Println(true && false);
		// 真＆偽＝偽

	fmt.Println(true || true);
		// 真｜真＝真

	fmt.Println(true || false);
		// 真｜偽＝真

	fmt.Println(!true);
		// 真の偽＝偽

	fmt.Println(!!false);
		// 偽の偽の偽＝偽

}
