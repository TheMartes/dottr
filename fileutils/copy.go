package fileutils

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

type FilePermission os.FileMode

// FilePermissions are defined in a way where
// 1. User (U), Group (G), Others (O) are always first
// 2. Followed by allowed permissions Read (R), Write (W), Execute (E)
// e.g. Permissions for user to Read and Group to execute granting nothing to others will be URGE = 0410
const (
    URWGROR FilePermission = 0644
)

func CopyFile(src string, dest string) {
    bytesRead, err := ioutil.ReadFile(src)

    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile(dest, bytesRead, fs.FileMode(URWGROR))

    if err != nil {
        log.Fatal(err)
    }
}
