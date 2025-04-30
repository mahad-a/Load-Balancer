# Golang Load Balancer

This project implements a lightweight HTTP load balancer in Go using round-robin scheduling. It forwards incoming requests to a pool of backend servers and ensures basic fault tolerance by skipping downed servers. A separate client tool is included to simulate traffic by sending concurrent requests to the load balancer.

## Features

- Reverse proxy with `httputil.NewSingleHostReverseProxy`
- Round-robin request distribution
- Thread-safe backend health checking
- Real backend simulation (math tasks per server)
- Logging request handling time
- Client load testing with goroutines and channels

## Project Structure

```
load_balancer/
├── main.go              # Load balancer logic and server pool
├── systems/
│   └── server.go        # Backend servers (3 math-based handlers)
├── client/
│   └── simulator.go          # Client to send requests at intervals
├── run_backends.ps1      # PowerShell script to launch backends easily
```

## Backend Behavior

Each backend server performs a different mathematical operation:
- **Port 8081**: Solves quadratic equations
- **Port 8082**: Computes Pythagorean theorem
- **Port 8083**: Calculates compound interest

## How to Run

### 1. Start Backend Servers

You have two options:

#### Option 1: Manually (One Terminal Per Server)

```bash
PORT=8081 go run systems/server.go
PORT=8082 go run systems/server.go
PORT=8083 go run systems/server.go
```

> If you're on Windows PowerShell, use:
> ```powershell
> $env:PORT=8081; go run systems/server.go
> ```

#### Option 2: Using the PowerShell Script

Run all backend servers automatically with a single command:

```powershell
./run_backends.ps1
```

Make sure PowerShell's execution policy allows running scripts. If needed, you can temporarily allow it with:

```powershell
Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
```

> This avoids permanently changing your system settings.

---

### 2. Start the Load Balancer

In a new terminal:

```bash
go run main.go
```

### 3. Start the Client Simulator

In another terminal:

```bash
cd client
go run simulator.go
```

---

## Sample Output

```bash
2025/04/29 Request handled by http://localhost:8081, duration: 1.23ms
2025/04/29 Calculating quadratic formula for a=2, b=4, c=1
...
```
