package rpc

import (
	"log"
)

type Event struct {
	Gid string `json:"gid"` // GID of the download
}

type Notifier interface {
	OnStart([]Event)
	OnPause([]Event)
	OnStop([]Event)
	OnComplete([]Event)
	OnError([]Event)
	OnBtComplete([]Event)
}

type DummyNotifier struct {
}

func (d *DummyNotifier) OnStart(es []Event) {
	log.Printf("%s started.\n", es)
}
func (d *DummyNotifier) OnPause(es []Event) {
	log.Printf("%s paused.\n", es)
}
func (d *DummyNotifier) OnStop(es []Event) {
	log.Printf("%s stopped.\n", es)
}
func (d *DummyNotifier) OnComplete(es []Event) {
	log.Printf("%s completed.\n", es)
}
func (d *DummyNotifier) OnError(es []Event) {
	log.Printf("%s error.\n", es)
}
func (d *DummyNotifier) OnBtComplete(es []Event) {
	log.Printf("bt %s completed.\n", es)
}
