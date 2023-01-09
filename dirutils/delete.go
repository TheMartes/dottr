package dirutils

import (
	"log"
	"os"
)

func RemoveFolder(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}
