package main

import(
	"github.com/urfave/cli/v2"
	"fmt"
	"log"
	"os"
	
)

func main(){
	app := &cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that checks the given domain is down.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			}


		}
	}
}
	

