package main

import "fmt"

func main()  {

	// ******************************
	// ********** 불 린 *************
	// ******************************

	fmt.Println(true && true);
		// 참이고 참이면 참
	fmt.Println(true && false);
		// 참이고 거짓이면 거짓
	fmt.Println(true || true);
		// 참이거나 참이면 참
	fmt.Println(true || false);
		// 참이거나 거짓이면 참
	fmt.Println(!true);
		// 참의 거짓이면 거짓
	fmt.Println(!!false);
		// 거짓의 거짓의 거짓 거짓
		}
