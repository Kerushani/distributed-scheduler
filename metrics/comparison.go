package metrics

import "fmt"

func Compare(greedy *Tracker, oracle *Tracker) {
	fmt.Println("==============================")
	fmt.Println("Scheduler Comparison")
	fmt.Println("==============================")

	fmt.Printf("%-30s %-12s %-12s\n", "Metric", "Greedy", "Oracle")
	fmt.Printf("%-30s %-12d %-12d\n", "Makespan", greedy.Makespan, oracle.Makespan)
	fmt.Printf("%-30s %-12.2f %-12.2f\n", "Avg Utilization (%)", greedy.AverageUtilization(), oracle.AverageUtilization())
	fmt.Printf("%-30s %-12.2f %-12.2f\n", "Avg Queue Wait", greedy.AverageQueueWait(), oracle.AverageQueueWait())
	fmt.Printf("%-30s %-12.2f %-12.2f\n", "Avg Turnaround Time", greedy.AverageTurnaroundTime(), oracle.AverageTurnaroundTime())
	fmt.Printf("%-30s %-12.2f %-12.2f\n", "Avg Service Time", greedy.AverageServiceTime(), oracle.AverageServiceTime())

	fmt.Println()
}
