# ASCII Art Web

A command-line tool written in Go that takes a string as input and prints it in a graphic representation using ASCII characters.

## Usage

```bash
go run . "your text here"
```

## Examples

```bash
go run . "Hello"
go run . "Hello\nWorld" # newline between words
go run . "Hello\n\nWorld" # blank line between words
go run . "" # empty string
```

## Features

- Supports uppercase and lowercase letters
- Supports numbers and special characters
- Supports spaces
- Supports \n for newlines in input

## Requirements

- Go 1.18+
- Only standard Go packages are used
