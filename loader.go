package cron

import (
	"github.com/robfig/cron/v3"
	"github.com/webcore-go/webcore/port"
)

type CronLoader struct {
	name string
}

func (a *CronLoader) SetName(name string) { a.name = name }
func (a *CronLoader) Name() string        { return a.name }

func (l *CronLoader) Init(args ...any) (port.Library, error) {
	// Inisialisasi scheduler di sini
	cl := &CronLibrary{
		scheduler: cron.New(),
	}

	// Framework Webcore akan mengelola pemanggilan Install() dan Connect()
	return cl, nil
}
