# 📝 Todo CLI App (Golang + Cobra)

A simple, fast, and minimal command-line Todo application written in Go using the [Cobra](https://github.com/spf13/cobra) library. Todo items are saved locally in a JSON file.

## 🚀 Features

- Add new todo items
- View todo list
- Delete todo items by ID
- Update todo content
- Toggle todo status (done/undone)

## 📦 Installation

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
| `todo toggle [id1] [id2]...`  | Toggle todo status (done <--> not done)    |

### 🧪 Example Usage

```bash
todo add "Learn Golang"
todo add "Buy milk"
todo list
todo update 1 "Learn Golang with Cobra"
todo toggle 1 2
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

## 🔧 Dependencies

- [Cobra](https://github.com/spf13/cobra)

Install with:

```bash
go get github.com/spf13/cobra
```

- [TableWriter](github.com/olekukonko/tablewriter)

Install with:

```bash
go get github.com/olekukonko/tablewriter
```

## ✍️ Contribution

Pull requests and suggestions are very welcome! Feel free to fork and add new features such as:

- Priorities or deadlines
- Categories/tags
- Filtering and searching
- Storage using SQLite or other databases

---

## 📣 Final Notes

This app is great for terminal lovers who want a fast and lightweight todo manager.
