package fileutils

import (
	"log"
	"os"
)

func RemoveFile(path string) {
    err := os.Remove(path)
    if err != nil {
        log.Fatal(err)
    }
}
