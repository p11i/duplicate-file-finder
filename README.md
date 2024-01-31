# Duplicate File Finder

A simple command-line tool written in Go to find duplicate files in a specified directory based on their hash.

## Features

- Identifies duplicate files based on their hash.
- Supports both MD5 and SHA-256 hashing methods.
- Flexible command-line options for specifying the root directory and hashing method.

## Installation

Clone the repository:
```bash
git clone https://github.com/p11i/duplicate-file-finder.git
cd duplicate-file-finder
```

Build the executable:
```bash
go build -o duplicate-file-finder cmd/duplicate-file-finder/main.go
```
## Usage
Run the tool with default options:
```bash
./duplicate-file-finder
```

Specify a different directory and hashing method:
```bash
./duplicate-file-finder -dir /path/to/directory -hash sha256
```

## Options
`-dir`: Root directory to search for duplicate files. Default is the current directory.
`-hash`: Hashing method to use (md5 or sha256). Default is md5.

## Testing
To run tests, use the following command:
```bash
go test ./...
```

## License
This project is licensed under the GPL-v3.0 License.
