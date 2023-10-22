package main

import "fmt"

type Animal struct {
	Nome string
	Tipo string
}

type Lista struct {
	Comeco *No
	Fim    *No
}

type No struct {
	Valor Animal
	Prox  *No
}

func (l *Lista) Add(animal Animal) {
	no := No{Valor: animal}
	if l.Comeco == nil {
		l.Comeco = &no

	}
	if l.Fim != nil{

		l.Fim.Prox = &no
	}
	l.Fim = &no
}

func (l *Lista) MostraLista() {
	no := l.Comeco
	for no != nil {
		fmt.Println(no.Valor.Nome)
		no = no.Prox
	}
}

func (l *Lista) remove(nome string){
	if l.Comeco.Valor.Nome == nome {
		l.Comeco = l.Comeco.Prox
		return
	}
	prev := l.Comeco
	no := l.Comeco.Prox
	for no != nil {
		if no.Valor.Nome == nome{
			prev.Prox = no.Prox
			break
		}
		prev = no
		no = no.Prox
	}
	if l.Fim == no{
		l.Fim = prev
	}

}

func main() {
	lista := Lista{}
	cachorro := Animal{"Caramelo", "Mamífero"}
	pardal := Animal{"Piu", "Passarinho"}
	cavalo := Animal{"pocoto", "Mamifero"}
	porco := Animal{"rodolfo", "Mamifero"}
	lista.Add(cavalo)
	lista.Add(porco)
	lista.Add(cachorro)
	lista.Add(pardal)
	fmt.Println("-------------------------------------------")
	lista.MostraLista()
	fmt.Println("-------------------------------------------")
	lista.remove("Piu")
	lista.MostraLista()
	fmt.Println("-------------------------------------------")
	lista.remove("rodolfo")
	lista.MostraLista()

}
