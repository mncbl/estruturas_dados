package main

import "fmt"

func main (){
	fmt.Println(ePrimo(7))
	fmt.Println(ePrimo(10))
	fmt.Println(fibo(7))
	fmt.Println(dia(2))
	fmt.Println(bissexto(1900))
	fmt.Println(bissexto(2012))



}

func ePrimo(n int) bool{
	var primo bool
	primo =true

	for i := 2; i < n; i++ {
		if n % i == 0 {
			primo = false
			fmt.Println(i)



		}
	}

	return primo

}


func fibo (n int ) int{
	fmt.Println("EXEC 2")
	var f1, f2 int  = 0 , 1
	fmt.Println("1")
	for i := 0; i < n-1; i++ {
		f2, f1 = (f1+ f2), f2

		fmt.Println(f2)


	}

	fmt.Println("				")
	return f2
}


func dia (n int) string {
	fmt.Println("EXEC 3")
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

func bissexto(n int) bool {
	fmt.Println("EXEC 4")
	var ver4, ver1 int
	ver1 = n % 100
	ver4 = n % 4
	if ver4 > 0 {
		fmt.Println("O ano NÃO é bissexto")
		return false
	} else if ver1 == 0 {
		fmt.Println("O ano NÃO é bissexto")
		return false
	}
	fmt.Println("O ano é SIM bissexto")
	return true
}
