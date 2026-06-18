func (n *Node) ExecuteTick(){
	for _, job := range n.RunningJobs {
		job.Duration -= 1

		if job.Duration <= 0 {
			n.UsedCPU -= job.CPU
			n.UsedMemory -= job.Memory

			delete(n.RunningJobs, job.ID)
		}
	}
}