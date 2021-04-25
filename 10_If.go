package main

import "fmt"

func main()  {
 sniffling();
 	//sniffling 호출

}

func sniffling()  {
	for i:=0; i<10; i++{
		if i % 2 == 0{
			//혹시 2로 나눠서 나머지가 0이면

			fmt.Println("짝수",i);
		} else {
			// 아니면

			fmt.Println("홀수",i);
		}
	}
}