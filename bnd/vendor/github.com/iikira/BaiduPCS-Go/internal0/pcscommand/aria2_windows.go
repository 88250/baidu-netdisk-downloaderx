package pcscommand

import (
	"bytes"
	"math/rand"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/andlabs/ui"
	"github.com/iikira/BaiduPCS-Go/baidupcs"
	"github.com/iikira/BaiduPCS-Go/internal0/pcsconfig"
	"github.com/iikira/BaiduPCS-Go/pcsutil/converter"
	"github.com/zyxar/argo/rpc"
)

const ARIA2_PORT = 6805

func GetAria2DownloadFunc(uiProgress *UIProgress, path, savePath, s string) baidupcs.DownloadFunc {
	return func(downloadURL string, jar *cookiejar.Jar) (speed int64, err error) {
		cookies := jar.Cookies((&url.URL{Scheme: "http", Host: "pan.baidu.com"}))
		cookie := "Cookie:BDUSS=" + cookies[0].Value
		dir := filepath.Dir(savePath)
		name := filepath.Base(savePath)
		secret := "hacpai.com"
		gid := RandomHexStr(16)

		//fmt.Println("aria2c.exe", "-s"+s, "-k2M", "-x"+s, "--header="+cookie, "\""+downloadURL+"\"")
		homeDir := pcsconfig.UserHome()
		aria2 := filepath.Join(homeDir, ".bnd", "aria2c.exe")
		aria2Cmd := exec.Command(aria2, "--daemon=true", "-s"+s, "--header="+cookie, "-k2M", "-x"+s, "-d", dir, "-o", name,
			"--file-allocation=prealloc", "--enable-color=false", "--no-conf=true", "--stop-with-process="+strconv.Itoa(os.Getpid()),
			"--enable-rpc=true", "--rpc-listen-port="+strconv.Itoa(ARIA2_PORT), "--rpc-secret="+secret, "--gid="+gid,
			downloadURL)
		aria2Cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		var outb, errb bytes.Buffer
		aria2Cmd.Stdout = &outb
		aria2Cmd.Stderr = &errb
		r, err := rpc.New("http://localhost:"+strconv.Itoa(ARIA2_PORT)+"/jsonrpc", secret)
		if nil != err {
			return
		}
		var total int64
		start := time.Now()
		go func() {
			defer r.Shutdown()

			for {
				time.Sleep(time.Second)

				info, err := r.TellStatus(gid)
				if nil != err {
					continue
				}

				if "error" == info.Status || "complete" == info.Status {
					break
				}

				cl, _ := strconv.ParseInt(info.CompletedLength, 10, 64)
				cls := converter.ConvertFileSize(int64(cl), 2)
				tl, _ := strconv.ParseInt(info.TotalLength, 10, 64)
				if 1 > tl {
					continue
				}
				total = tl
				tls := converter.ConvertFileSize(int64(tl), 2)
				s, _ := strconv.ParseInt(info.DownloadSpeed, 10, 64)
				ss := converter.ConvertFileSize(int64(s), 2)
				eta := float64(tl-cl) / float64(s)
				leftStr := (time.Duration(eta) * time.Second).String()
				ui.QueueMain(func() {
					uiProgress.PB.SetValue(int(float64(cl) / float64(tl) * 100))
					uiProgress.PL.SetText(cls + "/" + tls + " 速度 " + ss + " 估计剩余 " + leftStr)
				})
			}
		}()

		err = aria2Cmd.Start()
		if nil != err {
			Downloading = false

			return
		}

		err = aria2Cmd.Wait()
		elapsed := time.Now().Sub(start).Seconds()
		speed = int64(float64(total) / float64(elapsed))
		Downloading = false
		if nil != err {
			ui.QueueMain(func() {
				uiProgress.PF.SetText(ShortLabel("下载异常 " + errb.String()))
			})
		} else {
			ui.QueueMain(func() {
				uiProgress.PF.SetText(ShortLabel("下载完成 " + path))
			})
		}

		return
	}
}

func RandomHexStr(n int) string {
	var letter = []rune("abcdefABCDEF0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
