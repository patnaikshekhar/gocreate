.PHONY: build
build:
	go build -o ./bin/{{ .ProjectName }}

.PHONY: run
run: build
	./bin/{{ .ProjectName }}