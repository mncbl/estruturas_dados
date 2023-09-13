package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0
	fmt.Println(" tentando colcando a lista de forma crescente e add", 4)
	insereOrd(&lista, &n, 4)
	fmt.Println(lista)
	fmt.Println("   ")
	fmt.Println("tentandocolcando a lista de forma crescente e add", 12)
	insereOrd(&lista, &n, 12)
	fmt.Println(lista)
	fmt.Println("   ")
	fmt.Println("tentando colcando a lista de forma crescente e add", 2)
	insereOrd(&lista, &n, 2)
	fmt.Println(lista)
	fmt.Println("   ")
	fmt.Println("tentando colcando a lista de forma crescente e add", 6)
	insereOrd(&lista, &n, 6)
	fmt.Println(lista)
	fmt.Println("    ")
	fmt.Println("tentandocolocar na lista de forma crescente e add", 17)
	insereOrd(&lista, &n, 17)
	fmt.Println(lista)
	fmt.Println("   ")
	fmt.Println("tentando colocar a lista de forma crescente e add", 1)
	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n != M {

		fmt.Println("adicionado com sucesso")

	}else{
		fmt.Println("overflow")
		return

	}

	if busca1(*v, *n, novoValor) != -1 {
		fmt.Println("erro! o elemento já existe na lista")
		return
	}

	i := *n - 1

	for i >= 0 && v[i] > novoValor { // i > 0 verificando se o indice  de i vai ta dentro da lista

		v[i+1] = v[i]
		i--

	}

	v[i+1] = novoValor
	*n++

}

func busca1(v [M]int, n, x int) int {
	i := 0
	for i < n {
		if v[i] == x {
			return i
		}
		i++
	}
	return -1
}
