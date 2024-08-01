package main

import (
	"go-db/pkg/db"

	"github.com/abiosoft/ishell/v2"
)

func main() {
	mockDB := db.Connect()
	shell := ishell.New()

	shell.Println("Welcome! This is a simple key-value store.")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "create",
		Help: "create a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			value := c.Args[1]
			mockDB.Create(key, value)
			shell.Printf("Created %s = %s", key, value)
			shell.Println()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "get",
		Help: "get a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			shell.Printf("Got %s = %s", key, mockDB.Get(key))
			shell.Println()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "update",
		Help: "update a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			value := c.Args[1]
			mockDB.Update(key, value)
			shell.Printf("Updated %s = %s", key, value)
			shell.Println()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "delete",
		Help: "delete a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			mockDB.Delete(key)
			shell.Printf("Deleted %s", key)
			shell.Println()
		},
	})

	shell.Run()

}
