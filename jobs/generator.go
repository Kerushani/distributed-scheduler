package jobs

import "fmt"

type Profile string

// a generator that produces workloads with arrival patterns to simulate burst and steady load

const(
	ProfileBurst Profile = "burst"
	ProfileSteady Profile = "steady"
)

func Generate(profile Profile) []*Job {
	switch profile {
		case ProfileSteady:
			return steadyWorkload()
		default:
			return burstWorkload()
	}
}

func burstWorkload() []*Job{
	return []*Job{
		{ID: "job-1", CPU: 2, Memory: 4, Duration: 5, ArrivalTick: 0},
		{ID: "job-2", CPU: 4, Memory: 4, Duration: 3, ArrivalTick: 0},
		{ID: "job-3", CPU: 6, Memory: 8, Duration: 7, ArrivalTick: 0},
		{ID: "job-4", CPU: 1, Memory: 2, Duration: 2, ArrivalTick: 0},
	}
}

func steadyWorkload() []*Job{
	specs := []struct{
		id string
		cpu int
		memory int
		duration int
	}{
		{"job-1", 2, 4, 5},
		{"job-2", 4, 4, 3},
		{"job-3", 6, 8, 7},
		{"job-4", 1, 2, 2},
	}

	jobs := make([]*Job, len(specs))
	for i, s := range specs {
		jobs[i] = &Job{
			ID: s.id,
			CPU: s.cpu,
			Memory: s.memory,
			Duration: s.duration,
			ArrivalTick: i,
		}
	}
	return jobs
}

func DebugPrint(jobs []*Job) {
	for _, job := range jobs {
		fmt.Printf(" %s arrives at %d, CPU: %d, Memory: %d, Duration: %d\n", job.ID, job.ArrivalTick, job.CPU, job.Memory, job.Duration)
	}
}