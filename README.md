# 🐹 Go (Golang) Learning Journey

This repository tracks my hands-on learning progress with **Go (Golang)**. The goal is to document core language concepts, leverage my background in Python to draw practical comparisons, and adopt Go's idiomatic design patterns.

## 🚀 Hands-On Project: TodoList CLI

The initial exercise is a lightweight **Command Line Interface (CLI)** To-Do list application.

This project covers foundational concepts such as working with collection types, string manipulation, and functional modularization.

### 💡 Key Technical Concepts Applied

* **Slices (`[]string`):** Utilizing dynamic collections to store and organize task data.
* **Standard `strings` Package:**
  * `strings.Repeat`: Constructing structured ASCII header borders.
  * `strings.ToUpper`: Transforming string values to uppercase for formatting.
* **Iterative Loops (`for ... range`):** Traversing slices with zero-indexed control.
* **Custom Functions:** Creating `printTasks` to encapsulate and reuse presentation logic.
* **Formatted I/O (`fmt.Printf` / `fmt.Println`):** Clean stdout formatting with format specifiers.

## 📜 Source Code (`main.go`)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	opcoes := []string{
		"Watch Go Crash Course", 
		"Watch Nana's Golang Full Course", 
		"Reward myself with a cheesecake",
	}

	fmt.Println(strings.Repeat("=", 20) + strings.ToUpper("| Welcome to our TodoList App! |") + strings.Repeat("=", 20))

	// Calling printTasks to render options in a structured format
	printTasks(opcoes)

	fmt.Println("\nTutoriais:")
	fmt.Println(opcoes[0])
	fmt.Println(opcoes[1])

	fmt.Println("\nRewards:")
	fmt.Println(opcoes[2])

	fmt.Println("\nMy Projects:")
	fmt.Println(opcoes[1])
}

// Function to render tasks in a structured list format
func printTasks(task []string) {
	fmt.Println("\nList of my Todos:")
	for i, opcao := range task {
		fmt.Printf("%d. %s\n", i+1, strings.ToUpper(opcao))
	}
}
```

## 🛠️ How to Run

### Prerequisites

* Go 1.20+ installed on your system (Linux, macOS, or Windows).

### Local Execution

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   ```

2. **Initialize the Go module:**
   ```bash
   go mod init project_1
   ```

3. **Run the application:**
   ```bash
   go run main.go
   ```

## 🖥️ Expected Terminal Output

```text
====================| WELCOME TO OUR TODOLIST APP! |====================

List of my Todos:
1. WATCH GO CRASH COURSE
2. WATCH NANA'S GOLANG FULL COURSE
3. REWARD MYSELF WITH A CHEESECAKE

Tutoriais:
Watch Go Crash Course
Watch Nana's Golang Full Course

Rewards:
Reward myself with a cheesecake

My Projects:
Watch Nana's Golang Full Course
```

## 📌 Technical Roadmap & Future Improvements

* [ ] Implement interactive user input (`fmt.Scan` / `bufio`).
* [ ] Add dynamic task mutations (Create/Delete) using `append` and slice slicing.
* [ ] Implement persistence using JSON/flat file storage.
* [ ] Refactor tasks into structured types (`structs`) with state management (pending/completed).