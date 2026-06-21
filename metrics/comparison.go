package netrics

import "fmt"

func Compare(greedy *Tracker, oracle *Tracker) {
	fmt.Println("==============================")
	fmt.Println("Scheduler Comparison")
	fmt.Println("==============================")

	fmt.Print("%-30s %-12s %-12s\n", "Metric", "Greedy", "Oracle")
	fmt.Print("%-30s %-12s %-12s\n", "Makespan", greedy.Makespan, oracle.Makespan)
	fmt.Print("%-30s %-12s %-12s\n", "Avg Utilization (%)", greedy.AverageUtilization(), oracle.AverageUtilization())
	fmt.Print("%-30s %-12s %-12s\n", "Avg Turnaround Time", greedy.AverageTurnaroundTime(), oracle.AverageTurnaroundTime())

	fmt.Println()
}