package scheduler

import (
	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type GreedyScheduler struct{}

type Scheduler interface {
	Schedule(c *cluster.Cluster, jobs []*jobs.Job, tick int) []Decision
}

type Decision struct {
	JobID  string
	NodeID string
}

func scoreNode(n *cluster.Node, j *jobs.Job) int {
	remainingCPU := n.TotalCPU - n.UsedCPU - j.CPU
	remainingMem := n.TotalMemory - n.UsedMemory - j.Memory

	return -(remainingCPU + remainingMem)
}

func (s *GreedyScheduler) Schedule(c *cluster.Cluster, jobs []*jobs.Job, tick int) []Decision {
	decisions := []Decision{}

	for _, job := range jobs {
		bestNodeID := ""
		bestScore := -1 << 31

		for _, node := range c.Nodes {
			if !node.CanFit(job) {
				continue
			}

			score := scoreNode(node, job)

			if score > bestScore {
				bestScore = score
				bestNodeID = node.ID
			}
		}

		if bestNodeID != "" {
			decisions = append(decisions, Decision{
				JobID:  job.ID,
				NodeID: bestNodeID,
			})
		}
	}
	return decisions
}
