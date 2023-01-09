package preview

import (
	"log"
	"path"

	"github.com/TheMartes/dottr/config"
	"github.com/TheMartes/dottr/dirutils"
	"github.com/TheMartes/dottr/fileutils"
)

func Apply() {
	cfg := config.LoadConfig()

	// Reset the dottr folder
	dirutils.RemoveDottrFolder(cfg.DottrFolder)
	dirutils.CreateDottrFolder(cfg.DottrFolder)

	dottrFolder := cfg.DottrFolder

	// TODO: Copy all the files defined in config to the dottr and remove them from source
	for _, file := range cfg.FilesToSync {
		src := path.Join(dirutils.GetHomeDirectory(), file)
		dest := path.Join(dirutils.GetHomeDirectory(), dottrFolder, file)
		err := fileutils.CopyFile(src, dest)

		if err != nil {
			log.Fatal(err)
		}

		fileutils.RemoveFile(src)
	}

	for _, folder := range cfg.FoldersToSync {
		src := path.Join(dirutils.GetHomeDirectory(), folder.Path)
		dest := path.Join(dirutils.GetHomeDirectory(), dottrFolder, folder.Path)

		err := dirutils.CopyDirectory(src, dest)
		if err != nil {
			log.Fatal(err)
		}

		dirutils.RemoveFolder(src)
	}

	createSymlinks(cfg.FoldersToSync, cfg.FilesToSync, cfg.DottrFolder, SymlinkFromRepo)
}
