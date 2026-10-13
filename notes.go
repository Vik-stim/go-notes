package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("--- ЗАМЕТКИ ---")
		fmt.Println("1 — добавить заметку")
		fmt.Println("2 — показать все заметки")
		fmt.Println("3 — выйти")
		fmt.Print("Выбери: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Print("Напиши текст для заметки: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			
			oldData, err := os.ReadFile("notes.txt")
			if err != nil {
				oldData = []byte("")
			}
			oldText := string(oldData)
			
			fullText := oldText + "\n" + text
			err = os.WriteFile("notes.txt", []byte(fullText), 0644)
			if err != nil {
				fmt.Println("Ошибка записи:", err)
				return
			}
			fmt.Println("Заметка сохранена!")
		} else if choice == 2 {
			data, err := os.ReadFile("notes.txt")
			if err != nil {
				fmt.Println("Заметок пока нет")
			} else {
				fmt.Println("--- Все заметки ---")
				fmt.Println(string(data))
			}
		} else if choice == 3 {
			fmt.Println("Пока!")
			break
		}
	}
	
}