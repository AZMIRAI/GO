package main

import "fmt"

func main()  {
	ex_01();
		// ex_01() 呼出し。

	ex_02();
		// ex_02() 呼出し。

}

func ex_01()  {
	// 1から100の間の数字で、3で割り切れる数(3、6、9、...)をすべて出力するプログラムを作成しなさい。

	for i:=1; i<=100; i++ {
		// 変数iを1に指定後、100まで1ずつ足す。

		if i % 3 == 0{
			// もし変数iを3で割った値が0になるとき。

			fmt.Print(i," ");
				// iを出力して1マス空ける。
		}
	}
}

func ex_02()  {
	//1から100まで出力するプログラムを作成しなさい。 3の倍数に対しては数字の代わりに「Fizz」を出力して
	//5の倍数に対しては「Buzz」を出力しなさい。3と5の公倍数に対しては"FizzBuzz"を出力しなさい。

	for i:=1; i<=100; i++ {
		// 1から100まで変数iを1ずつ足す。

		if (i % 5 == 0)&&(i % 3 == 0){
			// 変数iを5で割って残った値が0で、変数iを3で割った値が0だったら。

			fmt.Print("fizzbuzz ");
				// fizzbuzzを出力する。

		}else if i % 3 == 0 {
			// 変数iを3に分けて残った値が0だったら。

			fmt.Print("fizz ");
				// fizzを出力する。

		}else if i % 5 == 0 {
			// 変数iを5に分けて余りの値が0だったら。

			fmt.Print("buzz ");
				// buzzを出力する。

		}else {
			// 以外は

			fmt.Print(i, " ");
				// 変数iを出力する。

		}
	}
}