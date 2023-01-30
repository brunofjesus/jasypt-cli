package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "jasypt-cli",
		Usage: "encrypts and decrypts jasypt strings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "algorithm",
				Aliases:  []string{"algo", "a"},
				Value:    defaultAlgorithm,
				Usage:    "use specified algorithm",
				Required: false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e", "enc"},
				Usage:   "encrypt [key] [text_to_encrypt]",
				Action: func(cCtx *cli.Context) error {
					password, text, algorithm := getCmdOptions(cCtx)

					result, err := Encrypt(text, password, algorithm)
					if err != nil {
						return err
					}
					fmt.Print(result)
					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d", "dec"},
				Usage:   "decrypt [key] [text_to_decrypt]",
				Action: func(cCtx *cli.Context) error {
					password, text, algorithm := getCmdOptions(cCtx)

					result, err := Decrypt(text, password, algorithm)
					if err != nil {
						return err
					}
					fmt.Print(result)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCmdOptions(cCtx *cli.Context) (string, string, string) {
	password := cCtx.Args().Get(0)
	text := cCtx.Args().Get(1)
	algorithm := cCtx.String("algorithm")

	return password, text, algorithm
}
