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
	profile := jobs.ProfileBurst

	workload := jobs.Generate(profile)

	fmt.Printf("Workload profile: %s\n", profile)
	jobs.DebugPrint(workload)
	fmt.Println()
	
	jobs.DebugPrint(workload)
	greedyCluster := cluster.NewCluster(3, 8, 16)

	greedyEngine := simulator.Engine{
		Cluster:     greedyCluster,
		Jobs:        cloneJobs(workload),
		PendingJobs: nil,
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
		PendingJobs: nil,
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