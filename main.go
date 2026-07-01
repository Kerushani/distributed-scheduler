package main

import (
	"fmt"

	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
	"distributed-scheduler/metrics"
	"distributed-scheduler/scheduler"
	"distributed-scheduler/simulator"
)

func cloneJobs(original []*jobs.Job) []*jobs.Job {
	clones := []*jobs.Job{}

	for _, job := range original {
		copy := *job
		clones = append(clones, &copy)
	}

	return clones
}

func runSimulation(name string, sched scheduler.Scheduler, specs []cluster.NodeSpec, workload []*jobs.Job) *metrics.Tracker {
	c := cluster.NewHeterogeneousCluster(specs)

	engine := simulator.Engine{
		Cluster: c, 
		Jobs: cloneJobs(workload),
		PendingJobs: nil,
		Scheduler: sched,
		Metrics: metrics.NewTracker(),
	}

	fmt.Printf("Running %s...\n", name)
	engine.Run(100)
	engine.Metrics.PrintReport(name)

	return engine.Metrics
}
func main() {
	profile := jobs.ProfileMixed

	workload := jobs.Generate(profile)

	fmt.Printf("Workload profile: %s\n", profile)
	jobs.DebugPrint(workload)
	fmt.Println()

	specs := []cluster.NodeSpec{
		{CPU:8, Memory:16},
		{CPU:4, Memory: 8},
		{CPU: 2, Memory: 16},
	}

	firstFit := runSimulation("First Fit", &scheduler.FirstFitScheduler{}, specs, workload)
	bestFit := runSimulation("Best Fit", &scheduler.BestFitScheduler{}, specs, workload)
	worstFit := runSimulation("Worst Fit", &scheduler.WorstFitScheduler{}, specs, workload)

	metrics.CompareSchedulers(map[string]*metrics.Tracker{
		"First-Fit": firstFit,
		"Best-Fit":  bestFit,
		"Worst-Fit": worstFit,
	})
}