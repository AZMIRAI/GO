package main

import "fmt"

func main()  {
	practice_01();
		// practice_01 호출

	practice_02();
		// practice_02 호출

	for_01();
		// for_01 호출

	for_02();
		// for_02 호출

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
		// 1에서 10까지 한줄로 치기

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
	// 1에서 10까지 반복해서 치기

}

func for_01()  {
	i := 1;
		// 변수 i를 선언

	for i < 11 {
		// 혹시 i가 11보다 작다면

		fmt.Println(i);
		i ++;
			// i를 출력하고 i값을 올리자

	}
}

func for_02()  {
	for i:=1; i<11; i++{
		// 자바랑 똑같이 선언하면 됨.. 하지만 괄호가 필요없다

		fmt.Println(i);
			// i를 출력

	}
}
