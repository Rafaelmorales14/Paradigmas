package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	var numero int

	fmt.Println("Digite o numero para a tabuada: ")
	fmt.Scanf("%d", &numero)
	fmt.Print("\n")
	
	for i:=0; i <= 10; i++ {
	    var resultado = numero * i
	    fmt.Printf("%d X %d = %d\n", numero, i, resultado)
	}
}

