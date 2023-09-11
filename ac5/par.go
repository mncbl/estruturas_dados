package main
import "fmt"

func main(){
	var x int
	c:= [6]int{1,2,3,4,5,6}
	fmt.Println("digite o valor que deseja ser proucurado")
	fmt.Scanln(&x)
	fmt.Println(achaPar(c, x))



}
func achaPar(v[6]int, x int)int{

	inicio:= 0
	final := len(v)-1
	for inicio < final{
		soma:= v[inicio]+ v[final]

	if soma == x  {
		fmt.Println("O Par que foi encontrado foi ", v[inicio], v[final])
		break


	}else if soma < x {
		inicio++

	}else{
		final--
	}




	}
	fmt.Print("Nao foi encontrado nenhum par -1 ")
	return -1



}