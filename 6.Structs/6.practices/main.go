package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note successed")
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
