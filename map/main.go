package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"bvb":    "yellow",
	// 	"united": "red",
	// }
	// var colors map[string]string
	colors := make(map[string]string)
	colors["bvb"] = "yellow"
	colors["united"] = "red"
	colors["rma"] = "white"
	// fmt.Println(colors["bvb"])
	// fmt.Println(colors["united"])
	fmt.Println(colors)
	// delete(colors, "united")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for team, color := range m {
		fmt.Printf("Color of %v is %v\n", team, color)
	}
}
