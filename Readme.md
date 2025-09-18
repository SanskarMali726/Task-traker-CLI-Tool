# 📝 Task CLI (Go)

A simple **Command Line Interface (CLI) task manager** built with **Go**.  
This tool lets you **add, update, delete, list, and manage tasks** with statuses such as `in-progress` and `done`.  
Tasks are stored locally in a JSON file, so no external database is required.

---

## 🚀 Features
- ➕ Add a new task  
- ✏️ Update an existing task  
- ❌ Delete a task  
- 🔄 Mark tasks as **in-progress** or **done**  
- 📋 List all tasks  
- 🎯 List tasks filtered by status (`done`, `in-progress`, etc.)  

---
## Inspired By / Resources

This project was inspired by the [Task Tracker project roadmap](https://roadmap.sh/projects/task-tracker) on roadmap.sh.

---

## Getting Started

### Prerequisites
- Go 1.24.0 or above
- Windows, Linux, or Mac

### Build & Run
Build the CLI executable:

```bash
go build -o task-cli main.go

##Run commands from the terminal:
# Add a task
./task-cli add "Finish project"
# Update a task
./task-cli update 1 "Finish Go CLI project"
# Delete a task
./task-cli delete 1
# Mark a task as in progress
./task-cli mark-in-progress 1
# Mark a task as done
./task-cli mark-done 1
# List tasks
./task-cli list
# List tasks by status
./task-cli list done
./task-cli list in-progress
./task-cli list to
```
---