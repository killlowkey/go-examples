.PHONY: run
run:
	go build -gcflags "all=-N -l"
	dlv --listen=:3000 --headless=true --api-version=2 --accept-multiclient exec ./delve

.PHONY: compile
compile:
	go build -gcflags "all=-N -l"

.PHONY: clean
clean:
	rm -rf delve.exe

.PHONY: insatll
install:
	sudo apt install inotify-tools
	go install github.com/go-delve/delve/cmd/dlv@latest
	go mod tidy