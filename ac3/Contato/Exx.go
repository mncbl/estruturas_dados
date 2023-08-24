package Contato // O nome do pacote deve ser maiúsculo para ser acessível de outros pacotes
import "fmt"

type Contato struct {
	Nome  string
	Email string
}

func ExcluiContato(a *[5]Contato){
	if (a[0] == Contato{}) {
		fmt.Println("Lista de contatos vazia!")

	}

	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind-1] = Contato{}
		}
	}


}
