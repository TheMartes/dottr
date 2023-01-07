package dirutils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func CheckIfDottrFolderExists(dottrFolder string) bool {
    // Check if dottr folder exists
    _, err := os.Stat(filepath.Join(GetHomeDirectory(), dottrFolder))

    if err != nil {
        return false
    }

    return true
}

func PromptToCreateDottrFolder(dottrFolder string) {
    prompt := promptui.Select{
        Label: "Do you want to create dottrConfig?",
        Items: []string{"Yes", "No"},
    }
    _, result, err := prompt.Run()
    if err != nil {
        log.Fatalf("Prompt failed %v\n", err)
    }

    if result == "Yes" {
        if err := os.Mkdir(filepath.Join(GetHomeDirectory(), dottrFolder), os.ModePerm); err != nil {
            log.Println(err)
        }
    } else {
        log.Fatal("Dottr folder failed to be created, please create the folder and then proceed dottr init once again.")
    }
}

func CreateDottrFolder(dottrFolder string) {
    err := os.Mkdir(filepath.Join(GetHomeDirectory(), dottrFolder), os.ModePerm)

    if err != nil {
        log.Fatal(err)
    }
}

func RemoveDottrFolder(dottrFolder string) {
    err := os.Remove(filepath.Join(GetHomeDirectory(), dottrFolder))

    if err != nil {
        log.Fatal(err)
    }
}
