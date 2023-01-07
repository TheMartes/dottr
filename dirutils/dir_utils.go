package dirutils

import (
	"log"
	"os"
)

func GetHomeDirectory() string {
    dirname, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

    return dirname
}
