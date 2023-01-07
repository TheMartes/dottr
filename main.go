package main

import (
	"log"
	"os"

	dottrInit "github.com/TheMartes/dottr/init"
	"github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name: "dottr",
        Version: "TBD",
        Commands: []*cli.Command{
            {
                Name:    "init",
                Usage:   "init dottr into your dotfiles based on the provided config",
                Action: func(cCtx *cli.Context) error {
                    dottrInit.InitDottr()
                    return nil
                },
            },
            {
                Name:    "commit",
                Usage:   "commit dottr to your dotfiles and reverse the symlinks",
                Action: func(cCtx *cli.Context) error {
                    dottrInit.Commit()
                    return nil
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
