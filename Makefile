test:
	go test -count=1000 .

benchmark: benchmark-marshal benchmark-unmarshal benchmark-copy
	@cat benchmark_marshal.txt benchmark_unmarshal.txt benchmark_copy.txt > benchmark.txt

benchmark-marshal:
	go test -benchmem -bench '^(Benchmark\w+Marshal)$$' . | tee benchmark_marshal.txt

benchmark-unmarshal:
	go test -benchmem -bench '^(Benchmark\w+Unmarshal)$$' . | tee benchmark_unmarshal.txt

benchmark-copy:
	go test -benchmem -bench '^(Benchmark\w+Copy)$$' . | tee benchmark_copy.txt

generate:
	protoc --go_out=. --go_opt=paths=source_relative ./pb/product.proto
