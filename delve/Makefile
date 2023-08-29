.PHONY: run
run:
	go build -gcflags "all=-N -l"
	dlv --listen=:3000 --headless=true --api-version=2 --accept-multiclient exec ./delve.exe

.PHONY: compile
compile:
	go build -gcflags "all=-N -l"

.PHONY: clean
clean:
	rm -rf delve.exe