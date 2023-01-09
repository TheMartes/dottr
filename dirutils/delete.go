package dirutils

import (
	"log"
	"os"
)

func RemoveFolder(path string) {
    err := os.Remove(path)
    if err != nil {
        log.Fatal(err)
    }
}
