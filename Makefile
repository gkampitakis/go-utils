test: 
	go vet ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	
test-strings:
	go test -run ReverseStrings ./...

test-arrays:
	go test -run Array ./...

tidyDeps:
	go mod tidy

listDeps:
	go list -m all

print:
	@cat Makefile

clean: 
	rm coverage.out

getVersion:
	@tail -n 1 go.mod |  cut -f2 -d":" | cut -c2-