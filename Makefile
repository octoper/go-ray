build_playground: clean_playground
	go build -o playground/bin/main playground/playground.go

playground:
	go run playground/playground.go

clean_playground:
	rm -R playground/bin/
