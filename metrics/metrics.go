package metrics

import (
	"fmt"
	"sort"

	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type Tracker struct {
	TotalTicks int

	CumulativeUtilization float64
	CumulativeFragmentation float64

	Makespan int

	JobQueueWaits map[string]int
	JobTurnarounds map[string]int
	JobServiceTimes map[string]int
}

func NewTracker() *Tracker {
	return &Tracker{
		JobQueueWaits: make(map[string]int),
		JobTurnarounds: make(map[string]int),
		JobServiceTimes: make(map[string]int),
	}
}

func (m *Tracker) Collect(
	c *cluster.Cluster,
	pendingJobs []*jobs.Job,
	completedJobs []*jobs.Job,
	currentTick int,
) {
	m.TotalTicks++

	totalCPU := 0
	usedCPU := 0

	for _, node := range c.Nodes {
		totalCPU += node.TotalCPU
		usedCPU += node.UsedCPU
	}

	if totalCPU > 0 {
		utilization := float64(usedCPU) / float64(totalCPU)
		m.CumulativeUtilization += utilization
	}
	m.CumulativeFragmentation += fragmentation(c, pendingJobs)

	for _, job := range completedJobs {
		if _, exists := m.JobTurnarounds[job.ID]; exists {
			continue
		}
		m.JobQueueWaits[job.ID] = job.StartTick - job.ArrivalTick
		m.JobTurnarounds[job.ID] = job.EndTick - job.ArrivalTick
		m.JobServiceTimes[job.ID] = job.EndTick - job.StartTick
	}	
	m.Makespan = currentTick
}

// Utilization alone hides stranded capacity. Fragmentation explains why jobs wait even when the cluster looks partially free.
func fragmentation (c *cluster.Cluster, pending []*jobs.Job) float64 {
	if len(pending) == 0 {
		return 0
	}

	total := 0.0

	for _, node := range c.Nodes {
		freeCPU := node.FreeCPU()
		freeMem := node.FreeMemory()

		if freeCPU == 0 && freeMem == 0 {
			continue
		}

		canPlaceAny := false

		for _, job := range pending {
			if freeCPU >= job.CPU && freeMem >= job.Memory {
				canPlaceAny = true
				break
			}
		}
		if !canPlaceAny {
			total += float64(freeCPU + freeMem)
		}
	}
	return total
}

func (m *Tracker) AverageUtilization() float64 {
	if m.TotalTicks == 0 {
		return 0
	}
	return (m.CumulativeUtilization / float64(m.TotalTicks)) * 100
}

func (m *Tracker) AverageFragmentation() float64 {
	if m.TotalTicks == 0 {
		return 0
	}
	return m.CumulativeFragmentation / float64(m.TotalTicks)
}

func (m *Tracker) AverageQueueWait() float64 {
	return averageFromMap(m.JobQueueWaits)
}

func (m *Tracker) AverageTurnaroundTime() float64 {
	return averageFromMap(m.JobTurnarounds)
}

func (m*Tracker) AverageServiceTime() float64 {
	return averageFromMap(m.JobServiceTimes)
}

func averageFromMap(values map[string]int) float64 {
	if len(values) == 0 {
		return 0
	}

	total := 0

	for _, value := range values {
		total += value
	}
	return float64(total) / float64(len(values))
}

func (m*Tracker) PercentileQueueWait(p float64) float64 {
	return percentileFromMap(m.JobQueueWaits, p)
}

func percentileFromMap(values map[string]int, p float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sorted := make([]int, 0, len(values))
	for _, v := range values {
		sorted = append(sorted, v)
	}
	sort.Ints(sorted)
	idx := int(p * float64(len(sorted)-1))
	return float64(sorted[idx])
}

func (m *Tracker) PrintReport(schedulerName string) {
	fmt.Println("==============================")
	fmt.Printf("%s Results\n", schedulerName)
	fmt.Println("==============================")

	fmt.Printf("Makespan: %d ticks\n", m.Makespan)
	fmt.Printf("Average Utilization: %.2f%%\n", m.AverageUtilization())
	// looks at stranded capacity due to there being a free node that is not at use because no pending job can fit into its capacity limit
	fmt.Printf("Average Fragmentation: %.2f\n", m.AverageFragmentation())
	fmt.Printf("Average Queue Wait: %.2f ticks\n", m.AverageQueueWait())
	fmt.Printf("Average Turnaround Time: %.2f ticks\n", m.AverageTurnaroundTime())
	fmt.Printf("Average Service Time: %.2f ticks\n", m.AverageServiceTime())
	fmt.Printf("P95 Queue Wait: %.2f ticks\n", m.PercentileQueueWait(0.95))

	fmt.Println()
}
