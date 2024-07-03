package main

import (
	"fmt"

	"github.com/husseinamine/whatsapp"
)

var TOKEN = "SECRET"
var BUSINESS_ID = "Test"

func main() {
	client := whatsapp.NewClient(BUSINESS_ID, TOKEN)

	factory := whatsapp.NewMessageFactory()

	factory.WithTo("+000000000")
	message := factory.WithTemplate("hello_world", "en_US")

	message.Save()

	S, E := client.SendMessage(message.Message)

	fmt.Println(S)
	fmt.Println(E)
}
