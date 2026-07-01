package metrics

import ("fmt"; "sort")

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

func CompareSchedulers(results map[string]*Tracker) {
	names := make([]string, 0, len(results))
	for name := range results {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Println("==============================")
	fmt.Println("Scheduler Comparison")
	fmt.Println("==============================")
	header := fmt.Sprintf("%-28s", "Metric")
	for _, name := range names {
		header += fmt.Sprintf(" %-14s", name)
	}
	fmt.Println(header)
	printIntRow := func(label string, fn func(*Tracker) int) {
		row := fmt.Sprintf("%-28s", label)
		for _, name := range names {
			row += fmt.Sprintf(" %-14d", fn(results[name]))
		}
		fmt.Println(row)
	}
	printFloatRow := func(label string, fn func(*Tracker) float64) {
		row := fmt.Sprintf("%-28s", label)
		for _, name := range names {
			row += fmt.Sprintf(" %-14.2f", fn(results[name]))
		}
		fmt.Println(row)
	}
	printIntRow("Makespan", func(t *Tracker) int { return t.Makespan })
	printFloatRow("Avg Utilization (%)", func(t *Tracker) float64 { return t.AverageUtilization() })
	printFloatRow("Avg Fragmentation", func(t *Tracker) float64 { return t.AverageFragmentation() })
	printFloatRow("Avg Queue Wait", func(t *Tracker) float64 { return t.AverageQueueWait() })
	printFloatRow("Avg Turnaround", func(t *Tracker) float64 { return t.AverageTurnaroundTime() })
	printFloatRow("Avg Service Time", func(t *Tracker) float64 { return t.AverageServiceTime() })
	fmt.Println()
}