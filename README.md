# Worker Generator CLI 🛠️

A Go CLI tool that generates worker project templates with customizable configurations.

## Features

- Generates ready-to-use Go worker projects
- Customizable project name and Go version
- Supports both CLI and programmatic usage
- Clean template structure with best practices

## Installation

### Method 1: Install directly (requires Go 1.24+)
```bash
go install github.com/LesterCerioli/worker-generator/cmd/worker-cli@latest
```

### Method 2: Build from source
```bash
git clone https://github.com/LesterCerioli/worker-generator.git
cd worker-generator
go build -o worker-cli ./cmd/worker-cli/main.go
sudo mv worker-cli /usr/local/bin/  # Optional: install system-wide
```

## Usege

### Basic Generation
```bash
worker-cli --name my-worker --version 1.24 --output ./my-project
```



![image](https://github.com/user-attachments/assets/75eb80aa-2b5a-41d5-86f1-b33988701f58)


### License

MIT © [LesterCerioli]
