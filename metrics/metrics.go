package metrics

import (
	"fmt"

	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type Tracker struct {
	TotalTicks int

	CumulativeUtilization float64

	Makespan int

	JobCompletionTimes map[string]int
}

func NewTracker() *Tracker {
	return &Tracker{
		JobCompletionTimes: make(map[string]int),
	}
}

func (m *Tracker) Collect(
	c *cluster.Cluster,
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

	for _, job := range completedJobs {
		if _, exists := m.JobCompletionTimes[job.ID]; !exists {
			m.JobCompletionTimes[job.ID] = job.EndTick - job.StartTick
		}
	}
	m.Makespan = currentTick
}

func (m *Tracker) AverageUtilization() float64 {
	if m.TotalTicks == 0 {
		return 0
	}
	return (m.CumulativeUtilization / float64(m.TotalTicks)) * 100
}

func (m *Tracker) AverageTurnaroundTime() float64 {
	if len(m.JobCompletionTimes) == 0 {
		return 0
	}

	total := 0
	for _, turnaround := range m.JobCompletionTimes {
		total += turnaround
	}

	return float64(total) / float64(len(m.JobCompletionTimes))
}

func (m *Tracker) PrintReport(schedulerName string) {
	fmt.Println("==============================")
	fmt.Printf("%s Results\n", schedulerName)
	fmt.Println("==============================")

	fmt.Printf("Makespan: %d ticks\n", m.Makespan)
	fmt.Printf("Average Utilization: %.2f%%\n", m.AverageUtilization())
	fmt.Printf("Average Turnaround Time: %.2f ticks\n", m.AverageTurnaroundTime())

	fmt.Println()
}
