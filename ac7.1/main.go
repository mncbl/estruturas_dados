package main

import "fmt"

type no struct {
	number int
	next   *no
	prev   *no
}

type doublylinkedlist struct {
	comeco *no
	len    int
}

func (dll *doublylinkedlist) remove(posicao int) {
	if dll.len >= posicao {
		var p *no = dll.comeco
		for i := 0; i < posicao-1; i++ {
			p = p.next
		}
		p.next = p.next.next
		if p.next != nil {
			p.next.prev = p
		}
		dll.len--
	} else {
		fmt.Println("Não foi encontrado")
	}
}

func (dll *doublylinkedlist) add(num, posicao int) {
	var temp = &no{number: num}
	var p *no = dll.comeco
	for i := 0; i < posicao-1; i++ {
		p = p.next
	}
	temp.next = p.next
	if p.next != nil {
		p.next.prev = temp
	}
	temp.prev = p
	p.next = temp
	dll.len++
}

func (dll *doublylinkedlist) MostraLista() {
	no := dll.comeco.next
	for no != nil {
		fmt.Println(no.number)
		no = no.next
	}
}

func (dll *doublylinkedlist) Busca(numero int) *no {
	no := dll.comeco.next
	for no != nil {
		if no.number == numero {
			return no
		}
		no = no.next
	}
	return nil
}

func main() {
	dll := &doublylinkedlist{comeco: &no{}}

	dll.add(1, 0)
	dll.add(9, 1)
	dll.add(10, 2)
	dll.MostraLista()

	fmt.Println("-------------------------------------------")

	dll.remove(2)
	fmt.Println("-------------------------------------------")
	dll.MostraLista()
	fmt.Println("-------------------------------------------")

	resultado := dll.Busca(9)
	indice:= (dll.len -1)
	if resultado != nil {
		fmt.Println("Valor encontrado:", resultado.number, indice)
	} else {
		fmt.Println("Valor não encontrado")
	}
}
