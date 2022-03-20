package main

import (
	"fmt"

	"github.com/czewski/tg-newsletter/pkg/constants"
	"github.com/czewski/tg-newsletter/pkg/server"
)

func main() {
	fmt.Println("hi")
	token := constants.ReadKey("BotKey")
	fmt.Println(token)
	server.CreateServer()
}
