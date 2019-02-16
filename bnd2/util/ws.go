package util

import (
	"gopkg.in/olahol/melody.v1"
)

var s *melody.Session

func SetPushChan(session *melody.Session) {
	s = session
}

func Push(msg []byte) {
	s.Write(msg)
}
