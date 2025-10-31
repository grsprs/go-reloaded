# Golden Test Suite – Go Reloaded

**Author:** Spiros Nikoloudakis  
**Date:** September 1, 2025  
**Version:** 1.0

## Overview
Official validation tests for the Go Reloaded text transformation engine. Each test verifies specific transformation rules and edge cases.

## Test Execution
```bash
go test ./tests -v
go run ./cmd/go-reloaded input.txt output.txt
```

---

## Core Validation Tests

### Test 1 – Number System Conversions
**Input:**  
`Add 1A (hex) and 1010 (bin) to get the result.`  
**Expected Output:**  
`Add 26 and 10 to get the result.`  
**Purpose:** Validate hexadecimal and binary numeric conversions.

---

### Test 2 – Basic Case Transformations
**Input:**  
`make this IMPORTANT (low) and that small (up)`  
**Expected Output:**  
`make this important and that SMALL`  
**Purpose:** Ensure correct single-word case transformations.

---

### Test 3 – Multi-Word Case Operations
**Input:**  
`transform these three words (cap, 3) properly`  
**Expected Output:**  
`Transform These Three words properly`  
**Purpose:** Validate batch case transformation logic with word counts.

---

### Test 4 – Punctuation Formatting
**Input:**  
`Hello , world ... how are you today ?`  
**Expected Output:**  
`Hello, world... how are you today?`  
**Purpose:** Test punctuation spacing and grouping consistency.

---

### Test 5 – Quote Cleaning
**Input:**  
`He said ' I am ready ' and left ' immediately '`  
**Expected Output:**  
`He said 'I am ready' and left 'immediately'`  
**Purpose:** Verify proper quote spacing cleanup.

---

### Test 6 – Article Correction
**Input:**  
`I saw a apple and a hour later a university.`  
**Expected Output:**  
`I saw an apple and an hour later a university.`  
**Purpose:** Ensure contextual “a/an” grammatical correction.

---

## Edge Case Tests

### Test 7 – Invalid Number Handling
**Input:**  
`Convert GH (hex) and 23 (bin) should remain.`  
**Expected Output:**  
`Convert GH (hex) and 23 (bin) should remain.`  
**Purpose:** Verify that invalid numeric patterns remain unchanged.

---

### Test 8 – Excessive Word Count
**Input:**  
`only two words here (up, 5) total`  
**Expected Output:**  
`ONLY TWO words here total`  
**Purpose:** Confirm correct behavior when modifier count exceeds word limit.

---

### Test 9 – Complex Punctuation
**Input:**  
`Wait ... really ! ? , ; : all together !!`  
**Expected Output:**  
`Wait... really!?,;: all together!!`  
**Purpose:** Validate multi-symbol punctuation clustering.

---

### Test 10 – Nested Transformations
**Input:**  
`1F (hex) (up) becomes uppercase number`  
**Expected Output:**  
`31 becomes uppercase number`  
**Purpose:** Confirm sequential transformation handling within the same token.

---

### Test 11 – Mixed Transformations
**Input:**  
`' QUOTED TEXT (low) ' needs formatting , correctly !`  
**Expected Output:**  
`'quoted text' needs formatting, correctly!`  
**Purpose:** Validate combined quote and punctuation handling.

---

## Integration Test

### Test 12 – Comprehensive Example
**Input:**  
`Start with 1F (hex) and apply (up, 2) transformations. Format ' this quote ' and handle a hour correctly. Don't forget punctuation , like this ... should work !`  

**Expected Output:**  
`Start with 31 and APPLY TRANSFORMATIONS. Format 'this quote' and handle an hour correctly. Don't forget punctuation, like this... should work!`  

**Purpose:** Full integration validation of all transformation stages working together.

---

## Validation Criteria
- All transformations applied **in correct logical order**
- Proper **spacing and punctuation formatting**
- Accurate **numeric conversions**
- Consistent **text case modification**
- Correct **quote and article handling**
- Original **semantic meaning preserved**

## Test Categories
- **Numeric**: Hex/binary conversions (Tests 1, 7, 10)
- **Case**: Single/multi-word transformations (Tests 2, 3, 8)
- **Formatting**: Punctuation and quotes (Tests 4, 5, 9, 11)
- **Grammar**: Article corrections (Test 6)
- **Integration**: Combined transformations (Test 12)

## Validation Requirements
- All 12 tests must pass without deviation
- Output must match expected results exactly
- Performance: <1s for all test inputs combined

## Implementation Status
✅ **All 12 tests passing** (100% compliance)
- Last validated: Implementation complete
- Test execution time: <1s
- Zero regressions detected
