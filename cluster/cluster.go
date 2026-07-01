package cluster

// Cluster is system state

import (
	"fmt"
	"sort"

	"distributed-scheduler/jobs"
)

type Cluster struct {
	Nodes map[string]*Node
}

type NodeSpec struct {
	CPU int
	Memory int
}

// makes a homogenous cluster -> cpu and memory are set
func NewCluster(nodeCount int, cpu int, memory int) *Cluster {
	specs := make([]NodeSpec, nodeCount)
	for i := range specs {
		specs[i] = NodeSpec{CPU: cpu, Memory: memory}
	}
	return NewHeterogeneousCluster(specs)
}

// makes a heterogenous cluster -> creates nodes with different CPU/Memory per node
func NewHeterogeneousCluster(specs []NodeSpec) *Cluster {
	nodes := make(map[string]*Node)

	for i, spec := range specs {
		id := fmt.Sprintf("node-%d", i)

		nodes[id] = &Node{
			ID: id,
			TotalCPU: spec.CPU,
			TotalMemory: spec.Memory,
			UsedCPU: 0,
			UsedMemory: 0,
			RunningJobs: make(map[string]*jobs.Job),
		}
	}
	return &Cluster{Nodes: nodes}
}

func (c *Cluster) AllNodes() []*Node {
	list := make([]*Node, 0, len(c.Nodes))
	for _, n := range c.Nodes {
		list = append(list, n)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}