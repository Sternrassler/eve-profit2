@echo off
REM EVE Profit Calculator 2.0 - Development Server Manager (Windows)
REM Based on Universal Development Guidelines - Clean Code + Automation

setlocal enabledelayedexpansion

REM Configuration
set BACKEND_PORT=9000
set FRONTEND_PORT=3001
set BACKEND_DIR=.\backend
set FRONTEND_DIR=.\frontend

REM Colors (Windows cmd doesn't support colors well, but we'll use echo)
set "RED=[ERROR]"
set "GREEN=[SUCCESS]"
set "YELLOW=[WARNING]"
set "BLUE=[EVE-DEV]"

REM Function to check if port is in use
:check_port
set port=%1
netstat -an | findstr ":%port% " >nul 2>&1
if %errorlevel% equ 0 (
    set port_in_use=1
) else (
    set port_in_use=0
)
goto :eof

REM Function to kill process on port
:kill_port
set port=%1
set service=%2
call :check_port %port%
if !port_in_use! equ 1 (
    echo %YELLOW% %service% running on port %port% - killing process...
    for /f "tokens=5" %%a in ('netstat -ano ^| findstr ":%port% "') do (
        taskkill /PID %%a /F >nul 2>&1
    )
    timeout /t 2 >nul
    echo %GREEN% %service% stopped successfully
) else (
    echo %BLUE% %service% not running on port %port%
)
goto :eof

REM Function to check Go installation
:check_go
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo %RED% Go is not installed! Please install Go 1.21+ for EVE Backend
    exit /b 1
)
for /f "tokens=3" %%i in ('go version') do (
    echo %BLUE% Go version: %%i
    goto :eof
)
goto :eof

REM Function to check Node.js installation
:check_node
where node >nul 2>&1
if %errorlevel% neq 0 (
    echo %RED% Node.js is not installed! Please install Node.js 18+ for EVE Frontend
    exit /b 1
)
for /f %%i in ('node --version') do (
    echo %BLUE% Node.js version: %%i
)
goto :eof

REM Function to start backend
:start_backend
echo %BLUE% Starting EVE Backend Server...

if not exist "%BACKEND_DIR%" (
    echo %RED% Backend directory not found: %BACKEND_DIR%
    exit /b 1
)

if not exist "%BACKEND_DIR%\cmd\server\main.go" (
    echo %RED% Backend main.go not found: %BACKEND_DIR%\cmd\server\main.go
    exit /b 1
)

call :kill_port %BACKEND_PORT% "Backend"

cd %BACKEND_DIR%
echo %BLUE% Executing: go run cmd/server/main.go
start /B go run cmd/server/main.go > ..\backend.log 2>&1
cd ..

echo %BLUE% Waiting for backend to start on port %BACKEND_PORT%...
set attempts=0
:wait_backend
call :check_port %BACKEND_PORT%
if !port_in_use! equ 1 (
    echo %GREEN% Backend started successfully on port %BACKEND_PORT%
    timeout /t 2 >nul
    curl -s http://localhost:%BACKEND_PORT%/api/v1/health >nul 2>&1
    if !errorlevel! equ 0 (
        echo %GREEN% Backend health check passed ✅
    ) else (
        echo %YELLOW% Backend started but health check failed
    )
    goto :eof
)

set /a attempts+=1
if !attempts! lss 30 (
    timeout /t 1 >nul
    goto wait_backend
)

echo %RED% Backend failed to start within 30 seconds
goto :eof

REM Function to start frontend
:start_frontend
echo %BLUE% Starting EVE Frontend Server...

if not exist "%FRONTEND_DIR%" (
    echo %RED% Frontend directory not found: %FRONTEND_DIR%
    exit /b 1
)

if not exist "%FRONTEND_DIR%\package.json" (
    echo %RED% Frontend package.json not found: %FRONTEND_DIR%\package.json
    exit /b 1
)

call :kill_port %FRONTEND_PORT% "Frontend"

if not exist "%FRONTEND_DIR%\node_modules" (
    echo %BLUE% Installing frontend dependencies...
    cd %FRONTEND_DIR%
    call npm install
    cd ..
)

cd %FRONTEND_DIR%
echo %BLUE% Executing: npx vite --port %FRONTEND_PORT%
start /B npx vite --port %FRONTEND_PORT% > ..\frontend.log 2>&1
cd ..

echo %BLUE% Waiting for frontend to start on port %FRONTEND_PORT%...
set attempts=0
:wait_frontend
call :check_port %FRONTEND_PORT%
if !port_in_use! equ 1 (
    echo %GREEN% Frontend started successfully on port %FRONTEND_PORT%
    goto :eof
)

set /a attempts+=1
if !attempts! lss 30 (
    timeout /t 1 >nul
    goto wait_frontend
)

echo %RED% Frontend failed to start within 30 seconds
goto :eof

REM Function to show status
:show_status
echo %BLUE% EVE Profit Calculator 2.0 - Server Status
echo ================================================

call :check_port %BACKEND_PORT%
if !port_in_use! equ 1 (
    echo %GREEN% Backend: Running on port %BACKEND_PORT%
    echo          Health: http://localhost:%BACKEND_PORT%/api/v1/health
) else (
    echo %YELLOW% Backend: Not running
)

call :check_port %FRONTEND_PORT%
if !port_in_use! equ 1 (
    echo %GREEN% Frontend: Running on port %FRONTEND_PORT%
    echo           URL: http://localhost:%FRONTEND_PORT%
) else (
    echo %YELLOW% Frontend: Not running
)

echo ================================================
goto :eof

REM Function to stop all services
:stop_all
echo %BLUE% Stopping EVE Profit Calculator 2.0 services...
call :kill_port %BACKEND_PORT% "Backend"
call :kill_port %FRONTEND_PORT% "Frontend"

if exist backend.log del backend.log
if exist frontend.log del frontend.log

echo %GREEN% All services stopped
goto :eof

REM Function to show help
:show_help
echo EVE Profit Calculator 2.0 - Development Server Manager (Windows)
echo.
echo Usage: %~nx0 [COMMAND]
echo.
echo Commands:
echo   start     Start both backend and frontend servers
echo   stop      Stop all running servers
echo   restart   Restart all servers
echo   status    Show server status
echo   backend   Start only backend server
echo   frontend  Start only frontend server
echo   test      Run backend tests
echo   help      Show this help message
echo.
echo URLs:
echo   Backend:  http://localhost:%BACKEND_PORT%
echo   Frontend: http://localhost:%FRONTEND_PORT%
echo.
echo Examples:
echo   %~nx0 start          # Start both servers
echo   %~nx0 backend        # Start only backend
goto :eof

REM Main script logic
set command=%1
if "%command%"=="" set command=help

if "%command%"=="start" (
    echo %BLUE% Starting EVE Profit Calculator 2.0 Development Environment
    call :check_go
    call :check_node
    call :start_backend
    call :start_frontend
    call :show_status
    echo %GREEN% Development environment ready! 🚀
    echo %BLUE% Backend: http://localhost:%BACKEND_PORT%
    echo %BLUE% Frontend: http://localhost:%FRONTEND_PORT%
) else if "%command%"=="stop" (
    call :stop_all
) else if "%command%"=="restart" (
    call :stop_all
    timeout /t 2 >nul
    call :start_backend
    call :start_frontend
) else if "%command%"=="status" (
    call :show_status
) else if "%command%"=="backend" (
    call :check_go
    call :start_backend
) else if "%command%"=="frontend" (
    call :check_node
    call :start_frontend
) else if "%command%"=="test" (
    echo %BLUE% Running backend tests...
    cd %BACKEND_DIR%
    go test ./...
    cd ..
) else (
    call :show_help
)
