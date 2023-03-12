appname = goapple

Default:
	go run main.go
start:
	./bin/$(appname)
build:
	go build -o bin/$(appname) main.go