package main

import "fmt"

func main() {
	var contatos [5]Contato
	fmt.Println("Informe o nome e o email:")
	var nome, email string


	novoContato := Contato{Nome: nome, Email: email}
	contatos = addContato(novoContato, contatos)

	fmt.Println("Lista de contatos:")
	for {


	for _, contato := range Contato {
		fmt.Printf("Nome: %s, Email: %s\n", contato.Nome, contato.Email)
		fmt.Scan(&nome, &email)

	}
}}


type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) alteraEmail(novoEmail string) {
	c.Email = novoEmail
}

func addContato(a Contato, contatos [5]Contato) [5]Contato {
	for indice, contato := range contatos {
		if contato == (Contato{}) {
			fmt.Println("Achou um espaço vazio no índice", indice)
			contatos[indice] = a
			break
		}
	}
	return contatos
}

func revContato(a Contato, contatos [5]Contato) [5]Contato {
	for indice, contato := range contatos {
		if contato == a {
			fmt.Println("Achou um contato para remover no índice", indice)
			contatos[indice] = Contato{}
			break
		}
	}
	return contatos
}
