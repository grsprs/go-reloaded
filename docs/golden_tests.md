# Go Reloaded – Golden Test Set

The following test cases describe how the program is expected to behave once implemented.  
They are expressed in natural language to illustrate understanding of the transformation rules.

---

## Test 1 – Number Conversion
**Input:**  
`Simply add 1E (hex) and 10 (bin) and you will see the result is 68.`

**Expected Output:**  
`Simply add 30 and 2 and you will see the result is 68.`

---

## Test 2 – Multi-word Casing
**Input:**  
`This is so exciting (up, 2)!`

**Expected Output:**  
`This is SO EXCITING!`

---

## Test 3 – Punctuation Formatting
**Input:**  
`Punctuation tests are ... kinda boring ,what do you think ?`

**Expected Output:**  
`Punctuation tests are... kinda boring, what do you think?`

---

## Test 4 – Quotation Formatting
**Input:**  
`As Elton John said: ' I am the most well-known homosexual in the world '`

**Expected Output:**  
`As Elton John said: 'I am the most well-known homosexual in the world'`

---

## Test 5 – Grammar Correction
**Input:**  
`There is no greater agony than bearing a untold story inside you.`

**Expected Output:**  
`There is no greater agony than bearing an untold story inside you.`

---

## Test 6 – Combined Rules
**Input:**  
`A honest man said: ' I saw 1E (hex) cats ,and 10 (bin) dogs (up, 3) !'`

**Expected Output:**  
`An honest man said: 'I saw 30 cats, and 2 DOGS!'`

---

## Test 7 – Complex Paragraph
**Input:**  
`it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.`

**Expected Output:**  
`It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.`

---

## General Validation Rule
If the output matches the examples above in structure, spacing, and logic, the transformation process can be considered functionally correct.
