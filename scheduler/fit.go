package scheduler

import (
	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type nodePicker func(nodes []*cluster.Node, job *jobs.Job) string

func placeJobs(c *cluster.Cluster, pending []*jobs.Job, pick nodePicker) []Decision {
	decisions := []Decision{}
	nodes := c.AllNodes()

	for _, job := range pending {
		fitting := fittingNodes(nodes, job)
		if len(fitting) == 0 {
			continue
		}

		nodeID := pick(fitting, job)

		if nodeID != "" {
			decisions = append(decisions, Decision {
				JobID: job.ID,
				NodeID: nodeID,
			})
		}
	}
	return decisions
}

func fittingNodes(nodes []*cluster.Node, job *jobs.Job) []*cluster.Node {
	fitting := []*cluster.Node{}
	for _, node := range nodes {
		if node.CanFit(job){
			fitting = append(fitting, node)
		}
	}
	return fitting
}

func scoreBestFit(n *cluster.Node, j *jobs.Job) int {
	remainingCPU := n.FreeCPU() - j.CPU
	remainingMem := n.FreeMemory() - j.Memory

	return -(remainingCPU + remainingMem)
}

func scoreWorstFit(n *cluster.Node, j *jobs.Job) int {
	remainingCPU := n.FreeCPU() - j.CPU
	remainingMem := n.FreeMemory() - j.Memory
	return remainingCPU + remainingMem
}

func pickByScore(nodes []*cluster.Node, job *jobs.Job, scoreFn func(*cluster.Node, *jobs.Job) int) string {
	bestNodeId := ""
	bestScore := -1 << 31

	for _, node := range nodes{
		score := scoreFn(node, job)
		if score > bestScore {
			bestScore = score
			bestNodeId = node.ID
		}
	}

	return bestNodeId
}