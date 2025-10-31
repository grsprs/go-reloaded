@echo off
echo ========================================
echo Go Reloaded - Live Demonstration
echo ========================================
echo.

echo [1] Input File:
echo ----------------------------------------
type demo\input.txt
echo.
echo.

echo [2] Running Transformation...
go run ./cmd/go-reloaded demo/input.txt demo/output.txt
echo.

echo [3] Output File:
echo ----------------------------------------
type demo\output.txt
echo.
echo.

echo [4] Test Suite:
echo ----------------------------------------
go test ./tests -v | findstr "PASS\|FAIL"
echo.

echo ========================================
echo Demonstration Complete!
echo ========================================