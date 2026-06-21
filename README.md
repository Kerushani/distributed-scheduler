Distributed Scheduler - A Brief Outline

This is a mini distributed job scheduling system that simulates how real cluster schedulers work. It uses a tick-based engine with two scheduler implementations, a greedy scheduler and a DP-based oracle, that can be swapped in for comparison.

The system has two scheduler paths (both run through the same engine):
- Greedy Scheduler -> uses best-fit scoring to pick a node per job
- DP Oracle (Benchmark Path) -> uses DP to find optimal job ordering on a single node (small N only)
Both can be compared to evaluate scheduling quality.

The goal is to simulate a distributed compute cluster and evaluate scheduling efficiency, resource utilization, and job completion time (makespan).

Some terminology:
- Job: Represents a unit of work
- Node: Represents a compute machine
- Cluster: Collection of nodes
- Scheduler (brain):
    - Greedy Scheduler
        - For each pending job, computes a score per node and assigns to best-fit node
    - DP Oracle
        - Takes pending jobs each tick
        - Computes optimal execution order on one node (not multi-node assignment yet)
        - Only for benchmarking purposes
- Engine (orchestrator): Runs the tick loop, applies scheduler decisions, executes jobs, collects metrics

The schedulers are compared on makespan (total completion time), average node utilization %, and average turnaround time.

At each tick:
- Scheduler reads pending jobs from the queue
- Scheduler assigns jobs to nodes
- Running jobs decrement remaining time
- Completed jobs are removed
- Metrics updated

Layers:
- Cluster (infra)
- Scheduler (brain)
- Engine (orchestrator)

Run:
```
go run .
```
