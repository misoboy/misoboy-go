# misoboy-go

> [Iris](https://iris-go.com) 웹 프레임워크를 활용한 Go 게시판 웹 애플리케이션입니다.

[![Go](https://img.shields.io/badge/Go-1.x-00ADD8?logo=go&logoColor=white)](https://golang.org)
[![Iris](https://img.shields.io/badge/Iris-Framework-5BC0EB)](https://iris-go.com)

## Overview

Go 언어와 Iris 웹 프레임워크를 학습하기 위한 프로젝트로, MVC 패턴으로 구성된 간단한 게시판(Board) 애플리케이션입니다.

## Project Structure

```
.
├── main.go          # Entry point
├── controllers/     # HTTP handlers
├── models/          # Data models
├── services/        # Business logic
├── repository/      # Data access layer
├── views/           # HTML templates
└── common/          # Shared utilities
```

## Getting Started

```bash
git clone https://github.com/misoboy/misoboy-go.git
cd misoboy-go
go run main.go
```

Application runs on `http://localhost:8080`

## License

MIT
