package main

import "fmt"

func main()  {
star_01();
star_02();
star_03();
}

func star_01()  {

	for i:=0; i<10; i++{
		fmt.Println("")
		for j:=0;j<10; j++{
			fmt.Print("*");
		}
	}
}

func star_02()  {
	for i:=0; i<10; i++ {
		for j:=0; j<i; j++{
		fmt.Print("*");
		}

		fmt.Println("")
		}
}

func star_03()  {
	for i:=0; i<10; i++{
		fmt.Println("")
		for j:=10; j>i; j--{
			if(j<10){
				fmt.Print("*")
			}

		}
	}
}