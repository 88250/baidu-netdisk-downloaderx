package util

import (
	"encoding/base64"
	"encoding/json"
	"runtime"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/parnurzeal/gorequest"
)

const RHY = "https://rhythm.b3log.org"

type StatData struct {
	Size  string
	Speed string
}

type BND2 struct {
	MachineID string `json:"machineId"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	Version   string `json:"version"`
	Size      string `json:"size"`
	Speed     string `json:"speed"`
	Name      string `json:"name"`
}

func Stat(statData *StatData) {
	machineId, err := machineid.ProtectedID("bnd2")
	if nil != err {
		machineId = "error"
	}
	bnd2 := &BND2{
		MachineID: machineId,
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Version:   Ver,
		Size:      statData.Size,
		Speed:     statData.Speed,
		Name:      "bnd2",
	}

	data, err := json.Marshal(bnd2)
	if nil != err {
		return
	}

	requestData := map[string]interface{}{
		"data": base64.StdEncoding.EncodeToString(AESEncrypt(data)),
	}

	response, data, errs := gorequest.New().Post(RHY + "/bnd").SendStruct(requestData).EndBytes()
	if nil != errs {
		return
	}
	if 200 != response.StatusCode {
		return
	}
	resData := map[string]interface{}{}
	if err := json.Unmarshal(data, &resData); nil != err {
		return
	}
	sc := resData["sc"].(float64)
	switch sc {
	case 0:
		return
	case 1:
		if "windows" != runtime.GOOS {
			return
		}
	}
}

func CheckUpgrade() (needUpgrade bool) {
	response, data, errs := gorequest.New().Get(RHY+"/version/bnd2").
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

	latestVer := result["kernelVer"].(string)
	AppId = result["appId"].(string)

	return Ver < latestVer
}
