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
			kv.Set(key, value)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "get",
		Help: "get a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			value := kv.Get(key)
			shell.Printf(value)
			shell.Println()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "unset",
		Help: "unset a kv pair",
		Func: func(c *ishell.Context) {
			key := c.Args[0]
			kv.Unset(key)
		},
	})

	shell.Run()

}
