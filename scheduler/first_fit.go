package scheduler

import (
	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type FirstFitScheduler struct{}

func (s *FirstFitScheduler) Schedule(c *cluster.Cluster, pending []*jobs.Job, tick int) []Decision {
	return placeJobs(c, pending, pickFirstFit)
}

func pickFirstFit(nodes []*cluster.Node, job *jobs.Job) string {
	for _, node := range nodes {
		if node.CanFit(job) {
			return node.ID
		}
	}
	return ""
}