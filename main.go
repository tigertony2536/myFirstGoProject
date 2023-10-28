package main

import "fmt"

// import "fmt"

// func changeString(text *string) {
// 	*text = "Black Jack"
// }

func main() {
	var card string = "Aces of Spades"
	card := "Aces of Spades"
	changeString(&card)
	fmt.Println(&card)
	test()
}
