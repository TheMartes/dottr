package preview

import (
	"log"
	"os"
	"path/filepath"

	"github.com/TheMartes/dottr/config"
	"github.com/TheMartes/dottr/dirutils"
	"github.com/avakarev/go-symlink"
)

func Init() {
	homeDir := dirutils.GetHomeDirectory()

	// Check if config exists
	dottrConfigPath := filepath.Join(homeDir, config.DottrConfigName)
	_, err := os.Stat(dottrConfigPath)

	if err != nil {
		log.Fatalf(
			"Config was not found in \"%s\". In order for dottr to work, %s needs to be present.",
			homeDir,
			config.DottrConfigName,
		)
	}

	// Load config
	loadedCfg := config.LoadConfig()
	createSymlinks(loadedCfg.FoldersToSync, loadedCfg.FilesToSync, loadedCfg.DottrFolder, SymlinkToRepo)

	if loadedCfg.IncludeConfig {
		sym := symlink.New(
			filepath.Join(homeDir, config.DottrConfigName),
			filepath.Join(homeDir, loadedCfg.DottrFolder, config.DottrConfigName),
		)

		sym.Link()
	} else {
		filepath := filepath.Join(homeDir, loadedCfg.DottrFolder, config.DottrConfigName)
		// os.Stat returns an err, which might not be used to the best here, but essentially if err == nil
		// that means file does exists
		_, fileDoestNotExist := os.Stat(filepath)

		// Check if file exists
		if fileDoestNotExist == nil {
			os.Remove(filepath)
		}
	}

	return
}
