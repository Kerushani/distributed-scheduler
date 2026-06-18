type Engine struct {
	Cluster *cluster.Cluster
	Jobs []*jobs.Jobs

	Tick int

	Scheduler scheduler.Scheduler

	Metrics *metrics.Tracker

	PendingJobs []*jobs.Job
	RunningJobs map[string]*jobs.Job
}

func (e *Engine) Run(maxTicks int) {
	for e.Tick < maxTicks{
		e.injectJobs()

		e.Scheduler.Schedule(e.Cluster, e.PendingJobs, e.Tick)

		e.executeTick()

		e.Metrics.Collect(e.Cluster, e.RunningJobs, e.Tick)

		e.Tick++

		if e.allJobsCompleted(){
			break
		}
	}
}