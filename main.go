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

func main() {
	workload := []*jobs.Job{
		{
			ID:       "job-1",
			CPU:      2,
			Memory:   4,
			Duration: 5,
		},
		{
			ID:       "job-2",
			CPU:      4,
			Memory:   4,
			Duration: 3,
		},
		{
			ID:       "job-3",
			CPU:      6,
			Memory:   8,
			Duration: 7,
		},
		{
			ID:       "job-4",
			CPU:      1,
			Memory:   2,
			Duration: 2,
		},
	}

	greedyCluster := cluster.NewCluster(3, 8, 16)

	greedyEngine := simulator.Engine{
		Cluster:     greedyCluster,
		Jobs:        cloneJobs(workload),
		PendingJobs: cloneJobs(workload),
		Scheduler:   &scheduler.GreedyScheduler{},
		Metrics:     metrics.NewTracker(),
	}

	fmt.Println("Running Greedy Scheduler...")
	greedyEngine.Run(100)

	greedyEngine.Metrics.PrintReport("Greedy")

	oracleCluster := cluster.NewCluster(3, 8, 16)

	oracleEngine := simulator.Engine{
		Cluster:     oracleCluster,
		Jobs:        cloneJobs(workload),
		PendingJobs: cloneJobs(workload),
		Scheduler:   &scheduler.DpOracleScheduler{},
		Metrics:     metrics.NewTracker(),
	}

	fmt.Println("Running Oracle Scheduler...")
	oracleEngine.Run(100)

	oracleEngine.Metrics.PrintReport("Oracle")

	metrics.Compare(
		greedyEngine.Metrics,
		oracleEngine.Metrics,
	)
}
