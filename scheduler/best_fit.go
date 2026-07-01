package scheduler

import (
	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type BestFitScheduler struct{}

func (s *BestFitScheduler) Schedule(c *cluster.Cluster, pending []*jobs.Job, tick int) []Decision {
	return placeJobs(c, pending, func(nodes []*cluster.Node, job *jobs.Job) string {
		return pickByScore(nodes, job, scoreBestFit)
	})
}