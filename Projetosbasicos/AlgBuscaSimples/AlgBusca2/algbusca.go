package main

import "fmt"

func main() {
	size := 0
	fmt.Println("Quantos valores deseja inserir na array?")
	_, e := fmt.Scan(&size)
	if e != nil || size <= 0 { //Verificando se a entrada deve ser rejeitada ou nao. (Se há algum erro ou se é um número <= a 0)
		fmt.Println("Por favor, insira um número maior que zero.")
	}
	z := ""
	fmt.Println("Em que direção devo atribuir os valores? ('<' esquerda para direita ou '>' direita para esquerda)")
	fmt.Scan(&z) //Definindo a direção em que são atribui os valores à X.
	fmt.Println("Escolha os", size, "valores para a array:")
	x := make([]int, size)
	if z == "<" {
		for i := 0; i < size; i++ { //Looping para definir os valores de X (Caso <)
			fmt.Scan(&x[i])
		}
	} else if z == ">" {
		for i := len(x) - 1; i >= 0; i-- { //Looping invertido para definir os valores de x (Caso >)
			fmt.Scan(&x[i])
		}
	} else { //Caso não reconheça < ou >
		fmt.Println("Comando inválido.")
		return
	}
	y := 0
	fmt.Println("Qual valor você gostaria de localizar?")
	fmt.Scan(&y) //Definindo o valor de Y.(Número à ser localizado)
	iteracoes := 0
	leitura := ""
	fmt.Println("De qual direção deve partir a leitura?('<' para esquerda, e '>' para direita)")
	fmt.Scan(&leitura)
	find := false
	if leitura == "<" {
		for i, v := range x {
			iteracoes++
			if v == y {
				fmt.Println("O valor:", v, "se encontra no indice:", i)
				find = true //Define que encontrou o valor Y dentro da Array X.
				break
			}
		}
	} else if leitura == ">" {
		for i := len(x) - 1; i >= 0; i-- {
			if x[i] == y {
				iteracoes++
				fmt.Println("O valor:", y, "se encontra no indice:", i)
				find = true //Define que encontrou o valor Y dentro da Array X.
				break
			}
		}
	} else { //Caso não reconheça < ou >
		fmt.Println("Comando inválido.")
		return
	}
	if find != true { //Caso não encontre o valor definido por Y dentro da array X.
		fmt.Println("Não encontrei este valor dentro da array X.")
	}
	fmt.Println("iterações:", iteracoes)
}
