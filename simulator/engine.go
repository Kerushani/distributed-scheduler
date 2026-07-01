package simulator

import (
	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
	"distributed-scheduler/metrics"
	"distributed-scheduler/scheduler"
)

type Engine struct {
	Cluster *cluster.Cluster
	Jobs    []*jobs.Job

	Tick int

	Scheduler scheduler.Scheduler

	Metrics *metrics.Tracker

	PendingJobs   []*jobs.Job
	CompletedJobs []*jobs.Job

	nextJobIndex int
}

func (e *Engine) applyDecisions(decisions []scheduler.Decision) {
	for _, d := range decisions {
		job := e.findJob(d.JobID)
		if job == nil {
			continue
		}
		node := e.Cluster.Nodes[d.NodeID]
		if node == nil {
			continue
		}
		if node.CanFit(job) {
			node.Assign(job, e.Tick)
			e.removePending(job.ID)
		}
	}
}

func (e *Engine) findJob(id string) *jobs.Job {
	for _, job := range e.PendingJobs {
		if job.ID == id {
			return job
		}
	}
	return nil
}

func (e *Engine) removePending(id string) {
	for i, job := range e.PendingJobs {
		if job.ID == id {
			e.PendingJobs = append(e.PendingJobs[:i], e.PendingJobs[i+1:]...)
			return
		}
	}
}

func (e *Engine) executeTick() {
	for _, node := range e.Cluster.Nodes {
		completed := node.ExecuteTick(e.Tick)
		e.CompletedJobs = append(e.CompletedJobs, completed...)
	}
}

func (e *Engine) loadJobs() {
	for e.nextJobIndex < len(e.Jobs) && e.Jobs[e.nextJobIndex].ArrivalTick <= e.Tick{
		e.PendingJobs = append(e.PendingJobs, e.Jobs[e.nextJobIndex])
		e.nextJobIndex++
	}
}

func (e *Engine) allNodesIdle() bool {
	for _, node := range e.Cluster.Nodes {
		if len(node.RunningJobs) > 0 {
			return false
		}
	}
	return true
}

func (e *Engine) allJobsReleased() bool {
	return e.nextJobIndex >= len(e.Jobs)
}

func (e *Engine) Run(maxTicks int) {
	for e.Tick < maxTicks {
		e.loadJobs()
		decisions := e.Scheduler.Schedule(e.Cluster, e.PendingJobs, e.Tick)
		e.applyDecisions(decisions)
		e.executeTick()
		e.Metrics.Collect(e.Cluster, e.PendingJobs, e.CompletedJobs, e.Tick)

		if e.allJobsReleased() && len(e.PendingJobs) == 0 && e.allNodesIdle() {
			break
		}
		e.Tick++
	}
}
