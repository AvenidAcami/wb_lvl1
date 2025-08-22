package main

import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(int) string {
	return "hsdfhushufsdfgjhgujiohgfujhifujihiujgfiujhujijfjoihfgihujfujighujiopfgujipofghujigopfhвввввв"
}

// К каким негативным последствиям он может привести:
// Ошибка если строка v длиной меньше 100
// Как это исправить:
// Сделать проверку длины v перед justString = v[:100]

// Что происходит с переменной justString?
// Переменная инициализируется и ей присваивается часть v
