package Contato

func  AlterarEmail(a *[5]Contato, i int ,novoEmail string) {
	for i := 0; i < len(a);{
		a[i].Email = novoEmail
		break
	}


}