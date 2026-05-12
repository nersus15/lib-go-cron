package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/webcore-go/webcore/infra/logger"
	"github.com/webcore-go/webcore/port"
)

// CronLibrary mengimplementasikan port.Library
type CronLibrary struct {
	scheduler *cron.Cron
}

type CronLoader struct {
	name string
}

func (a *CronLoader) SetName(name string) {
	a.name = name
}

func (a *CronLoader) Name() string {
	return a.name
}

func (l *CronLoader) Init(args ...any) (port.Library, error) {
	c := cron.New()

	cl := &CronLibrary{
		scheduler: c,
	}

	err := cl.Install(args...)
	if err != nil {
		return nil, err
	}

	cl.Connect()
	return cl, nil
}

// ----------------------- CronLibrary Methods -------------------

func (c *CronLibrary) Install(args ...any) error {
	logger.Info("Installing Cron Library")
	return nil
}

func (c *CronLibrary) Connect() error {
	c.scheduler.Start()
	logger.Info("Cron Scheduler connected and started")
	return nil
}

func (c *CronLibrary) Uninstall() error {
	c.scheduler.Stop()
	logger.Info("Cron Library uninstalled and scheduler stopped")
	return nil
}

// AddFunc memudahkan registrasi job dari luar library
func (c *CronLibrary) AddFunc(spec string, cmd func()) (int, error) {
	id, err := c.scheduler.AddFunc(spec, cmd)
	if err != nil {
		return 0, fmt.Errorf("failed to schedule cron job: %w", err)
	}
	logger.Debug("Cron job scheduled", "spec", spec, "entry_id", id)
	return int(id), nil
}

func (c *CronLibrary) Stop() {
	c.scheduler.Stop()
	logger.Info("Cron Scheduler stopped")
}
