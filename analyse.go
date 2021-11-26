package main

import "fmt"

func analyse() {
	for i := range categories {
		var n int = 0
		for j := range books {
			if categories[books[j].category].name == categories[i].name {
				n++
			}
		}
		fmt.Println(categories[i].name+":", n)
	}
}
