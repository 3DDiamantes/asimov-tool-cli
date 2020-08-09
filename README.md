# asimov-tool-cli
CLI tool for manage builds, tests, deploys and other stuff related to development.

# Build
```
go build -o builds/asimov.exe cmd/main.go
```

# Set the token
## Windows
```
setx ASIMOV_TOOL_CLI_TOKEN <your-github-api-token>
```