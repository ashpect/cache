# Cache
A cache implementation in Go. The cache is thread safe and supports multiple policies.
You can write your own policy by implementing your logic as per the boilerplate provided in main.go. To read examples, refer to main.go.
The main cache storage is a map[string]string and methods, builder and director is provided for it. You will find common algorithms in handlers. An obvious improvement is to use go generics instead of string but since I have implemented methods, it will be dirty to use go generics.

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
- [x] Use go generics or interfaces instead of string for generic cache

### Misc
- [x] Implement example in main.go 
- [ ] Add tests for testing thread safety
