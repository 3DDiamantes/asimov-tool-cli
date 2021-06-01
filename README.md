# asimov-tool-cli
CLI tool for manage builds, tests, deploys and other stuff related to development.

# Commands
| Command     | Description                |
|-------------|----------------------------|
| `new-feature` | Create a feature's branch |
| `build`       | Create a build |

## Create new feature
**Command:** `new-feature`,`nf`  
**Description:** Create a new feature branch and upload it to GitHub.  
**Usage:**
```
asimov nf exampleBranch
```
## Build project
**Command:** `build`  
**Description:** Build the project.   
Name can be set manually with `--name` flag, otherwise the name will be _project_branch-vM.m.p_.  
Target can be set manually with `--target` flag, otherwise the target will be the current machine. Supported targets: `arm`, `linux`, `mac`, `win`.

**Usage:** 
```
asimov build --target arm --name something
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