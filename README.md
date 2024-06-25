# Cache
A cache implementation in Go. The cache is thread safe and supports multiple policies.
You can write your own policy by implementing your logic as per the boilerplate provided in main.go. To read examples, refer to main.go.

## TODO :

### Pkg
- [x] Implement a builder for cache
- [x] Implement standard policies
    - [x] LRU
    - [x] LIFO
    - [x] FIFO
    - [x] LFU
    - [x] MRU
- [x] Add thread safety to cache (Not using sync.map because writes are frequent)
- [ ] Use go generics instead of string 

### Misc
- [x] Implement example in main.go 
- [ ] Add test for thread safety for learning purposes
