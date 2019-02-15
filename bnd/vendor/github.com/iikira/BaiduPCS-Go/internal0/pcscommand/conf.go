package pcscommand

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/iikira/BaiduPCS-Go/pcsutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	Ver            = "4.0.0"
	WindowTitle    = "BND v" + Ver
	WindowTitleGBK = ""
	PlsSelectFile  = "2. 点我选择要下载的文件（在网盘中将需要下载的文件移到根目录后再点【1. 刷新文件列表】）"
	BND2           = ""

	SK = []byte("696D887C9AA0611B")
)

var HomeDir = pcsutil.ExecutablePathJoin("")
var BndDir = filepath.Join(HomeDir, ".bnd")
var SaveDir = filepath.Join(HomeDir, "Downloads")
var BDUSS = ""
var AppID = 260149

func init() {
	titleData, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(WindowTitle)), simplifiedchinese.GBK.NewEncoder()))
	WindowTitleGBK = string(titleData)
}
