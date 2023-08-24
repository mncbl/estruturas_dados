package Contato // O nome do pacote deve ser maiúsculo para ser acessível de outros pacotes



func AdicionaContato( a*[5]Contato,c Contato)  {
	for ind, contato := range a {
		if (contato == Contato{}) {
			a[ind] = c
			break
		}
	}


}

