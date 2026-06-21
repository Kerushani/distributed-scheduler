package cluster

// Cluster is system state

import (
	"fmt"
	"distributed-scheduler/jobs"
)

type Cluster struct {
	Nodes map[string]*Node
}

func NewCluster(nodeCount int, cpu int, memory int) *Cluster {
	nodes := make(map[string]*Node)

	for i := 0; i < nodeCount; i++ {
		id := fmt.Sprintf("node-%d", i)

		nodes[id] = &Node{
			ID: id,
			TotalCPU: cpu,
			TotalMemory: memory,
			UsedCPU: 0,
			UsedMemory: 0,
			RunningJobs: make(map[string]*jobs.Job),
		}
	}
	return &Cluster{
		Nodes: nodes,
	}
}

func (c *Cluster) AllNodes() []*Node {
	list := []*Node{}
	for _, n := range c.Nodes {
		list = append(list, n)
	}
	return list
}