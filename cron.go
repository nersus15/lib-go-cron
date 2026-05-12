package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/webcore-go/webcore/infra/logger"
	"github.com/webcore-go/webcore/port"
)

// ==========================================
// LIBRARY IMPLEMENTATION
// ==========================================

// CronLibrary mengimplementasikan port.Connector
type CronLibrary struct {
	scheduler *cron.Cron
}

// Pastikan CronLibrary memenuhi kontrak port.Connector
var _ port.Connector = (*CronLibrary)(nil)

func (c *CronLibrary) Install(args ...any) error {
	logger.Info("CronLibrary: Installed")
	return nil
}

func (c *CronLibrary) Uninstall() error {
	logger.Info("CronLibrary: Uninstalled")
	return nil
}

// Connect dipanggil otomatis oleh framework untuk menyalakan scheduler
func (c *CronLibrary) Connect() error {
	c.scheduler.Start()
	logger.Info("CronLibrary: Scheduler Engine Started")
	return nil
}

// Disconnect dipanggil otomatis oleh framework saat aplikasi mati (graceful shutdown)
func (c *CronLibrary) Disconnect() error {
	c.scheduler.Stop()
	logger.Info("CronLibrary: Scheduler Engine Stopped")
	return nil
}

// AddFunc adalah helper untuk meregistrasi task dari application layer
func (c *CronLibrary) AddFunc(spec string, cmd func()) (int, error) {
	id, err := c.scheduler.AddFunc(spec, cmd)
	if err != nil {
		return 0, fmt.Errorf("cron failed to schedule: %w", err)
	}
	return int(id), nil
}

// ==========================================
// LOADER IMPLEMENTATION
// ==========================================

type CronLoader struct {
	name string
}

func (a *CronLoader) SetName(name string) { a.name = name }
func (a *CronLoader) Name() string         { return a.name }

func (l *CronLoader) Init(args ...any) (port.Library, error) {
	// Inisialisasi scheduler di sini
	cl := &CronLibrary{
		scheduler: cron.New(),
	}

	// Framework Webcore akan mengelola pemanggilan Install() dan Connect()
	return cl, nil
}
