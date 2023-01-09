package main

import (
	"log"
	"os"

	"github.com/TheMartes/dottr/preview"
	"github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name: "dottr",
        Version: "TBD",
        Commands: []*cli.Command{
            {
                Name:    "preview",
                Usage:   "preview dottr structure based on the provided config",
                Action: func(cCtx *cli.Context) error {
                    preview.Init()
                    return nil
                },
            },
            {
                Name:    "commit",
                Usage:   "commit to dottr. (should be run after you saw preview)",
                Action: func(cCtx *cli.Context) error {
                    preview.Apply()
                    return nil
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
