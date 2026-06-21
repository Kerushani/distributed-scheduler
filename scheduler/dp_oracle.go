package scheduler

import (
	"math"

	"distributed-scheduler/cluster"
	"distributed-scheduler/jobs"
)

type DpOracleScheduler struct{}

const MAX_DP_JOBS = 10

func (s *DpOracleScheduler) Schedule(
	_ *cluster.Cluster,
	jobs []*jobs.Job,
	_ int,
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

	parent := make([]int, size)
	for i := range parent {
		parent[i] = -1
	}

	for mask := 0; mask < size; mask++ {
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}

			newMask := mask | (1 << j)
			newCost := dp[mask] + jobs[j].Duration

			if newCost < dp[newMask] {
				dp[newMask] = newCost
				parent[newMask] = j
			}
		}
	}

	order := reconstructOrder(parent, n)

	decisions := []Decision{}
	for _, j := range order {
		decisions = append(decisions, Decision{
			JobID:  jobs[j].ID,
			NodeID: "node-0",
		})
	}

	return decisions
}

func reconstructOrder(parent []int, n int) []int {
	order := []int{}

	mask := (1 << n) - 1

	for mask > 0 {
		j := parent[mask]
		order = append([]int{j}, order...)
		mask = mask & ^(1 << j)
	}

	return order
}
