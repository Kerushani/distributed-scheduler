Distributed Scheduler - A Breif Outline

This is a mini distributed job scheduling system that simulates how real cluster schedulers work. It uses a greedy runtime engine and a DP-based oracle for evaluation.

The system will have two paths:
- Real Time Scheduler -> which will use Greedy Scoring + Bin Packing Heuristic
- Offline Optimizer (Benchmark Path): DP-based
Both can be compared to evaluate scheduling quality.

The goal is to simulate a distributed compute cluster and evaluate scheduling efficiency, resource utilization, job completion time (makespan), and queue latency. 

Some terminology:
- Job: Represents a unit of work
- Node: Represents a compute machine
- Cluster: Collection of nodes
- Scheduler (Core Engine): 
    - Greedy Scheduler 
        - For each incoming jobs, it will compute score per node and assign to best-fit node
    - DP optimizer
        - Takes a batch of jobs
        - Computes optimal assignment (small N only)
        - Only for benchmarking purposes

The schedulers will be compared on their makespan (total completion time), node utilization %, job waiting time, and scheduling efficiency gap.

At each tick:
- Scheduler receives new jobs (or pulls from queue)
- Greedy scheduler assigns jobs
- Running jobs decrement remaining time
- Completed jobs are removed
- Metrics updated

Layers:
- Cluster (infra)
- Scheduler (brain)
- Engine (orchestrator)