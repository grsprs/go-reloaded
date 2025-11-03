@echo off
echo Restoring original golden tests...

REM Backup current implementation
copy "tests\golden_test.go" "tests\golden_test_new.go.bak"
copy "tests\expected_outputs\test3.txt" "tests\expected_outputs\test3_new.txt.bak"
copy "internal\processor\case.go" "internal\processor\case_new.go.bak"

REM Restore from git or backup (you'll need to implement this based on your version control)
echo To restore original tests, you need to:
echo 1. Restore tests\golden_test.go from git history
echo 2. Restore tests\expected_outputs\test3.txt from git history  
echo 3. Restore internal\processor\case.go from git history
echo.
echo Current implementation backed up as *.bak files
echo.
echo Use: git checkout HEAD~1 -- tests/golden_test.go tests/expected_outputs/test3.txt internal/processor/case.go
pause