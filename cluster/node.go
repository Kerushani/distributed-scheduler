package cluster

import "distributed-scheduler/jobs"

type Node struct {
	ID          string
	TotalCPU    int
	TotalMemory int
	UsedCPU     int
	UsedMemory  int
	RunningJobs map[string]*jobs.Job
}

func (n *Node) FreeCPU() int {
	return n.TotalCPU - n.UsedCPU
}

func (n *Node) FreeMemory() int {
	return n.TotalMemory - n.UsedMemory
}

func (n *Node) CanFit(job *jobs.Job) bool {
	return n.UsedCPU+job.CPU <= n.TotalCPU && n.UsedMemory+job.Memory <= n.TotalMemory
}

func (n *Node) Assign(job *jobs.Job, tick int) {
	if job.StartTick == 0 {
		job.StartTick = tick
	}

	n.UsedCPU += job.CPU
	n.UsedMemory += job.Memory

	n.RunningJobs[job.ID] = job
	job.AssignedNode = n.ID
}

func (n *Node) ExecuteTick(tick int) []*jobs.Job {
	var completed []*jobs.Job

	for id, job := range n.RunningJobs {
		job.Duration--

		if job.Duration <= 0 {
			job.EndTick = tick + 1
			n.UsedCPU -= job.CPU
			n.UsedMemory -= job.Memory
			completed = append(completed, job)
			delete(n.RunningJobs, id)
		}
	}

	return completed
}
