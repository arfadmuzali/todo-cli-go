# 📝 Todo CLI App (Golang + Cobra)

A simple, fast, and minimal command-line Todo application written in Go using the [Cobra](https://github.com/spf13/cobra) library. Todo items are saved locally in a JSON file.

## 🚀 Features

- Add new todo items
- View todo list
- Delete todo items by ID
- Update todo content
- Toggle todo status (done/undone)

## 📦 Installation

Make sure your Go version is up to date

   ```bash
   go install github.com/arfadmuzali/todo-cli-go/cmd/todo@latest
   ```

## 📦 Manual Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/arfadmuzali/todo-cli-go.git
   cd todo-cli-go
   ```

2. **Build the project:**

   ```bash
   go build -o todo
   ```

3. **(Optional) Install globally:**

   ```bash
   sudo mv todo /usr/local/bin/
   ```

4. **(Optional) Install locally:**

   ```bash
   sudo mv todo $HOME/.local/bin/
   ```

## 🛠 Usage

```bash
todo [command] [arguments]
```

### 📋 Available Commands

| Command                       | Description                                |
|-------------------------------|--------------------------------------------|
| `todo list`                   | Display all todo items                     |
| `todo add "Task content"`     | Add a new todo item                        |
| `todo delete [id1] [id2]...`  | Delete one or more todos by ID             |
| `todo update [id] "New task"` | Update the content of a todo item by ID    |
| `todo status [id1] [id2]...`  | Toggle todo status (done <--> not done)    |

### 🧪 Example Usage

```bash
todo add "Learn Golang"
todo add "Buy milk"
todo list
todo update 1 "Learn Golang with Cobra"
todo status 1 2
todo delete 2
```

## 🗃 Storage Structure

All todos are stored in a local JSON file. Example `todos.json` content|

```json
[
  {
    "id": 1,
    "task": "Learn Golang",
    "done": true
  },
  {
    "id": 2,
    "task": "Buy milk",
    "done": false
  }
]
```

## 📂 Project Structure

```
.
├── cmd/               # Cobra command implementations (add, list, delete, etc.)
├── internal/          # Core logic and utilities
├── main.go
└── README.md
```

```

~/.config/todo         # Storage
└── todos.json
```
## 🐧 Tips for Linux (i use arch btw)

If you want your todo list to show automatically at startup, you can create a simple script like this:

   ```bash
#!/bin/bash

# File location
FILE="$HOME/.config/todo/todos.json"

# Check if the file exists
if [[ ! -f "$FILE" ]]; then
  echo "File $FILE not found."
  exit 1
fi

# Extract all descriptions from the todo file
descriptions=$(jq -r '.[].Description' "$FILE")

# Exit if there are no todos
if [[ -z "$descriptions" ]]; then
  exit 0
fi

# Send a notification for each todo
while IFS= read -r desc; do
  notify-send "Todo" "$desc"
done <<< "$descriptions"

   ```

