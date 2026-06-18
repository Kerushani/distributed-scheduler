pakage sim

import (
	"scheduler-sim/cluster"
	"scheduler-sim/jobs"
	"scheduler-sim/scheduler"
	"scheduler-sim/metrics"
)

type Engine struct {
	Cluster *cluster.Cluster
	Jobs []*jobs.Jobs

	Tick int

	Scheduler scheduler.Scheduler

	Metrics *metrics.Tracker

	PendingJobs []*jobs.Job
	RunningJobs map[string]*jobs.Job
	CompletedJobs []*jobs.Job
}

type Decision struct {
	JobID string
	NodeID string
}

func (e *Enginer) applyDecisions(decisions []Decision){
	for _, d := range decisions{
		job := e.findJob(d.JobID)
		node := e.Cluster.Nodes[d.NodeID]

		if node.CanFit(job){
			node.Assign(job, e.Tick)

			e.removePending(job.ID)
		}
	}
}

func (e *Engine) executeTick() {
	for _, node := range e.Cluster.Nodes{
		node.Tick()

		for _, job := range node.RunningJobs{
			if job.Duration <= 0 {
				e.CompletedJobs = append(e.CompletedJobs, job)
			}
		}
	}
}

func (e *Engine) loadJobs() {
	//no-op for v1 since no batch mode
}

func (e *Engine) Run(maxTicks int) {
	for e.Tick < maxTicks{
		e.loadJobs()
		decisions := e.Scheduler.Schedule(e.Cluster, e.PendingJobs, e.Tick)
		e.applyDecisions(decisions)
		e.ExecuteTick()
		e.Metrics.Collect(e.Cluster, e.PendingJobs, e.CompletedJobs, e.Tick)

		if len(e.PendingJobs) == 0 && e.allNodesIdle(){
			break
		}
		e.Tick++
	}
}