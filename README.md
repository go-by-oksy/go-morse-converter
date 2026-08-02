# Go Morse Converter

A small web application that converts Russian text to Morse code and Morse code back to text.

The application accepts a text file through a web form, detects the input format, performs the conversion, displays the result in the browser, and saves it as a generated file.

## Features

- Upload text files through a web interface
- Convert Russian text to Morse code
- Convert Morse code back to text
- Automatically detect the conversion direction
- Support digits and common punctuation
- Reject empty input
- Save converted results with cross-platform filenames
- Unit tests for the conversion service
- Automated formatting, testing and build checks with GitHub Actions

## Tech Stack

- Go 1.24
- `net/http`
- Multipart file uploads
- HTML
- Unit testing with the Go standard library
- GitHub Actions

The project uses only the Go standard library and does not require third-party dependencies.

## Project Structure

```text
.
├── .github/workflows   # CI configuration
├── cmd                 # Application entry point
├── internal
│   ├── handlers        # HTTP handlers and file processing
│   ├── server          # HTTP server configuration
│   └── service         # Conversion direction detection
├── pkg
│   └── morse           # Morse encoding and decoding
├── index.html          # Upload form
├── go.mod
└── README.md
```

## Running Locally

### Requirements

- Go 1.24 or newer
- Git

Clone the repository:

```bash
git clone https://github.com/go-by-oksy/go-morse-converter.git
cd go-morse-converter
```

Run the application from the repository root:

```bash
go run ./cmd
```

Open the application in a browser:

```text
http://localhost:8080
```

Choose a text file and submit the form.

The converted content will be:

- displayed in the browser;
- saved in the repository directory with a name similar to:

```text
converted-20260802-120000.000000000.txt
```

## Usage Examples

Text input:

```text
Привет
```

Result:

```text
.--. .-. .. .-- . -
```

Morse input:

```text
.--. .-. .. .-- . -
```

Result:

```text
ПРИВЕТ
```

Only dots, hyphens and whitespace are treated as Morse input. Other content is processed as regular text.

## Testing

Run all tests:

```bash
go test ./...
```

Check that the application builds:

```bash
go build ./...
```

The GitHub Actions workflow also checks formatting, runs the tests and verifies the build on every push and pull request.

## Project Background

This project was completed as part of the Yandex Practicum Go development course and is based on the provided project template.

The HTTP handlers, conversion workflow, error handling, unit tests and subsequent portfolio improvements were completed by [Oksana](https://github.com/go-by-oksy).