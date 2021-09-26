package main

import (
	"fmt"
)

func main() {
	fmt.Println("\033[36m" + "Witaj w panelu administracyjnym mLib" + "\033[0m")
	menu()
}

func menu() {
	fmt.Println("Wybierz numer polecenia")
	fmt.Println("\033[33m" + "0 - Generuj losowe liczby (narzędzie pomocnicze)" + "\033[0m")
	fmt.Println("\033[33m" + "1 - Wyświetl bazę danych" + "\033[0m")
	fmt.Println("\033[33m" + "2 - Przeładuj bazę danych)" + "\033[0m")
	fmt.Println("\033[33m" + "3 - Usuń bazę danych)" + "\033[0m")
	var command int
	fmt.Scan(&command)
	if command == 0 {
		randomize()
	}
	if command == 1 {
		get()
	}
	if command == 2 {
		delete()
		post()
	}
	if command == 3 {
		delete()
	}
	fmt.Println("\nCoś jeszcze?")
	menu()
}
