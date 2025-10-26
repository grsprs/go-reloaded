#!/usr/bin/env bash
set -euo pipefail
echo "Run `go test ./...` to execute tests"
go test ./...
