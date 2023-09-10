benchmark:
	go test ./internal/benchmark --bench=. --benchtime 10s
	go test ./internal/benchmark --bench=. --benchmem
