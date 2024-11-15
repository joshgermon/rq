# rq

[!WARNING]
This project is still in early development and not yet ready for production use. Please star the repository to follow along with updates.

[![Go Report Card](https://goreportcard.com/badge/github.com/joshgermon/rq)](https://goreportcard.com/report/github.com/joshgermon/rq)
[![License](https://img.shields.io/github/license/joshgermon/rq)]

A streamlined CLI tool for testing APIs using OpenAPI/Swagger specifications. Edit and send API requests without leaving your terminal.



A streamlined CLI tool for testing APIs using OpenAPI/Swagger specifications. Edit and send API requests without leaving your terminal.

```bash
# Quick start
rq call users-api

📋 Select endpoint (type to search):
> /users/create

📋 Select method:
> POST

📝 Request body from spec:
{
  "name": "string",
  "email": "user@example.com",
  "organization": {
    "id": "org_123",
    "role": "member"
  }
}
```

## Features

- 🔍 **Fuzzy search** through all your API endpoints
- ✏️ **Edit requests** from the command line or in $EDITOR
- 🎯 **Smart examples** from OpenAPI specs
- 🔐 **Environment management** for different API endpoints and auth tokens **(coming soon)**
- 💾 **Save and reuse** common requests **(coming soon)**

## Installation

```bash

go install github.com/joshgermon/rq
```

## Quick Start

From within the working directory of your api where you have your OpenAPI/Swagger spec, run the following command:

```bash
rq call

```

## Contributing


Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.


## Why rq?

While tools like Postman, Insomnia, and various OpenAPI CLIs exist, `rq` focuses on:

1. **Speed**: Quick endpoint discovery and request editing
2. **Terminal-first**: No context switching to GUI applications
4. **Simplicity**: Minimal interface, maximum productivity

## License

MIT
