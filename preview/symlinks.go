package preview

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/TheMartes/dottr/config"
	"github.com/TheMartes/dottr/dirutils"
	"github.com/avakarev/go-symlink"
)

type SymlinkWay string

const (
	SymlinkToRepo   SymlinkWay = "SymlinkToRepo"
	SymlinkFromRepo SymlinkWay = "SymlinkFromRepo"
)

func createSymlinks(folders []config.FoldersToSync, files []string, dottrFolder string, symlinkWay SymlinkWay) {
	homeDir := dirutils.GetHomeDirectory()

	// Check if dottr folder exists
	_, err := os.Stat(filepath.Join(homeDir, dottrFolder))

	if err != nil {
		log.Println("Dottr folder doest not exists")

		if !dirutils.CheckIfDottrFolderExists(dottrFolder) {
			dirutils.PromptToCreateDottrFolder(dottrFolder)
		}
	}

	// Link folders
	for _, folder := range folders {
		var sym symlink.Symlink

		if symlinkWay == SymlinkToRepo {
			sym = symlink.New(
				filepath.Join(homeDir, folder.Path),              // Source
				filepath.Join(homeDir, dottrFolder, folder.Path), // Target
			)
		} else {
			sym = symlink.New(
				filepath.Join(homeDir, dottrFolder, folder.Path), // Target
				filepath.Join(homeDir, folder.Path),              // Source
			)
		}

		if err := sym.Link(); err != nil {
			fmt.Printf("\033[31m not linked:\033[0m %s (%s)\n", err, folder.Path)
		} else {
			fmt.Printf("Dir: %s\033[32m linked! \033[0m \n", folder.Path)
		}

	}

	// Link files
	for _, file := range files {
		var sym symlink.Symlink

		if symlinkWay == SymlinkToRepo {
			sym = symlink.New(
				filepath.Join(homeDir, file),              // Source
				filepath.Join(homeDir, dottrFolder, file), // Target
			)
		} else {
			sym = symlink.New(
				filepath.Join(homeDir, dottrFolder, file), // Target
				filepath.Join(homeDir, file),              // Source
			)
		}

		if err := sym.Link(); err != nil {
			fmt.Printf("\033[31m not linked:\033[0m %s (%s)\n", err, file)
		} else {
			fmt.Printf("File: %s\033[32m linked! \033[0m \n", file)
		}

	}
}
