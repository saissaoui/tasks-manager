# Tasks Manager

A simple and efficient CLI tool to organize and manage your tasks. Built with Go and Cobra CLI framework.

## Features

- ✅ Add tasks with customizable details
- 📅 Set due dates for tasks
- 🎯 Assign priority levels
- 🚀 Simple and intuitive command-line interface

## Installation

### Prerequisites
- Go 1.19 or later

### Build from source
```bash
git clone <repository-url>
cd tasks-manager
go build -o tasks-manager
```

## Usage

### Adding Tasks

Use the `add` command to create new tasks with the following options:

```bash
./tasks-manager add -t "Task name" -d "Due date" -p "Priority level"
```

#### Options:
- `-t, --task`: Task name (required)
- `-d, --date`: Due date (required)
- `-p, --priority`: Priority level (required)

#### Examples:

```bash
# Add a high-priority task
./tasks-manager add -t "Complete project proposal" -d "2025-01-15" -p "high"

# Add a medium-priority task
./tasks-manager add -t "Buy groceries" -d "2025-01-10" -p "medium"

# Using long flags
./tasks-manager add --task "Call dentist" --date "2025-01-12" --priority "low"

# Using the alias 'a' for add
./tasks-manager a -t "Review code" -d "2025-01-08" -p "high"
```

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `add`   | `a`   | Add a new task with name, date, and priority |

## Help

Get help for any command:

```bash
./tasks-manager --help
./tasks-manager add --help
```

## Project Structure

```
tasks-manager/
├── cmd/
│   ├── root.go      # Root command configuration
│   └── add.go       # Add command implementation
├── internal/        # Internal packages
├── pkg/            # Public packages
├── main.go         # Application entry point
├── go.mod          # Go module file
└── README.md       # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).
