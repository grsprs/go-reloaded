# Go Reloaded – Analysis and Design Document

## 1. Problem Description
The objective of **Go Reloaded** is to design a text processing tool that can read a text file, detect transformation indicators (such as `(hex)` or `(cap)`), and produce a reformatted output.  
This task emphasizes the analysis and system design behind such transformations, not the implementation itself.

---

## 2. Transformation Rules

### 🔢 Number Conversion
- Replace hexadecimal numbers before `(hex)` with their decimal equivalents.  
  Example: `1E (hex)` → `30`
- Replace binary numbers before `(bin)` with their decimal equivalents.  
  Example: `10 (bin)` → `2`

---

### 🔠 Case Transformations
- `(up)` → converts the previous word to uppercase.  
- `(low)` → converts the previous word to lowercase.  
- `(cap)` → capitalizes the previous word.  
- `(up, n)`, `(low, n)`, `(cap, n)` → apply the transformation to *n* previous words.

Example:  
`This is so exciting (up, 2)` → `This is SO EXCITING`

---

### ✍️ Punctuation Formatting
- Punctuation marks (`.,!?;:`) should attach directly to the preceding word and be followed by a space before the next word.  
- Multiple punctuation groups (`!?`, `...`) should remain together.  

Example:  
`I was sitting over there ,and then BAMM !!` → `I was sitting over there, and then BAMM!!`

---

### 🧩 Quotation Handling
- Single quotes (`'`) should be placed directly around the quoted text.  
Example: `' I am happy '` → `'I am happy'`

---

### 📝 Article Correction
- Replace `a` with `an` when followed by a word beginning with a vowel or an “h”.  
Example: `a honest story` → `an honest story`

---

## 3. Architectural Comparison

### 🏗️ Pipeline Model
Each transformation rule is handled in a separate, sequential processing step.  
**Advantages:**
- Modular structure  
- Easy testing per stage  
- Readable and maintainable  

### ⚙️ Finite-State Machine (FSM)
Processes the text in a single pass, changing behavior based on the current state.  
**Advantages:**
- Efficient  
- Compact logic  

**Disadvantages:**  
- More complex to debug and modify.

---

## 4. Architectural Choice
The **Pipeline** model is selected as the preferred architecture.  
It allows clearer separation of logic, easier audits, and independent testing of each transformation stage.

---

## 5. Summary
**Go Reloaded** focuses on text clarity and structured processing.  
The project highlights logical consistency, architectural reasoning, and the ability to document thought processes effectively before implementation.
