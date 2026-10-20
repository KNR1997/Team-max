// This nested module isolates the skill example files under agent/skills/
// from the root todo module, so `go build ./...`, `go mod tidy`, and
// `golangci-lint run ./...` at the repo root do not scan them.
module skills-assets

go 1.22
