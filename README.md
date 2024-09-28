
# GoAPI

![GoAPI Banner](./assets/go-logo-white.svg)

![License](https://img.shields.io/github/license/13city/goapi)
![Go Version](https://img.shields.io/github/go-mod/go-version/13city/goapi)
![Build Status](https://img.shields.io/github/actions/workflow/status/13city/goapi/go.yml?branch=main)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Business Use Cases](#business-use-cases)
- [Installation](#installation)
- [Usage](#usage)
  - [1. Simple Port Scanner](#1-simple-port-scanner)
  - [2. HTTP Proxy Server](#2-http-proxy-server)
  - [3. Encryption/Decryption Utility](#3-encryptiondecryption-utility)
  - [4. Log Analysis Tool](#4-log-analysis-tool)
  - [5. Secure File Transfer Tool](#5-secure-file-transfer-tool)
  - [6. Password Strength Analyzer](#6-password-strength-analyzer)
- [Project Structure](#project-structure)
- [Adding and Styling Images](#adding-and-styling-images)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

---

## Introduction

**goapi** is a Go-based API platform designed to provide a suite of essential cybersecurity tools. With scalability and performance in mind, **goapi** integrates multiple utilities into a single cohesive application, streamlining cybersecurity operations for businesses of all sizes.

![GoAPI Architecture Diagram](./assets/goapi-architecture-diagram.svg)

---

## Features

- **Simple Port Scanner**: Scans open ports on target hosts to identify vulnerabilities.
- **HTTP Proxy Server**: Routes HTTP requests through the proxy for secure and anonymous browsing.
- **Encryption/Decryption Utility**: Robust encryption and decryption functionalities for protecting sensitive data.
- **Log Analysis Tool**: Analyzes server logs to detect suspicious activity.
- **Secure File Transfer Tool**: Securely uploads and downloads files with encryption.
- **Password Strength Analyzer**: Evaluates password strength to meet security standards.

---

## Business Use Cases

- **Network Security Assessment**: Regularly assess network vulnerabilities.
- **Secure Communication**: Ensure secure data transmission through the proxy server.
- **Data Protection**: Safeguard sensitive data in transit and at rest with encryption.
- **Threat Detection**: Analyze logs to detect potential security breaches in real-time.
- **File Management**: Securely handle file transfers with encryption.
- **Password Enforcement**: Enforce strong password policies to reduce unauthorized access risks.

---

## Installation

### Prerequisites

- **Go** (Install from [Go's official website](https://golang.org/dl/))
- **Git** (To clone the repository)

### Steps

```bash
# Clone the repository
git clone https://github.com/13city/goapi.git
cd goapi

# Install dependencies
go mod tidy

# Build the application
go build -o app main.go

# Run the application
./app
```

The server will start on `http://localhost:8080`.

---

## Usage

Once the application is running, you can interact with the various endpoints using `curl` or any API testing tool.

---

### 1. Simple Port Scanner

```bash
# Scan open ports on a target host
curl "http://localhost:8080/portscanner?target=scanme.nmap.org&startPort=20&endPort=80"
```

---

### 2. HTTP Proxy Server

```bash
# Send a request through the proxy
curl "http://localhost:8080/proxy?url=http://example.com"
```

---

### 3. Encryption/Decryption Utility

```bash
# Encrypt text
curl "http://localhost:8080/encrypt?text=HelloWorld&key=your-32-byte-long-encryption-key!!"

# Decrypt text
curl "http://localhost:8080/decrypt?text=encryptedText&key=your-32-byte-long-encryption-key!!"
```

---

### 4. Log Analysis Tool

```bash
# Analyze a log file
curl -F "logfile=@/path/to/logfile.log" http://localhost:8080/loganalysis
```

---

### 5. Secure File Transfer Tool

```bash
# Upload a file
curl -F "file=@/path/to/yourfile.txt" http://localhost:8080/upload

# Download a file
curl "http://localhost:8080/download?filename=yourfile.txt" -O
```

---

### 6. Password Strength Analyzer

```bash
# Analyze password strength
curl "http://localhost:8080/passwordstrength?password=YourP@ssw0rd!"
```

---

## Project Structure

Here is the detailed project structure for **goapi**:

```plaintext
goapi/
├── go.mod                   # Module definition
├── go.sum                   # Dependencies
├── main.go                  # Main application
├── portscanner/             # Port Scanner module
│   └── portscanner.go
├── proxy/                   # Proxy Server module
│   └── proxy.go
├── encryption/              # Encryption/Decryption module
│   └── encryption.go
├── loganalysis/             # Log Analysis module
│   └── loganalysis.go
├── filetransfer/            # Secure File Transfer module
│   └── filetransfer.go
├── passwordstrength/        # Password Strength Analyzer module
│   └── passwordstrength.go
├── images/                  # Images for the repository
│   ├── banner.png
│   └── architecture.png
├── uploads/                 # Uploaded files storage
├── commands.txt             # Script commands
└── README.md                # Documentation
```

---

## Contributing

Contributions are welcome! Follow these steps to contribute:

```bash
# Fork the repository
git clone https://github.com/13city/goapi.git

# Create a new branch
git checkout -b feature/YourFeatureName

# Commit your changes
git add .
git commit -m "Add feature: YourFeatureName"

# Push to your fork and create a pull request
git push origin feature/YourFeatureName
```

---

## License

This project is licensed under the MIT License.

---

## Contact

For questions or suggestions, feel free to reach out:

- GitHub: [13city](https://github.com/13city)

