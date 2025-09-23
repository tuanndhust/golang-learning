package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface { //dam bao 1 struct co mot so method nhat dinh
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface { //embedded interfaces
	Saver
	// displayer
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("any value")

	title, content := getNoteData()
	todoText := getUserInput("todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	//--------------------------
	// todo.Display()
	// err = saveData(todo)
	outputData(todo)

	// if err != nil {
	// 	return
	// }

	//----------------------------
	// userNote.Display()
	// err = saveData(userNote)
	outputData(userNote)

	// if err != nil {
	// 	return
	// }

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	// if err != nil {
	// 	return "", "", err
	// }
	content := getUserInput("Note content:")
	// if err != nil {
	// 	return "", "", err
	// }
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // dung doc string khi xuong dong (delim)
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") //windows xuong dong se gom ca 2 ky tu "\n\r"
	return text
}

//todo

// func getTodoData() string {
// 	return getUserInput("Todo text: ")
// }

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value interface{}) { //any value allowed
	//cach 1
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
	}

	floatVal, ok := value.(int)
	if ok {
		fmt.Println("Float: ", floatVal)
	}

	//cach 2

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float64: ", value)
	// case string:
	// 	fmt.Println(value)
	// default:
	// 	//...
	// }
	// fmt.Println(value)
}
