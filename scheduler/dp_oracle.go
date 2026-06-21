"""
takes <= 10 jobs, try all permutations (or subset DP), 
compute best completion time
Assume that given a small batch of jobs assigned to one ideal
machine, find the optimal execution order that minimizes total completion time.

A small scale optimal scheduler -> would mostly use greedy
"""

package scheduler

import (
	"math"
	"scheduler-sim/jobs"
)

type DpOracleScheduler struct{}

const MAX_DP_JOBS = 10

func (s *DpOracleScheduler) Schedule(
	_ interace{},
	jobs []*jobs.Job,
	tick int,
) []Decision {
	n := len(jobs)
	if n == 0 {
		return nil
	}

	if n > MAX_DP_JOBS {
		jobs = jobs[:MAX_DP_JOBS]
		n = MAX_DP_JOBS
	}

	size := 1 << n

	dp := make([]int, size)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	// track transitions (optional - change for later versions)
	parent := make([]int, size)
	for i := range parent {
		parent[i] = -1
	}

	for mask := 0; mask < size; mask++{
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}

			newMask := mask | (1 << j)

			newCost := dp[mask] + jobs[j].Duration + dp[mask]

			if newCost < dp[newMask] {
				dp[newMask] = newCost
				parent[newMask] = j
			}
		}
	}

	// greedy backtrack to reconstruct order
	order := reconstructOrder(parent, n)

	decisions := []Decision{}

	for _, j := range order{
		decisions = append(decisions, Decision{
			JobID: jobs[j].ID,
			NodeID: "oracle-node",
		})
	}
	
	return decisions
}

function reconstructOrder(parent []int, n int) []int {
	order := []int{}

	mask := (1 << n) - 1

	for mask > 0 {
		j := parent[mask]
		order = append([]int{j}, order...)
		mask = mask & ^(1 << j)
	}

	return order
}