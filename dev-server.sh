#!/bin/bash
# EVE Profit Calculator 2.0 - Development Server Manager
# Based on Universal Development Guidelines - Clean Code + Automation

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BACKEND_PORT=9000
FRONTEND_PORT=3001
BACKEND_DIR="./backend"
FRONTEND_DIR="./frontend"

# Function: Print colored output
print_status() {
    echo -e "${BLUE}[EVE-DEV]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function: Check if port is in use
is_port_in_use() {
    local port=$1
    lsof -i :$port >/dev/null 2>&1
    return $?
}

# Function: Kill process on port
kill_port() {
    local port=$1
    local service=$2
    
    if is_port_in_use $port; then
        print_warning "$service running on port $port - killing process..."
        lsof -ti :$port | xargs kill -9 2>/dev/null || true
        sleep 2
        
        if is_port_in_use $port; then
            print_error "Failed to kill $service on port $port"
            return 1
        else
            print_success "$service stopped successfully"
        fi
    else
        print_status "$service not running on port $port"
    fi
}

# Function: Check if Go is installed
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed! Please install Go 1.21+ for EVE Backend"
        exit 1
    fi
    
    local go_version=$(go version | grep -oE '[0-9]+\.[0-9]+' | head -1)
    print_status "Go version: $go_version"
}

# Function: Check if Node.js is installed
check_node() {
    if ! command -v node &> /dev/null; then
        print_error "Node.js is not installed! Please install Node.js 18+ for EVE Frontend"
        exit 1
    fi
    
    local node_version=$(node --version)
    print_status "Node.js version: $node_version"
}

# Function: Check if npm is installed
check_npm() {
    if ! command -v npm &> /dev/null; then
        print_error "npm is not installed! Please install npm for frontend dependencies"
        exit 1
    fi
}

# Function: Start Backend
start_backend() {
    print_status "Starting EVE Backend Server..."
    
    if [ ! -d "$BACKEND_DIR" ]; then
        print_error "Backend directory not found: $BACKEND_DIR"
        exit 1
    fi
    
    if [ ! -f "$BACKEND_DIR/cmd/server/main.go" ]; then
        print_error "Backend main.go not found: $BACKEND_DIR/cmd/server/main.go"
        exit 1
    fi
    
    # Kill existing backend
    kill_port $BACKEND_PORT "Backend"
    
    # Start backend in background
    cd $BACKEND_DIR
    print_status "Executing: go run cmd/server/main.go"
    nohup go run cmd/server/main.go > ../backend.log 2>&1 &
    local backend_pid=$!
    cd ..
    
    # Wait for backend to start
    print_status "Waiting for backend to start on port $BACKEND_PORT..."
    local attempts=0
    local max_attempts=30
    
    while [ $attempts -lt $max_attempts ]; do
        if is_port_in_use $BACKEND_PORT; then
            print_success "Backend started successfully on port $BACKEND_PORT (PID: $backend_pid)"
            
            # Test health endpoint
            sleep 2
            if curl -s http://localhost:$BACKEND_PORT/api/v1/health >/dev/null 2>&1; then
                print_success "Backend health check passed ✅"
            else
                print_warning "Backend started but health check failed"
            fi
            return 0
        fi
        
        sleep 1
        attempts=$((attempts + 1))
        echo -n "."
    done
    
    echo
    print_error "Backend failed to start within $max_attempts seconds"
    return 1
}

# Function: Start Frontend
start_frontend() {
    print_status "Starting EVE Frontend Server..."
    
    if [ ! -d "$FRONTEND_DIR" ]; then
        print_error "Frontend directory not found: $FRONTEND_DIR"
        exit 1
    fi
    
    if [ ! -f "$FRONTEND_DIR/package.json" ]; then
        print_error "Frontend package.json not found: $FRONTEND_DIR/package.json"
        exit 1
    fi
    
    # Kill existing frontend
    kill_port $FRONTEND_PORT "Frontend"
    
    # Install dependencies if node_modules doesn't exist
    if [ ! -d "$FRONTEND_DIR/node_modules" ]; then
        print_status "Installing frontend dependencies..."
        cd $FRONTEND_DIR
        npm install
        cd ..
    fi
    
    # Start frontend in background
    cd $FRONTEND_DIR
    print_status "Executing: npx vite --port $FRONTEND_PORT"
    nohup npx vite --port $FRONTEND_PORT > ../frontend.log 2>&1 &
    local frontend_pid=$!
    cd ..
    
    # Wait for frontend to start
    print_status "Waiting for frontend to start on port $FRONTEND_PORT..."
    local attempts=0
    local max_attempts=30
    
    while [ $attempts -lt $max_attempts ]; do
        if is_port_in_use $FRONTEND_PORT; then
            print_success "Frontend started successfully on port $FRONTEND_PORT (PID: $frontend_pid)"
            return 0
        fi
        
        sleep 1
        attempts=$((attempts + 1))
        echo -n "."
    done
    
    echo
    print_error "Frontend failed to start within $max_attempts seconds"
    return 1
}

# Function: Show status
show_status() {
    print_status "EVE Profit Calculator 2.0 - Server Status"
    echo "================================================"
    
    if is_port_in_use $BACKEND_PORT; then
        print_success "Backend: Running on port $BACKEND_PORT"
        echo "         Health: http://localhost:$BACKEND_PORT/api/v1/health"
    else
        print_warning "Backend: Not running"
    fi
    
    if is_port_in_use $FRONTEND_PORT; then
        print_success "Frontend: Running on port $FRONTEND_PORT"
        echo "          URL: http://localhost:$FRONTEND_PORT"
    else
        print_warning "Frontend: Not running"
    fi
    
    echo "================================================"
}

# Function: Stop all services
stop_all() {
    print_status "Stopping EVE Profit Calculator 2.0 services..."
    kill_port $BACKEND_PORT "Backend"
    kill_port $FRONTEND_PORT "Frontend"
    
    # Clean up log files
    [ -f backend.log ] && rm backend.log
    [ -f frontend.log ] && rm frontend.log
    
    print_success "All services stopped"
}

# Function: Show logs
show_logs() {
    local service=$1
    
    case $service in
        "backend"|"b")
            if [ -f backend.log ]; then
                print_status "Backend logs (last 50 lines):"
                tail -50 backend.log
            else
                print_warning "Backend log file not found"
            fi
            ;;
        "frontend"|"f")
            if [ -f frontend.log ]; then
                print_status "Frontend logs (last 50 lines):"
                tail -50 frontend.log
            else
                print_warning "Frontend log file not found"
            fi
            ;;
        *)
            print_status "All logs:"
            if [ -f backend.log ]; then
                echo "=== Backend Logs ==="
                tail -25 backend.log
            fi
            if [ -f frontend.log ]; then
                echo "=== Frontend Logs ==="
                tail -25 frontend.log
            fi
            ;;
    esac
}

# Function: Show help
show_help() {
    echo "EVE Profit Calculator 2.0 - Development Server Manager"
    echo ""
    echo "Usage: $0 [COMMAND]"
    echo ""
    echo "Commands:"
    echo "  start     Start both backend and frontend servers"
    echo "  stop      Stop all running servers"
    echo "  restart   Restart all servers"
    echo "  status    Show server status"
    echo "  backend   Start only backend server"
    echo "  frontend  Start only frontend server"
    echo "  logs      Show logs from both servers"
    echo "  logs b    Show backend logs only"
    echo "  logs f    Show frontend logs only"
    echo "  test      Run backend tests"
    echo "  help      Show this help message"
    echo ""
    echo "URLs:"
    echo "  Backend:  http://localhost:$BACKEND_PORT"
    echo "  Frontend: http://localhost:$FRONTEND_PORT"
    echo ""
    echo "Examples:"
    echo "  $0 start          # Start both servers"
    echo "  $0 backend        # Start only backend"
    echo "  $0 logs b         # Show backend logs"
}

# Main script logic
main() {
    local command=${1:-"help"}
    
    case $command in
        "start")
            print_status "Starting EVE Profit Calculator 2.0 Development Environment"
            check_go
            check_node
            check_npm
            start_backend
            start_frontend
            show_status
            print_success "Development environment ready! 🚀"
            print_status "Backend: http://localhost:$BACKEND_PORT"
            print_status "Frontend: http://localhost:$FRONTEND_PORT"
            ;;
        "stop")
            stop_all
            ;;
        "restart")
            stop_all
            sleep 2
            main start
            ;;
        "status")
            show_status
            ;;
        "backend"|"b")
            check_go
            start_backend
            ;;
        "frontend"|"f")
            check_node
            check_npm
            start_frontend
            ;;
        "logs")
            show_logs $2
            ;;
        "test")
            print_status "Running backend tests..."
            cd $BACKEND_DIR
            go test ./...
            cd ..
            ;;
        "help"|"-h"|"--help"|*)
            show_help
            ;;
    esac
}

# Run main function with all arguments
main "$@"
