test: 
	go vet ./...
	go test ./...

test-strings:
	go test -run ReverseStrings ./...

test-arrays:
	go test -run Array ./...

tidyDeps:
	go mod tidy

listDeps:
	go list -m all

print:
	cat Makefile

