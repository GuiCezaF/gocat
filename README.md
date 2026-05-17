# GOCAT

`gocat` is a small Go command-line tool for inspecting text files from the terminal.

This project is a work in progress (WIP). The current implementation is intentionally minimal and will likely grow with more file inspection features over time.

It currently provides two commands:

- `read` prints the contents of a file
- `countLines` counts the number of lines in a file

## Requirements

- Go 1.26.3 or newer

## Installation

```bash
git clone https://github.com/GuiCezaF/gocat.git
cd gocat
go build -o gocat
```

You can also run it directly with Go:

```bash
go run . --help
```

## Usage

### Read a file

```bash
./gocat read path/to/file.txt
```

Example:

```bash
./gocat read README.md
```

### Count lines in a file

```bash
./gocat countLines path/to/file.txt
```

Example:

```bash
./gocat countLines README.md
```

## Commands

### `read [file_path]`

Reads the file and prints its contents to standard output.

### `countLines [file_path]`

Counts the total number of lines in the file and prints the result.

## Project Structure

```text
.
├── main.go
├── cmd/
│   ├── root.go
│   ├── read.go
│   └── countLines.go
└── internal/
    └── reader/
        └── file_reader.go
```

## Notes

- The current CLI expects a valid file path argument for each command.
- Error handling is minimal and the command exits with a non-zero status if file access fails.

## License

See [LICENSE](LICENSE).
