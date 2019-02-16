package util

import (
	"bytes"
	"os"
	"os/exec"
	u "os/user"
	"runtime"
	"strings"
	"time"

	"github.com/mitchellh/go-ps"
)

var ppid = os.Getppid()

func ParentExited() {
	for range time.Tick(2 * time.Second) {
		process, e := ps.FindProcess(ppid)
		if nil == process || nil != e {
			logger.Info("can't find parent process, exited")

			os.Exit(0)
		}
	}
}

func IsWindows() bool {
	return "windows" == runtime.GOOS
}

func IsLinux() bool {
	return "linux" == runtime.GOOS
}

func IsMac() bool {
	return "darwin" == runtime.GOOS
}

func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func UserHome() string {
	user, err := u.Current()
	if nil == err {
		return user.HomeDir
	}

	if IsWindows() {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() string {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		logger.Errorf("get user home path failed [%s]", err)

		return ""
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		logger.Errorf("blank output when reading home directory")

		return ""
	}

	return result
}

func homeWindows() string {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		logger.Errorf("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")

		return ""
	}

	return home
}
