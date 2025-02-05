# EPATH=[[(${PWD}|${PWD}/bin)]]

hello:
	echo "hello"

build: 
	go build -o bin/sbx cmd/Sbx/main.go

dev:
	@go run bin/lnk-wgroup.go

run:
	go run cmd/WaitGroups/main.go
