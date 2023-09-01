package main

import "fmt"

func main() {
	var n int
	fmt.Print("informe 0 número de discos, SUPERIOR A 0 E INFERIOR A 16!!!: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("impossivel ter pecas negativas ")
		return
	}
	if n>15 {
		fmt.Println("O numero de movmentos maior que 60.000")
		return


	}

	movimento := hanoi(n, "1º torre", "2ºtorre", "3ºtorre")
	fmt.Println("Movimentos necessários: ", movimento)

}

func hanoi(n int, origem string, destino string, trabalho string) int {
	if  n >0 {
		movimento := 0
		movimento++
		movimento = movimento + hanoi(n-1, origem, trabalho, destino)
		fmt.Println("Mova o disc o", n, " de ", origem, " para ", destino)
		fmt.Println("_________________________________________________________")
		movimento = movimento + hanoi(n-1, trabalho, destino, origem)

		return movimento

	}
	return 0

}
