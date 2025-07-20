# EVE-Profit2

## 🚀 EVE Online Profit Analysis Tool

A comprehensive backend service for analyzing profit opportunities in EVE Online using market data and character information.

## 📖 Documentation

All project documentation is located in the [`docs/`](./docs/) directory:

### 📋 Project Documentation
- **[Project Context](./docs/PROJECT_CONTEXT.md)** - Project overview and objectives
- **[Project Status](./docs/PROJECT_STATUS.md)** - Current development status and progress
- **[Character API Specs](./docs/CHARACTER_API_SPECS.md)** - Character API specifications
- **[Session Management](./docs/PROJECT_SESSION_MANAGEMENT.md)** - Development session guidelines

### 🎯 Universal Development Guidelines
- **[Development Guidelines](./docs/UNIVERSAL_DEVELOPMENT_GUIDELINES.md)** - Universal development standards
- **[Clean Code Guidelines](./docs/UNIVERSAL_CLEAN_CODE_GUIDELINES.md)** - Clean Code and SOLID principles
- **[Testing Guidelines](./docs/UNIVERSAL_TESTING_GUIDELINES.md)** - TDD and testing best practices
- **[Session Management Guidelines](./docs/UNIVERSAL_SESSION_MANAGEMENT_GUIDELINES.md)** - Universal session management

## 🏗️ Project Structure

```
eve-profit2/
├── backend/                # Go Backend Services
│   ├── cmd/server/        # Application entry point
│   ├── internal/          # Private application code
│   ├── pkg/               # Public libraries
│   ├── tests/             # Backend tests
│   └── scripts/           # Backend-specific scripts
├── docs/                  # Project documentation
├── scripts/               # Shared project scripts
└── README.md             # This file
```

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Git

### Running the Backend
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Running Tests
```bash
cd backend
go test ./...
```

## 📚 Development

Before starting development, please read:
1. [Universal Development Guidelines](./docs/UNIVERSAL_DEVELOPMENT_GUIDELINES.md)
2. [Universal Testing Guidelines](./docs/UNIVERSAL_TESTING_GUIDELINES.md)
3. [Project Context](./docs/PROJECT_CONTEXT.md)

## 🎮 About EVE Online

This tool analyzes market data from EVE Online, a space-based MMORPG, to help players identify profitable trading and manufacturing opportunities.

## 📄 License

This project is for educational and personal use in accordance with EVE Online's Third-Party Developer License Agreement.
