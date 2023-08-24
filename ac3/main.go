package main

import (
	"ac3/Contato"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contatos [5]Contato.Contato
	var nome, email, opcao, novoEmail string
	var i int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("(1) para adicionar, (2) para remover (3) indice lista (4)altera email, Qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			Contato.AdicionaContato(&contatos, Contato.Contato{Nome: nome, Email: email})

		case "2":
			Contato.ExcluiContato(&contatos)

		case "3":
			fmt.Println("Contatos:")
			for i, contato := range contatos {
				if contato.Nome == "" {
					fmt.Println("não tem cadastro no indice ", i)
					continue
			}else{
				fmt.Println("Índice:",i)
				fmt.Println("Nome:",contato.Nome)
				fmt.Println("Email:", contato.Email)
				fmt.Println("___________________________________________-")
			}
		}
		case "4":
			fmt.Println("Informe o índice do contato a ser editado: ")
			fmt.Scanln(&i)
			fmt.Println("Informe o novo e-mail: ")
			fmt.Scanln(&novoEmail)
			Contato.AlterarEmail(&contatos, i, novoEmail)



		default:
			fmt.Println("PROGRAMA ENCERRADO")
			return
		}
	}

	fmt.Println(contatos)

}
