Implementations of varius useful data structures in Go.


Currently includes:
* Min Heap
* Max Heap


## Testing

To run tests for entire project
```sh
go test ./...
```

## Coverage

To show coverage report for entire project
```sh
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```