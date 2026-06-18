"""
A node is a resource container -> can think about this 
like a simplified Kubernetes node / resource container
"""
package cluster

import "scheduler-sim/jobs"

type Node struct {
	ID string
	TotalCPU int
	TotalMemory int
	UsedCPU int
	UsedMemory int
	RunningJobs map[string]*jobs.Job
}

func (n *Node) CanFit(job *jobs.Job) bool {
	return n.UsedCPU+job.CPU <= n.TotalCPU && n.UsedMemory+job.Memory <= n.TotalMemory
}

func (n *Node) Assign(job *jobs.Job, tick int){
	if job.StartTick == 0 {
		job.StartTick = tick
	}

	n.UsedCPU += job.CPU
	n.UsedMemory += job.Memory

	n.RunningJobs[job.ID] = job
	job.AssignedNode = n.ID
}

func (n *Node) ExecuteTick(){
	for id, job := range n.RunningJobs {
		job.Duration--

		if job.Duration <= 0 {
			n.UsedCPU -= job.CPU
			n.UsedMemory -= job.Memory

			job.EndTick = job.StartTick
			delete(n.RunningJobs, id)
		}
	}
}