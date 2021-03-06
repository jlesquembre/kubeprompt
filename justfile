VERSION := `git describe --abbrev=0 --tags $(git rev-list --tags --max-count=1)`
VERSION_FLAG := "-X " + `go list ./pkg/version` + ".Version=" + VERSION + "-dev"

build:
  go build \
  -ldflags "-w -s {{VERSION_FLAG}}" \
  -o bin/kubeprompt ./cmd/kubeprompt
