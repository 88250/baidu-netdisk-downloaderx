package util

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/b3log/bnd2/log"
	"github.com/dustin/go-humanize"
	"github.com/parnurzeal/gorequest"
)

var logger = log.NewLogger(os.Stdout)

var AppId = "260149"
var BDUSS = ""

func DowanloadURL(path string) string {
	ret := &url.URL{
		Scheme: "http",
		Host:   "pcs.baidu.com",
		Path:   "/rest/2.0/pcs/file",
	}

	q := ret.Query()
	q.Set("app_id", AppId)
	q.Set("method", "download")
	q.Set("path", path)
	ret.RawQuery = q.Encode()

	return ret.String()
}

type File struct {
	FsId   uint64 `json:"fs_id"`
	Path   string `json:"path"`
	Name   string `json:"server_filename"`
	IsDir  int8   `json:"isdir"`
	Size   uint64 `json:"size"`
	HSize  string `json:"hSize"`
	Mtime  int64  `json:"mtime"`
	HMtime string `json:"hMtime"`
}

func Ls(path, by, order string) []*File {
	lsURL := &url.URL{
		Scheme: "http",
		Host:   "pcs.baidu.com",
		Path:   "/rest/2.0/pcs/file",
	}

	q := lsURL.Query()
	q.Set("app_id", AppId)
	q.Set("method", "list")
	q.Set("path", path)
	q.Set("by", by)
	q.Set("order", order)
	q.Set("limit", "0-2147483647")
	lsURL.RawQuery = q.Encode()

	var ret []*File
	result := map[string]interface{}{}
	response, body, errs := gorequest.New().Get(lsURL.String()).AppendHeader("Cookie", "BDUSS="+BDUSS).Timeout(5*time.Second).Retry(3, time.Second).EndStruct(&result)
	if nil != errs {
		logger.Errorf("Request Baidu PCS failed [err=%s]", errs)

		return ret
	}
	if http.StatusOK != response.StatusCode {
		logger.Errorf("Request Baidu PCS failed [code=%d, body=%s]", response.StatusCode, body)

		return ret
	}

	list := result["list"].([]interface{})
	bytes, err := json.Marshal(list)
	if nil != err {
		logger.Errorf("Parse Baidu PCS failed [err=%s]", err)

		return ret
	}

	if err = json.Unmarshal(bytes, &ret); nil != err {
		logger.Errorf("Parse Baidu PCS failed [err=%s]", err)

		return ret
	}

	for i := 0; i < len(ret); i++ {
		f := ret[i]
		f.HSize = humanize.Bytes(f.Size)
		f.HMtime = time.Unix(f.Mtime, 0).Format("2006-01-02 15:04:05")
	}

	return ret
}
