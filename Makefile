test: 
	go test

tidyDeps:
	go mod tidy

listDeps:
	go list -m all

print:
	cat Makefile

