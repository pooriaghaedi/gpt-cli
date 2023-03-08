package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	secret := check()
    app := &cli.App{
        Action: func(c *cli.Context) error {
            args := strings.Join(c.Args().Slice(), " ")
            fmt.Println(args)
            gpt(secret, args)

            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}


func check() string {
    _, err := os.Open(os.ExpandEnv("$HOME/.go-chatgpt/credentials"))
    if err != nil {
        fmt.Println(err)
		initial()
    }

	creds, err := os.ReadFile(os.ExpandEnv("$HOME/.go-chatgpt/credentials")) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    secret := string(creds) 
	return secret
}


