package main

import "fmt"

func main() {
	// color := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "4bf745",
	// }
	// var color map[string]string

	color := make(map[string]string)
	color["red"] = "#ff0000"
	delete(color, "red")
	fmt.Println(color)
}
