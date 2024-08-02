package main

import (
	"go-example/pkg/kvstore"

	"github.com/abiosoft/ishell/v2"
)

func main() {
	kv := kvstore.CreateKV()
	shell := ishell.New()

	shell.Println("Welcome! This is a simple key-value store.")

	shell.AddCmd(&ishell.Cmd{
		Name: "set",
		Help: "set a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			value := c.Args[1]
			// TODO: Implement the set command
			shell.Printf("Set %s = %s", key, value)
			shell.Println()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "get",
		Help: "get a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			// TODO: Implement the get command
			shell.Printf("Got %s = %s", key, kv.Get(key))
			shell.Println()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "delete",
		Help: "delete a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			// TODO: Implement the delete command
			shell.Printf("Deleted %s", key)
			shell.Println()
		},
	})

	shell.Run()

}
