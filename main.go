package main

import (
	"fmt"

	"github.com/HenrySaldanha/Go.FlashCards/routes"
)

func main() {
	fmt.Println("Initializing Go Server")
	routes.HandleRequest()
}
