package cron_scheduler

import "github.com/robfig/cron/v3"

type Registry struct {
	s *cron.Cron

	OnJobAddSuccess func(job *RegistryItem)
	OnJobAddFailure func(job *RegistryItem, err error)
}

func NewRegistry(opts ...cron.Option) *Registry {
	return &Registry{
		s: cron.New(opts...),
	}
}

func (r *Registry) AddJobs(jobs []RegistryItem) {
	for _, job := range jobs {
		if !job.Enabled {
			continue
		}

		jobID, err := r.s.AddFunc(job.Spec, job.TaskFunc)
		if err != nil {
			if r.OnJobAddFailure != nil {
				r.OnJobAddFailure(&job, err)
			}

			continue
		}
		job.JobID = &jobID

		if r.OnJobAddSuccess != nil {
			r.OnJobAddSuccess(&job)
		}

		if job.RunOnRegister {
			go job.TaskFunc()
		}
	}
}

func (r *Registry) Start() {
	r.s.Start()
}
