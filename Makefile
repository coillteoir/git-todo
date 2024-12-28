git-todo: main.go cmd/*.go
	go fmt 
	go build 

run: git-todo
	go run .
