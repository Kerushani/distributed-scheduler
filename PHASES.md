# Distributed Scheduler — v2 Roadmap & Progress

## Project summary

Tick-based Go simulator comparing cluster schedulers (greedy / fit heuristics / DP oracle) on a heterogeneous cluster. Measures makespan, utilization, queue wait, turnaround, fragmentation.

```
Engine → Scheduler → Cluster/Nodes → Jobs
         ↓
       Metrics
```

## Phase status

| Phase | Topic | Status |
|-------|-------|--------|
| **1** | Workloads & queueing (`ArrivalTick`, generator, `loadJobs`, queue wait / turnaround metrics) | **Done** |
| **2** | Heterogeneous cluster & bin packing (first/best/worst-fit, fragmentation) | **In progress** |
| **3** | Real multi-node DP oracle + scheduling efficiency gap | Not started |
| **4** | Scheduler policies (priority, fairness, gang scheduling) | Not started |
| **5** | Failure & preemption | Not started |
| **6** | Observability & experiments (CSV, sweeps, CLI flags) | Not started |

## Phase 1 — done

- `jobs/job.go` — `ArrivalTick` field
- `jobs/generator.go` — `ProfileBurst`, `ProfileSteady`
- `simulator/engine.go` — `loadJobs()` with `nextJobIndex` cursor
- `metrics/metrics.go` — queue wait, turnaround, service time, P95 queue wait
- `main.go` — generator-driven workload, empty `PendingJobs` at start

**Key formulas:**
- Queue wait = `StartTick - ArrivalTick`
- Turnaround = `EndTick - ArrivalTick`
- Service time = `EndTick - StartTick`

## Phase 2 — in progress

**Goal:** Nodes differ in size; compare placement strategies; measure fragmentation.

**Build order:**
1. `cluster/node.go` — `FreeCPU()`, `FreeMemory()`
2. `cluster/cluster.go` — `NodeSpec`, `NewHeterogeneousCluster`
3. `scheduler/fit.go` — shared `placeJobs` + scoring
4. `scheduler/first_fit.go`, `best_fit.go`, `worst_fit.go`
5. `metrics/metrics.go` — fragmentation metric
6. `simulator/engine.go` — pass `PendingJobs` to `Collect`
7. `jobs/generator.go` — `ProfileMixed`
8. `main.go` — heterogeneous cluster + compare 3 schedulers
9. `metrics/comparison.go` — `CompareSchedulers` for N schedulers

**Keep:** `greedy.go` (equivalent to best-fit) and `dp_oracle.go` for later Phase 3 work.

## Phase 3+ (planned)

- **3:** Multi-node DP oracle, scheduling efficiency gap %
- **4:** Priority, fair share, gang scheduling
- **5:** Node failure, draining, preemption
- **6:** Per-tick traces, CSV export, `cmd/experiment`, CLI flags

## Run

```bash
go run .
```
