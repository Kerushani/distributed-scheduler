"""
A job is a unit of compute work / workload -> a job is doing stateful execution.
"""

package jobs

type Job struct {
	ID string

	CPU int
	Memory int
	Duration int //remaining ticks

	Priority int

	StartTick int
	EndTick int

	AssignedNode string
}