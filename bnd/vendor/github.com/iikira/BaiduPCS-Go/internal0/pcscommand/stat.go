package pcscommand

import (
	"encoding/base64"
	"encoding/json"
	"runtime"
	"strconv"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/parnurzeal/gorequest"
)

const RHY = "https://rhythm.b3log.org"

type StatData struct {
	Size  string
	Speed string
}

type BND struct {
	MachineID string `json:"machineId"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
	Size      string `json:"size"`
	Speed     string `json:"speed"`
	Name      string `json:"name"`
}

func Stat(statData *StatData) {
	machineId, err := machineid.ProtectedID("bnd")
	if nil != err {
		machineId = "error"
	}
	bnd := &BND{
		MachineID: machineId,
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Version:   Ver,
		Size:      statData.Size,
		Speed:     statData.Speed,
		Name:      "bnd",
	}

	data, err := json.Marshal(bnd)
	if nil != err {
		return
	}

	requestData := map[string]interface{}{
		"data": base64.StdEncoding.EncodeToString(AESEncrypt(data)),
	}

	gorequest.New().Post(RHY + "/bnd").SendStruct(requestData).EndBytes()
}

func CheckUpgrade() (needUpgrade bool, tip, newVerURL string) {
	needUpgrade, tip, newVerURL = true, "检测到环境异常请重启，如果重启多次仍然无效请重新下载最新版后再试", "https://share.weiyun.com/57zViCm"
	response, data, errs := gorequest.New().Get(RHY+"/version/bnd/latest").
		Timeout(3*time.Second).Retry(3, time.Second).EndBytes()
	if nil != errs {
		return
	}
	if 200 != response.StatusCode {
		return
	}

	result := map[string]interface{}{}
	if err := json.Unmarshal(data, &result); nil != err {
		return
	}

	dataStr := result["data"].(string)
	data, err := base64.StdEncoding.DecodeString(dataStr)
	if nil != err {
		return
	}

	data = AESDecrypt(data)

	result = map[string]interface{}{}
	if err := json.Unmarshal(data, &result); nil != err {
		return
	}

	latestVer := result["bndVersion"].(string)
	tip = result["bndTip"].(string)
	BND2 = result["bnd2Download"].(string)
	AppID, _ = strconv.Atoi(result["appId"].(string))
	if Ver < latestVer {
		return true, tip, result["bndDownloadHP"].(string)
	}

	return false, tip, result["bndDownloadHP"].(string)
}
