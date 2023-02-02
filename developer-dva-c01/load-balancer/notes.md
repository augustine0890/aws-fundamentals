# AWS Load Balancers

## Application Load Balancer (ALB)
- Elastic Load Balancing automaticallt distributes your incoming traffic across multiple targets, such as EC2 instances, containers, and IP addresses, in one or multiple Availability Zones.
- It monitors the health of its registered targets and routes traffic only to healthy targets.
- Elastic Load Balancing scales your load balancer as your incomming traffic changes over time. It can automatically scale to the vast majority of workloads.

## ALB Components
- A __load balancer__ serves as the single point of contact for clients. You add one or more listeners to your load balancer.
- A __listener__ checks for connection requests from client, using the protocol and port that you configure.
  - The rules that you define for a listener determine how the load balancer routes request to its registered targets.
  - You must define a default rule for each listener, and you can optionally define additional rules.
- Each __target group__ routes requests to one or more registered targets, such as EC2 instances, using the protocol and port number that you specify.
  - You can configure health checks on a per target group basis. Health checks are performed on all targets registered to a target group that is specified in a listener rule for your load balancer.

<p align="center" width="100%">
    <img width="90%" src="/developer-dva-c01/load-balancer/load-balancer.png">
</p>
