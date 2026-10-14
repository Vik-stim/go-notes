package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func combineNotes(oldText, newText string) string {
	return oldText + "\n" + newText
}

func addNote(scanner *bufio.Scanner) {
	fmt.Print("Напиши текст для заметки: ")
	scanner.Scan() // больше не создаём новый, а пользуемся переданным
	text := scanner.Text()

	oldData, err := os.ReadFile("notes.txt")
	if err != nil {
		oldData = []byte("")
	}
	oldText := string(oldData)

	fullText := combineNotes(oldText, text)
	err = os.WriteFile("notes.txt", []byte(fullText), 0644)
	if err != nil {
		fmt.Println("Ошибка записи:", err)
		return
	}
	fmt.Println("Заметка сохранена!")
}

func showNotes() {
	data, err := os.ReadFile("notes.txt")
	if err != nil {
		fmt.Println("Заметок пока нет")
	} else {
		fmt.Println("--- Все заметки ---")
		fmt.Println(string(data))
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("--- ЗАМЕТКИ ---")
		fmt.Println("1 — добавить заметку")
		fmt.Println("2 — показать все заметки")
		fmt.Println("3 — выйти")
		fmt.Print("Выбери: ")

		scanner.Scan()                     // читаем строку ввода
		input := scanner.Text()            // например, "1"
		choice, err := strconv.Atoi(input) // превращаем в число
		if err != nil {
			fmt.Println("Введи число от 1 до 3")
			continue // пропускаем круг цикла, меню напечатается заново
		}

		if choice == 1 {
			addNote(scanner)
		} else if choice == 2 {
			showNotes()
		} else if choice == 3 {
			fmt.Println("Пока!")
			break
		}
	}

}
