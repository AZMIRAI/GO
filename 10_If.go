package main

import "fmt"

func main()  {

	// ******************************
	// ********** IF 文 *************
	// ******************************

	sniffling();
 		//sniffling 呼出し

}

func sniffling()  {
	for i:=0; i<10; i++{
		if i % 2 == 0{
			//もし2を割り算して残りが0だったら

			fmt.Println("偶数",i);
		} else {
			// それじゃなかったら

			fmt.Println("奇数",i);
		}
	}
}
