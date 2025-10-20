# Go Reloaded

### A Go-based Text Transformation Project

**Go Reloaded** is a project focused on text analysis, transformation, and formatting automation.  
The objective is to design a simple but structured command-line tool that can process a text file, apply specific transformation rules, and produce a corrected output file.

---

## 🎯 Objective
The project aims to explore the logic of text manipulation in Go.  
Participants analyze a set of language rules, design a clear architecture, and define functional test cases.  
The emphasis is on *analysis, design, and clarity of logic*, not on rapid coding.

---

## 💡 Learning Goals
- Understanding how textual data can be processed programmatically.  
- Applying structured architectural thinking (Pipeline or FSM models).  
- Practicing documentation and peer collaboration through audits.  
- Designing a clean, modular, and testable approach in Go.

---

## 🧱 Core Description
The tool will read an **input file**, apply a series of text transformations, and write the **output** to another file.  
Transformations include:
- Hexadecimal and binary number conversion  
- Text case modification (upper, lower, capitalize)  
- Grammar-based article correction (`a` → `an`)  
- Punctuation and quotation cleanup

Each rule contributes to producing a clean and grammatically consistent text.

---

## 🧭 Process Overview
1. **Analysis Phase** – define and document the problem clearly.  
2. **Architecture Choice** – compare Pipeline and FSM models, justify your selection.  
3. **Golden Test Definition** – create a set of natural language test cases that describe correct behavior.  
4. **Audit Phase** – exchange feedback with peers to ensure mutual understanding.

---

## ⚙️ Example Usage (after implementation)
```bash
$ go run . input.txt output.txt
```

The program should process the text according to the defined transformation rules and output the corrected version.

---

## ✨ Purpose
The project combines structured thinking, collaboration, and technical documentation.  
It helps participants strengthen their understanding of text processing, while improving clarity, precision, and communication — core skills in any software development context.
