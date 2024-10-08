test:
	go test -count=1000 .

benchmark:
	@go test -benchmem -bench ./... | tee benchmark.txt

benchmark-marshal:
	@go test -benchmem -bench '^(BenchmarkMarshal)$$' . | tee benchmark_marshal.txt

benchmark-unmarshal:
	@go test -benchmem -bench '^(BenchmarkUnmarshal)$$' . | tee benchmark_unmarshal.txt

benchmark-copy:
	@go test -benchmem -bench '^(BenchmarkCopy)$$' . | tee benchmark_copy.txt

generate:
	protoc --go_out=. --go_opt=paths=source_relative ./pb/product.proto
