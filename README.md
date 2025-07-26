# EVE-Profit2

## 🚀 EVE Online Profit Analysis Tool

A comprehensive backend service for analyzing profit opportunities in EVE Online using market data and character information.

## 📖 Documentation

**🎯 Primary Rules:** [`.github/copilot-instructions.md`](.github/copilot-instructions.md) - **GitHub Copilot automatically applies these coding standards**

All project documentation is located in the [`docs/`](./docs/) directory:

### 📋 Project Documentation
- **[Project Status](./docs/PROJECT_STATUS.md)** - **Current Phase 7 status and all development progress**
- **[Project Context](./docs/PROJECT_CONTEXT.md)** - Project overview and objectives
- **[Project Session Management](./docs/PROJECT_SESSION_MANAGEMENT.md)** - Development session guidelines

### 🎯 Extended Development Guidelines
- **[Development Guidelines](./docs/UNIVERSAL_DEVELOPMENT_GUIDELINES.md)** - Extended development standards  
- **[Clean Code Guidelines](./docs/UNIVERSAL_CLEAN_CODE_GUIDELINES.md)** - Extended Clean Code principles
- **[Testing Guidelines](./docs/UNIVERSAL_TESTING_GUIDELINES.md)** - Extended TDD patterns

### 📚 Additional Documentation
Check the [`docs/`](./docs/) directory for any additional project documentation.

## ⚠️ Before Starting Development

**🎯 Quick Start - Essential Reading Order:**
1. **[`.github/copilot-instructions.md`](.github/copilot-instructions.md)** - Primary coding rules (GitHub Copilot auto-applies)
2. **[`docs/PROJECT_STATUS.md`](./docs/PROJECT_STATUS.md)** - Current Phase 7 development status
3. **Run tests:** `go test ./... && npm test && npx playwright test` (All 135 tests must pass)

**Required Reading:**
1. **[Project Status](./docs/PROJECT_STATUS.md)** - Current development phase and priorities
2. **[Universal Development Guidelines](./docs/UNIVERSAL_DEVELOPMENT_GUIDELINES.md)** - Required project structure and standards
3. **[Universal Clean Code Guidelines](./docs/UNIVERSAL_CLEAN_CODE_GUIDELINES.md)** - Code quality and SOLID principles
4. **[Universal Testing Guidelines](./docs/UNIVERSAL_TESTING_GUIDELINES.md)** - TDD workflow and testing standards
5. **[Project Context](./docs/PROJECT_CONTEXT.md)** - Project background and objectives

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

### Development Workflow
1. Check **[Project Status](./docs/PROJECT_STATUS.md)** for current phase
2. Follow **[Universal Session Management](./docs/UNIVERSAL_SESSION_MANAGEMENT_GUIDELINES.md)** for session setup
3. Use **[Project Session Management](./docs/PROJECT_SESSION_MANAGEMENT.md)** for project-specific guidelines
4. Implement features following **[Universal Development Guidelines](./docs/UNIVERSAL_DEVELOPMENT_GUIDELINES.md)**

## 🎮 About EVE Online

This tool analyzes market data from EVE Online, a space-based MMORPG, to help players identify profitable trading and manufacturing opportunities.

## 📄 License

This project is for educational and personal use in accordance with EVE Online's Third-Party Developer License Agreement.
