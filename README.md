# Trackr

Trackr is a minimalist command-line application for logging daily activities, moods, habits, or anything else worth tracking. It uses Go and SQLite to store data locally and securely.

## Features

- Add entries with a category and a value
- Automatically timestamps each log
- Local SQLite database (no network dependency)
- Modular code structure for easy extension

## Example Usage

```bash
go run main.go add mood 7
go run main.go add prayer "Read John 3"
go run main.go add gym "push day"
```

## Project Structure
trackr/
├── cmd/        # Cobra CLI commands
│   ├── add.go
│   └── root.go
├── db/         # SQLite logic
│   └── db.go
├── models/     # Data models
│   └── log.go
├── go.mod
├── main.go
├── README.md


## Installation
Clone the repo
```bash
git clone https://github.com/CBYeuler/trackr.git
cd trackr
```
Install dependencies
```bash
go mod tidy
```
Build or run the app
```bash
go run main.go add test "test entry"
```
Note: This project uses github.com/mattn/go-sqlite3, which requires CGO. To run on Windows, install TDM-GCC and ensure CGO_ENABLED=1.
Alternatively, switch to modernc.org/sqlite for a CGO-free build.

## Planned Features:
trackr list: list logs by category or date

trackr export: export logs to markdown or CSV

Weekly summaries and filtering

TUI/ASCII visualizations

Optional encryption suppor

License
This project is open source and available under the MIT License.