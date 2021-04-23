package main

import "fmt"

func main()  {
	fmt.Print("数字を入力してください : ");
		// 문자열 출력한다

	var input float64;
		// 변수 input를 float64자료형으로 선언한다

	fmt.Scanf("%f", &input);
		// %f 자리에 변수 input의 값이 들어가게된다

	output := input * 2;
		// 변수 output을 선언하고 값은 변수 input 값의 2배로 지정한다

	fmt.Println(output);
		// 변수 output을 출력한다

}