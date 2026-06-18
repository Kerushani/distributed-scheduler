type Scheduler interface {
	Schedule(c *cluster.Cluster), jobs []*jobs.Job, tick int
}

func scoreNode(n *cluster.Node, j *jobs.Job) int {
	remainingCPU += n.TotalCPU - n.UsedCPU - j.CPU
	remainingMem := n.TotalMemory - n.UsedMemory - j.Memory

	return - remainingCPU - remainingMem
}

func (s *GreedyScheduler) Schedule(c *cluster.Cluster, jobs []*jobs.Job, tick int) {
	for _, job := range jobs {
		bestNode := ""
		bestScore := -1 << 31

		for _, node := range c.Nodes {
			if !canFit(node, job) {
				continue
			}

			score := scoreNode(node, job)

			if score > bestScore {
				bestScore = score
				bestNode = node.ID
			}
		}

		if bestNode != "" {
			assignJob(c.Nodes[bestNode], job, tick)
		}
	}
}