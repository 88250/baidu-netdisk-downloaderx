package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/zyxar/argo/rpc"
)

var R rpc.Protocol

func StartAria2() {
	aria2 := filepath.Join(BndDir, "aria2c")
	if IsWindows() {
		aria2 += ".exe"
	}

	x := "16"
	if IsWindows() {
		x = "128"
	}

	cookie := "Cookie:BDUSS=" + BDUSS
	ua := "User-Agent=netdisk;7.8.1;Red;android-android;4.3"
	referer := "http://pan.baidu.com/disk/home"
	aria2Cmd := exec.Command(aria2, "--daemon=true", "--header="+cookie, "--header="+ua, "--header="+referer,
		"-x"+x, "-s1024", "-k1m",
		"--file-allocation=prealloc", "--enable-color=false", "--no-conf=true", "--stop-with-process="+strconv.Itoa(os.Getpid()),
		"--enable-rpc=true", "--rpc-listen-port="+strconv.Itoa(AriaPort), "--rpc-secret=b3log.org")
	hideWindow(aria2Cmd)

	go func() {
		if bytes, err := aria2Cmd.CombinedOutput(); nil != err {
			logger.Fatalf("start aria2 failed [%s]", bytes)
		}
	}()

	time.Sleep(time.Second)
	var err error
	R, err = rpc.New("http://localhost:"+strconv.Itoa(AriaPort)+"/jsonrpc", "b3log.org")
	if nil != err {
		logger.Fatalf("start aria2 rpc failed [%s]", err)
	}

	info, _ := R.GetVersion()
	logger.Debug("connected to aria2 v" + info.Version)
}
