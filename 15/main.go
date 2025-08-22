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
// 1) Ошибка если строка v длиной меньше 100
// 2) Утечка памяти, т.к. v не будет освобождена сборщиком мусора
// Как это исправить:
// 1) Сделать проверку длины v перед justString = v[:100]
// 2) justString = string([]byte(v[:100])) вместо justString = v[:100]

// Что происходит с переменной justString?
// Переменная инициализируется и получает первые 100 символов v
