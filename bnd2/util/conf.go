package util

import (
	"path/filepath"
)

const Ver = "2.0.0"

var SK = []byte("696D887C9AA0611B")
var UserAgent = "BND2/v" + Ver

const (
	ServerPort = 6804
	AriaPort   = 6805
)

var HomeDir = UserHome()
var BndDir = filepath.Join(HomeDir, ".bnd2")
