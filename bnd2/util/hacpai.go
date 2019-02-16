package util

import (
	"path/filepath"
)

type user struct {
	SaveDir   string `json:"saveDir"`
}

var User = &user{SaveDir: filepath.Join(HomeDir, "Downloads")}
