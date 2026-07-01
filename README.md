Distributed Scheduler - A Brief Outline

Refer to [PHASES.md](PHASES.md) for the full v2 roadmap and progress.

This is a mini distributed job scheduling system that simulates how real cluster schedulers work. A tick-based engine runs workloads over discrete time steps, assigns jobs to nodes, and collects metrics. Schedulers are swappable behind a common interface.

**Phase 1 (done):** Jobs arrive over time via workload profiles (burst, steady). Metrics track queue wait, turnaround, and service time.

**Phase 2 (done):** Clusters can be heterogeneous (nodes with different CPU/memory). Three bin-packing schedulers — first-fit, best-fit, worst-fit — are compared on a mixed workload. Fragmentation measures stranded capacity when pending jobs cannot use free resources on a node.

The current `main.go` runs first-fit, best-fit, and worst-fit on a heterogeneous 3-node cluster with `ProfileMixed`. A DP oracle (`dp_oracle.go`) remains for Phase 3 benchmarking.

The goal is to simulate a distributed compute cluster and evaluate scheduling efficiency, resource utilization, job completion time (makespan), and placement quality.

Some terminology:
- Job: Represents a unit of work
- Node: Represents a compute machine
- Cluster: Collection of nodes
- Scheduler (brain):
    - First-Fit / Best-Fit / Worst-Fit — bin-packing placement on heterogeneous nodes
    - Greedy Scheduler — equivalent to best-fit (kept for compatibility)
    - DP Oracle — single-node ordering benchmark (Phase 3 will add multi-node assignment)
- Engine (orchestrator): Runs the tick loop, applies scheduler decisions, executes jobs, collects metrics

The schedulers are compared on makespan, average node utilization %, queue wait, turnaround, service time, and fragmentation.

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

Current Output:
<img width="383" height="336" alt="Screenshot 2026-06-20 at 11 54 06 PM" src="https://github.com/user-attachments/assets/b98db85b-dc38-4655-ac27-ced657d24aff" />

