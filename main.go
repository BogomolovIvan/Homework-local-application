package main

import (
	"bufio"
	"fmt"
	"os"
)

type Note struct {
	Number  int
	Name    string
	Surname string
	Note    string
}

func main() {
	count := 1
	allNotes := make([]Note, 0, 10)
	note := setNote(count)

	allNotes = append(allNotes, note)

	for {
		res := ""
		fmt.Println("\nЧто делать дальше?")
		fmt.Println(
			"\"y\" - продолжить; " +
				" \"n\" - завершить; " +
				" \"p\" - показать все")
		fmt.Scan(&res)

		switch res {
		case "y":
			count++
			note := setNote(count)
			allNotes = append(allNotes, note)
		case "p":
			fmt.Println("Вы записали ваши заметки: ")

			for _, n := range allNotes {
				printNote(n)
			}
		case "n":
			return
		default:
			continue
		}
	}
}

func printNote(n Note) {
	fmt.Printf(
		"Запись: №%d  "+
			"\nИмя: %s; "+
			"\nФамилия: %s; "+
			"\nЗаметка: %s \n\n", n.Number, n.Name, n.Surname, n.Note)
}

func setNote(number int) Note {
	var note Note

	fmt.Println("Введи своё имя")
	fmt.Scan(&note.Name)
	fmt.Println("Введи свою фамилию")
	fmt.Scan(&note.Surname)
	fmt.Println("Напиши заметку")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		note.Note = scanner.Text()
	}

	note.Number = number

	return note
}
