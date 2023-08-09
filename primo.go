package main

import "fmt"

func main (){
	fmt.Println(ePrimo(7))
	fmt.Println(ePrimo(10))
	fmt.Println(fibo(8))
	fmt.Println(dia(2))
	fmt.Println(bissexto(2017))
}

func ePrimo(n int) bool{
	var primo bool
	primo =true
	for i := 2; i < n; i++ {
		if n % i == 0 {
			primo = false
			fmt.Println(i)
			fmt.Println("------")

		}
	}
	return primo


}

func fibo (n int ) int{
	var f1, f2 int  = 0 , 1
	fmt.Println("1")
	for i := 0; i < n-1; i++ {
		f2, f1 = (f1+ f2), f2

		fmt.Println(f2)


	}

	fmt.Println("___________________")
	return f1
}
func dia (n int) string {
	switch n {
	case 1, 8, 15, 22, 29:
		fmt.Println("domingo")
	case 2,9, 16, 23, 30:
		fmt.Println("segunda")
	case 3, 10, 17, 24, 31:
		fmt.Println("Terça")
	case 4, 11, 18, 25:
		fmt.Println("Quarta")
	case 5,12, 19, 26:
		fmt.Println("Quinta")
	case 6,13, 20, 27:
		fmt.Println("Sexta")
	case 7,14, 21, 28:
		fmt.Println("Sabado")
	default:
		fmt.Println("dia invalido ")

		}
		return  ("_____________________")
}


func bissexto(n int ) bool {
	var ver int
	ver = n % 4
	if ver == 0 {
			return true
	}
	return false



}