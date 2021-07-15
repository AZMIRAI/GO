package main

import "fmt"

func main() {

	// Raw String Literal. 복수라인
	rawLiteral := `아리랑\n
아리랑\n
	아라리요`

	// Interpreted String Literal

	interLiteral := "아리랑아리랑\n아라리요"
	// 아래와 같이 + 를 사용해서 두라인에 걸쳐 사용 할 수도 있다
	// interLiteral := "아리랑아리랑\n" +
	// "아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	forex()
	scan()
}

func forex() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

func scan() {
	var age int
	var name string

	fmt.Print("이름과 나이를 띄어쓰기로 구분하여 입력하세요 : ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("이름 : %s\n나이 : %d\n", name, age)

}
