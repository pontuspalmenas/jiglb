# jiglb

Toy LB project.

## Planned features
- L7 multi-target group load balancing
- HTTPS with TLS termination
- Path based routing
- JWT validation
- Client certificate validation (mTLS)

## Flow
1. Accept connection from client (C).
1. Inspect HTTP data (headers, path, etc.) to use for routing.
1. Find the appropriate target group (TG) based on routing criteria (e.g. path).
1. Pick target (T) from TG using load balancing algorithm.
1. Connect to T, relay response to C.

## Slask
- We can think of the problem as two parts:
- Implement a proxy
  - [ ] Timeouts
  - [ ] Concurrency
  - [ ] Robust error handling
- Implement a load balancer
  - [ ] Radix tree
  - [ ] Algorithm
  - [ ] Target groups, targets
  - [ ] Track state for distributing load

## Todo
- [ ] Plugins
- 
