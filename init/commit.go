package init

import (
	"github.com/TheMartes/dottr/config"
	"github.com/TheMartes/dottr/dirutils"
)

func Commit() {
    cfg := config.LoadConfig()

    // Reset the dottr folder
    dirutils.RemoveDottrFolder(cfg.DottrFolder)
    dirutils.CreateDottrFolder(cfg.DottrFolder)

    // TODO: Copy all the files defined in config to the dottr and remove them from source
}
