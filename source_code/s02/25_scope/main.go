// File name: ...\s02\25_scope\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

/*
Глобальный и локальный.
{ } - создают свой скоп
*/

package main

import "fmt"
// _ГП - глобальная перемнная
var g int = 10

// g := 20

func main() {

	fmt.Println(g) // 10

	{
		// скобки создают свой scope
		// доступ к гл.пер. есть
		g := 20
		fmt.Println(g) // 20
		f := 2.3
		fmt.Println(f) // 2.3
	}

	// fmt.Println(f) //compiler error - undefined 'f'

	fmt.Println(g) // 10
	g += 1
	fmt.Println(g) // 11
	{
		g += 4 // обращение к _ГП
		fmt.Println(g) //15
	}

	fmt.Println(g) // 15
	
	sayHello()
	// fmt.Println(s)		//compiler error - undefined 's'
}

func sayHello() {
	// у функци локальный scope
	fmt.Println("Hello!")
	fmt.Println(g)

	s := "from testScope()"
	fmt.Println(s)
}
