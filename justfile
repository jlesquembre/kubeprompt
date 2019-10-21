VERSION_FLAG := "-X " + `go list ./pkg/version` + ".Version=1.0.0"

build:
  go build \
  -ldflags "-w -s {{VERSION_FLAG}}" \
  -o bin/kubeprompt ./cmd/kubeprompt
