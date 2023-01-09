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
    createSymlinks(loadedCfg.FoldersToSync, loadedCfg.FilesToSync, loadedCfg.DottrFolder)
    
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

func createSymlinks(folders []config.FoldersToSync, files []string, dottrFolder string) {
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
        sym := symlink.New(
            filepath.Join(homeDir, folder.Path), // Source
            filepath.Join(homeDir, dottrFolder, folder.Path), // Target
        )

        if err := sym.Link(); err != nil {
            fmt.Printf("\033[31m not linked:\033[0m %s (%s)\n", err, folder.Path)
        } else {
            fmt.Printf("Dir: %s\033[32m linked! \033[0m \n", folder.Path)
        }

    }

    // Link files
    for _, file := range files {
        sym := symlink.New(
            filepath.Join(homeDir, file), // Source
            filepath.Join(homeDir, dottrFolder, file), // Target
        )

        if err := sym.Link(); err != nil {
            fmt.Printf("\033[31m not linked:\033[0m %s (%s)\n", err, file)
        } else {
            fmt.Printf("File: %s\033[32m linked! \033[0m \n", file)
        }

    }
}
