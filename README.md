# asimov-tool-cli
CLI tool for manage builds, tests, deploys and other stuff related to development.

# Commands
- `new-feature`,`nf`: Create a new feature branch and upload it to GitHub. E.g.:
    ```
    asimov nf testBranch
    ```

# Build
## Windows
```
go build -o builds/asimov.exe cmd/main.go
```

# Set the token
## Windows
```
setx ASIMOV_TOOL_CLI_TOKEN <your-github-api-token>
```