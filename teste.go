package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Aluno struct{
	nome string
	nota float64
}

type ListaAlunos struct{
	Alunos []Aluno
}

//METODOS ALUNO
func (a *Aluno) criarAluno(nome string, nota float64){
	a.nome = nome
	a.nota = nota
}

//------------------------------SYSTEM-------------------------------------

// Função para limpar o terminal
func limparTela() {
	// Verifica o sistema operacional
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func pausar(){
	fmt.Print("Pressione Enter para continuar...")
	var buf [1]byte
	os.Stdin.Read(buf[:]) // Aguarda qualquer tecla ser pressionada
}

//------------------------------ORDENAÇÃO-------------------------------------
func insertionSortNome(alunoLista *[]Aluno){
	// fmt.Println(*alunoLista)
	// fmt.Println((*alunoLista)[0])

	// fmt.Println(len(*alunoLista))

	// time.Sleep(time.Second * 2)
	 i := 0
	 j := 0
	 aux := Aluno{}
	 for i=0; i < len(*alunoLista); i++{
		aux = (*alunoLista)[i]
		for j = i; (j>0) && aux.nome < (*alunoLista)[j-1].nome; j--{
			(*alunoLista)[j] = (*alunoLista)[j-1]
		}
		(*alunoLista)[j] = aux
	 }
}

func insertionSortNota(alunoLista *[]Aluno){
	// fmt.Println(*alunoLista)
	// fmt.Println((*alunoLista)[0])

	// fmt.Println(len(*alunoLista))

	// time.Sleep(time.Second * 2)
	 i := 0
	 j := 0
	 aux := Aluno{}
	 for i=0; i < len(*alunoLista); i++{
		aux = (*alunoLista)[i]
		for j = i; (j>0) && aux.nota > (*alunoLista)[j-1].nota; j--{
			(*alunoLista)[j] = (*alunoLista)[j-1]
		}
		(*alunoLista)[j] = aux
	 }
}

//------------------------------PROGRAMA-------------------------------------
func imprimeMenu(){
	limparTela()

	fmt.Println("-------MENU--------")
	fmt.Println("0 - Para finalizar o programa")
	fmt.Println("1 - Adicionar aluno")
	fmt.Println("2 - listar alunos da turma")
	fmt.Println("3 - Deletar aluno")
	fmt.Println("4 - Atualizar aluno")

	fmt.Printf("\n Digite a opção desejada: ")
}

func adicionarAluno(alunoLista *[]Aluno, tamanhoLista *int){
	limparTela()

	var nome string
	var nota float64
	
	fmt.Println("-------ADICIONAR ALUNOS--------")
	fmt.Println("-Digite -1 a qualquer momento para voltar ao MENU-")


	fmt.Printf("\nDigite o NOME do aluno:")
	fmt.Scan(&nome)
	if(nome == "-1"){
		return
	}
	fmt.Printf("\nDigite a NOTA do aluno:")
	fmt.Scan(&nota)

	for {
		if(nota == -1){
			return
		}
		
		if nota >=0.00 && nota <= 10.00 {
			break
		}

		fmt.Println("Digite uma nota de 0.0 a 10.0: ")
		fmt.Scan(&nota)
	}



	novoAluno := Aluno{}//cria uma variavel do tipo Aluno
	novoAluno.criarAluno(nome, nota)//chama o metodo de criação de aluno
	*alunoLista = append(*alunoLista, novoAluno)//adiciona o aluno criado a lista

	//verifica se o tamanho da lista aumentou em relação ao inicio da função
	//se sim, aluno foi cadastrado na lista
	if *tamanhoLista != len(*alunoLista){
		fmt.Println("Aluno cadastrado com sucesso")
		*tamanhoLista++
		}else{
			fmt.Println("Ocorreu um erro")
		}
		
	time.Sleep(time.Millisecond * 1000)
	adicionarAluno(alunoLista, tamanhoLista)
	
}

func listarAlunos(alunoLista []Aluno){
	limparTela()
	fmt.Println("-------LISTA DE ALUNOS--------")
	fmt.Println("-Digite -1 para voltar ao MENU-")
	fmt.Println("-Digite 1 para ordenar por ordem alfabetica-")
	fmt.Println("-Digite 2 para ordenar pelas notas em ordem decrescente-\n")
	var opcao int8

	if len(alunoLista) == 0 {
		fmt.Println("Nenhum aluno cadastrado")
		return
	}

	for index, value := range alunoLista{
		fmt.Printf("%d - %s - nota: %.2f\n", index+1, value.nome, value.nota)
	}


	fmt.Scanf("%d", &opcao)
	fmt.Scan(&opcao)
	// pausar()

	//usando AUX para nao modificar a lista original na ordenação
	aux := []Aluno{}
	aux = append(aux, alunoLista...)

	if(opcao == 1){
		insertionSortNome(&aux)
		listarAlunos(aux)
	}else if(opcao == 2){
		insertionSortNota(&aux)
		listarAlunos(aux)
	}

}

func pesquisarAluno(alunoLista []Aluno, nome string) (Aluno, bool) {
	for _, value := range alunoLista{
		if strings.ToLower(value.nome) == strings.ToLower(nome){
			return value, true
		}
	}
	return Aluno{}, false
}

func deletarAluno(alunoLista *[]Aluno, tamanhoLista *int){
	limparTela()

	if len(*alunoLista) == 0 {
		fmt.Println("Nenhum aluno cadastrado")
		time.Sleep(time.Millisecond * 2000)
		return
	}

	var nome string
	fmt.Println("-------LISTA DE ALUNOS--------")
	fmt.Println("-Digite -1 para voltar ao MENU-")

	fmt.Println("Digite o nome do aluno a ser deletado")
	fmt.Scanln(&nome)
	fmt.Scanf("%s", &nome)

	if(nome == "-1"){
		return
	}
	
	//verifica se aluno existe na lista
	alunoPesquisado, isEncontrado := pesquisarAluno(*alunoLista, nome)
	if isEncontrado == false{
		fmt.Println("Aluno não encontrado")
		time.Sleep(time.Millisecond * 2000)
	}else{
		aux := []Aluno{}
		for _, value := range *alunoLista{
			
			if value != alunoPesquisado{
				aux = append(aux, value)
			}
		}

		if *tamanhoLista >= len(aux){
			fmt.Println("Aluno deletado com sucesso!")
			*alunoLista = aux
			*tamanhoLista--
		}
	}


	time.Sleep(time.Millisecond * 1000)
	
	deletarAluno(alunoLista, tamanhoLista)
}

func atualizarAluno(alunoLista *[]Aluno, tamanhoLista int){
	limparTela()

	if tamanhoLista == 0 {
		fmt.Println("Nenhum aluno cadastrado")
		time.Sleep(time.Millisecond * 2000)
		return
	}

	var nome string
	var nota float64

	fmt.Println("-------ATUALIZAR ALUNO--------")
	fmt.Println("-Digite -1 para voltar ao MENU-")

	fmt.Print("Digite o nome do aluno a ser ATUALIZADO: ")
	fmt.Scanln(&nome)
	fmt.Scanf("%s", &nome)

	if(nome == "-1"){
		return
	}

	//verifica se aluno existe na lista
	alunoPesquisado, isEncontrado := pesquisarAluno(*alunoLista, nome)
	if isEncontrado == false{
		fmt.Println("Aluno não encontrado")
		time.Sleep(time.Millisecond * 2000)
		atualizarAluno(alunoLista, tamanhoLista)

	}else{
		fmt.Println("Aluno encontrado!")

		aux := []Aluno{}
		for _, value := range *alunoLista{

			if value != alunoPesquisado{
				aux = append(aux, value)
			}else{
				fmt.Print("Digite o nome atualizado do aluno: ")
				fmt.Scanln(&nome)
				fmt.Scanf("%s", &nome)
				fmt.Print("Digite a nota atualizada do aluno: ")
				fmt.Scanln(&nota)
				fmt.Scanf("%f", &nota)

				for {
					if(nota == -1){
						return
					}
					
					if nota >=0.00 && nota <= 10.00 {
						break
					}
			
					fmt.Println("Digite uma nota de 0.0 a 10.0: ")
					fmt.Scan(&nota)
				}
				
				alunoAtualizado := Aluno{nome: nome, nota: nota}
				aux = append(aux, alunoAtualizado)



			}
		}

		if tamanhoLista == len(aux){
			fmt.Println("Aluno atualizado com sucesso")
			time.Sleep(time.Millisecond * 2000)
			*alunoLista = aux
		}
	}

}

func main(){
	//criando lista do tipo aluno
	alunosLista := []Aluno{}
	
	//povoando a lista
	a1 := Aluno{nome: "Emily", nota: 10}
	a2 := Aluno{nome: "Danylo", nota: 7}

	alunosLista = append(alunosLista, a1, a2)
	//controle do tamanho da lista
	tamanhoLista := len(alunosLista)
	var ptrTamanhoLista *int = &tamanhoLista

	//estrutura de escolha
	var opcao int
	for ok := true; ok; ok = (opcao  != 0){
		imprimeMenu()
		fmt.Scan(&opcao)

		switch opcao{
			case 0: break
			case 1: adicionarAluno(&alunosLista, ptrTamanhoLista)
			case 2: listarAlunos(alunosLista)
			case 3: deletarAluno(&alunosLista, ptrTamanhoLista)
			case 4: atualizarAluno(&alunosLista, *ptrTamanhoLista)
			default: fmt.Println("Valor invalido")
		}
	}	
}