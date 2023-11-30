# Golang OS ARCH Detection

## Usage
### To Setup:
Initialise module:
```
go mod init example.com/golang-os-arch-detection
```

Verify and download dependencies:
- Ensure go.mod and go.sum are up-to-date
```
go mod verify
```

### To Install:

Install non-default packages (e.g. `golang.org/x/sys/unix`)
```
go get golang.org/x/sys/unix
```

### To Run:
```
go run main.go
```

### To Build (binary/executable):

For example, for Linux ARM64/AARCH64:
```
GOOS=linux GOARCH=arm64 go build
```

## Development:
VSCode DevContainer (uses Ubuntu with go and make features).

Requirements:
- VSCode Extension: ms-vscode-remote.remote-containers
- Docker Desktop

To start:
1. Make sure Docker Desktop is running!
2. Go to Command Palette (macOS: Cmd + Shift + P)
3. Search and Select: `Dev Containers: Rebuild And Reopen In Container`
4. It should automatically start building container.
5. If successful, it should not show any popups and bottom left corner of vscode should show green box with: `Dev Container: Ubuntu @ desktop-linux` or something similar.
6. Go to Command Palette (macOS: Cmd + Shift + P)
7. Search and Select: `Terminal: Create New Terminal`

## Local Environment
### To Check:

To see what the current OS and Architecture is:
```
go env GOOS GOARCH
```

Sample output (for macOS with Apple Silicon M1/M2):
```
darwin
arm64
```