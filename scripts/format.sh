#!/usr/bin/env bash
set -euo pipefail
echo "=========================================="
echo " Formatting Go code"
echo "=========================================="
go fmt ./...
go vet ./...
echo "✅ Formatting and vet completed."
