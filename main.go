package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Student struct{
	name string
	grade float64
}

type StudentList struct{
	Students []Student
}

//METODOS ALUNO
func (a *Student) CreateStudent(name string, grade float64){
	a.name = name
	a.grade = grade
}

//------------------------------SYSTEM-------------------------------------

// Função para limpar o terminal
func clearScreen() {
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

func pause(){
	fmt.Print("Pressione Enter para continuar...")
	var buf [1]byte
	os.Stdin.Read(buf[:]) // Aguarda qualquer tecla ser pressionada
}

//------------------------------ORDENAÇÃO-------------------------------------
func sortStudentsByName(studentList *[]Student){
	 i := 0
	 j := 0
	 aux := Student{}
	 for i=0; i < len(*studentList); i++{
		aux = (*studentList)[i]
		for j = i; (j>0) && aux.name < (*studentList)[j-1].name; j--{
			(*studentList)[j] = (*studentList)[j-1]
		}
		(*studentList)[j] = aux
	 }
}

func sortStudentsByGrade(studentList *[]Student){
	 i := 0
	 j := 0
	 aux := Student{}
	 for i=0; i < len(*studentList); i++{
		aux = (*studentList)[i]
		for j = i; (j>0) && aux.grade > (*studentList)[j-1].grade; j--{
			(*studentList)[j] = (*studentList)[j-1]
		}
		(*studentList)[j] = aux
	 }
}

//------------------------------PROGRAMA-------------------------------------
func printMenu(){
	clearScreen()

	fmt.Println("-------MENU--------")
	fmt.Println("0 - Para finalizar o programa")
	fmt.Println("1 - Adicionar aluno")
	fmt.Println("2 - listar alunos da turma")
	fmt.Println("3 - Deletar aluno")
	fmt.Println("4 - Atualizar aluno")

	fmt.Printf("\n Digite a opção desejada: ")
}

func addStudent(studentList *[]Student, sizeStudentList *int){
	clearScreen()

	var name string
	var grade float64
	
	fmt.Println("-------ADICIONAR ALUNOS--------")
	fmt.Println("-Digite -1 a qualquer momento para voltar ao MENU-")


	fmt.Printf("\nDigite o NOME do aluno: ")
	fmt.Scan(&name)
	if (name == "-1"){
		return
	}

	studentFound, isFound := getStudentByName(*studentList, name)
	if isFound {
		fmt.Printf("Já existe um aluno cadastrado com esse nome!\n")
		fmt.Printf("Aluno: %s - Nota: %0.2f", studentFound.name, studentFound.grade)

		//chama a função de volta
		time.Sleep(time.Millisecond * 1000)
		addStudent(studentList, sizeStudentList)
	}

	
	fmt.Printf("\nDigite a NOTA (0.0 a 10.0) do aluno: ")
	fmt.Scan(&grade)

	for {
		if(grade == -1){
			return
		}
		
		if grade >=0.00 && grade <= 10.00 {
			break
		}

		fmt.Println("Digite uma nota de 0.0 a 10.0: ")
		fmt.Scan(&grade)
	}



	newStudent := Student{}//cria uma variavel do tipo Aluno
	newStudent.CreateStudent(name, grade)//chama o metodo de criação de aluno
	*studentList = append(*studentList, newStudent)//adiciona o aluno criado a lista

	//verifica se o tamanho da lista aumentou em relação ao inicio da função
	//se sim, aluno foi cadastrado na lista
	if *sizeStudentList != len(*studentList){
		fmt.Println("Aluno cadastrado com sucesso!")
		*sizeStudentList++
		}else{
			fmt.Println("Ocorreu um erro!")
		}
		
	time.Sleep(time.Millisecond * 1000)
	addStudent(studentList, sizeStudentList)
	
}

func displayStudents(studentList []Student){
	clearScreen()
	fmt.Println("-------LISTA DE ALUNOS--------")
	fmt.Println("-Digite -1 para voltar ao MENU-")
	fmt.Println("-Digite 1 para ordenar por ordem alfabetica-")
	fmt.Printf("-Digite 2 para ordenar pelas notas em ordem decrescente-\n\n")
	var option int8

	if len(studentList) == 0 {
		fmt.Println("Nenhum aluno cadastrado")
		return
	}

	for index, value := range studentList{
		fmt.Printf("%d - %s - nota: %.2f\n", index+1, value.name, value.grade)
	}

	fmt.Printf("\n Digite a opção desejada: ")
	fmt.Scanf("%d", &option)
	fmt.Scan(&option)
	// pausar()

	//usando AUX para nao modificar a lista original na ordenação
	aux := []Student{}
	aux = append(aux, studentList...)

	if(option == 1){
		sortStudentsByName(&aux)
		displayStudents(aux)
	}else if(option == 2){
		sortStudentsByGrade(&aux)
		displayStudents(aux)
	}

}

func getStudentByName(studentList []Student, name string) (Student, bool) {
	for _, value := range studentList{
		if strings.EqualFold(value.name, name) {
			return value, true
		}
	}
	return Student{}, false
}

func deleteStudent(studentList *[]Student, sizeStudentList *int){
	clearScreen()

	if len(*studentList) == 0 {
		fmt.Println("Nenhum aluno cadastrado")
		time.Sleep(time.Millisecond * 2000)
		return
	}

	var name string
	fmt.Println("-------LISTA DE ALUNOS--------")
	fmt.Println("-Digite -1 para voltar ao MENU-")

	fmt.Printf("Digite o nome do aluno a ser deletado: ")
	fmt.Scanln(&name)
	fmt.Scanf("%s", &name)

	if(name == "-1"){
		return
	}
	
	//verifica se aluno existe na lista
	studentFound, isFound := getStudentByName(*studentList, name)
	if !isFound {
		fmt.Println("\nAluno não encontrado!")
		time.Sleep(time.Millisecond * 2000)
	}else{
		aux := []Student{}
		for _, value := range *studentList{
			
			if value != studentFound{
				aux = append(aux, value)
			}
		}

		if *sizeStudentList >= len(aux){
			fmt.Println("\nAluno deletado com sucesso!")
			*studentList = aux
			*sizeStudentList--
		}
	}


	time.Sleep(time.Millisecond * 1000)
	
	deleteStudent(studentList, sizeStudentList)
}

func updateStudent(studentList *[]Student, sizeStudentList int){
	clearScreen()

	if sizeStudentList == 0 {
		fmt.Println("Nenhum aluno cadastrado")
		time.Sleep(time.Millisecond * 2000)
		return
	}

	var name string
	var grade float64

	fmt.Println("-------ATUALIZAR ALUNO--------")
	fmt.Println("-Digite -1 para voltar ao MENU-")

	fmt.Print("Digite o nome do aluno a ser ATUALIZADO: ")
	fmt.Scanln(&name)
	fmt.Scanf("%s", &name)

	if(name == "-1"){
		return
	}

	//verifica se aluno existe na lista
	studentFound, isFound := getStudentByName(*studentList, name)
	if !isFound {
		fmt.Println("Aluno não encontrado")
		time.Sleep(time.Millisecond * 2000)
	}else{
		fmt.Printf("Aluno encontrado!\n\n")

		aux := []Student{}
		for _, value := range *studentList{

			if value != studentFound{
				aux = append(aux, value)
			}else{
				fmt.Print("Digite o NOVO NOME do aluno: ")
				fmt.Scanln(&name)
				fmt.Scanf("%s", &name)
				fmt.Print("Digite a NOVA NOTA (0.0 a 10.0) do aluno: ")
				fmt.Scanln(&grade)
				fmt.Scanf("%f", &grade)

				for {
					if(grade == -1){
						return
					}
					
					if grade >=0.00 && grade <= 10.00 {
						break
					}
			
					fmt.Println("Digite uma nota de 0.0 a 10.0: ")
					fmt.Scan(&grade)
				}
				
				updatedStudent := Student{name: name, grade: grade}
				aux = append(aux, updatedStudent)



			}
		}

		if sizeStudentList == len(aux){
			fmt.Printf("Aluno atualizado com sucesso!\n\n")
			time.Sleep(time.Millisecond * 2000)
			*studentList = aux
		}
	}

}

func main(){
	//criando lista do tipo aluno
	studentList := []Student{}
	
	//povoando a lista
	a1 := Student{name: "Emily", grade: 10}
	a2 := Student{name: "Danylo", grade: 7}

	studentList = append(studentList, a1, a2)
	//controle do tamanho da lista
	sizeStudentList := len(studentList)
	var pointerSizeStudentList *int = &sizeStudentList

	//estrutura de escolha
	var option int
	for ok := true; ok; ok = (option  != 0){
		printMenu()
		fmt.Scan(&option)

		switch option{
			case 0: break
			case 1: addStudent(&studentList, pointerSizeStudentList)
			case 2: displayStudents(studentList)
			case 3: deleteStudent(&studentList, pointerSizeStudentList)
			case 4: updateStudent(&studentList, *pointerSizeStudentList)
			default: fmt.Println("Valor invalido")
		}
	}	
}