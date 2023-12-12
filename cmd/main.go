package main

import (
	"examTasks/store"
	"fmt"
)

func main() {
	data := "dataJson/first_exam.json"
	s, err := store.ReadJSonFile(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//s.Task1()
	//s.Task2()
	//s.Task3()
	//s.Task4()
	//s.Task5()
	//s.Task6()
	//s.Task7()
	//s.Task8()
	s.Task9()
	//s.Task10() 
	//s.Task11()
	//s.Task12()
	//s.Task13()
	//s.Task14()
	//s.Task15()

}
