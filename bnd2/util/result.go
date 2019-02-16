package util

import (
	"encoding/json"

	)

type Result struct {
	Code int         `json:"code"`
	Cmd  string      `json:"cmd"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

func NewCmdResult(cmd string) *Result {
	ret := NewResult()
	ret.Cmd = cmd

	return ret
}

func (r *Result) Bytes() []byte {
	ret, err := json.Marshal(r)
	if nil != err {
		logger.Errorf("marshal result [%#v] failed [%s]", r, err)
	}

	return ret
}
