package envy

import "fmt"

func Run() {
	fmt.Println(Version)

	port := Load("X")
	fmt.Println("VALUE", port)
	fmt.Println("Leido")
}
