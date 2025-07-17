package cron_scheduler

import (
	"reflect"
	"runtime"

	"github.com/robfig/cron/v3"
)

type RegistryItem struct {
	JobID *cron.EntryID

	Enabled       bool
	RunOnRegister bool

	Spec     string
	TaskFunc func()
}

// Gets the name of the cron job based on the function name.
func (i *RegistryItem) Name() string {
	return runtime.FuncForPC(reflect.ValueOf(i.TaskFunc).Pointer()).Name()
}
