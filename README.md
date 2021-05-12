# stop-tilt
A hackweek proof-of-concept for stopping the current Tilt run || local resources via the API.

## Usage
### To stop the current Tilt session:
```
go run cmd/main.go
```
(User will be asked for confirmation. Note that this is a POC and doesn't handle multiple Tilt sessions well.)

### To stop one or more serving local resources:
 ```
 go run cmd/main.go resource1 [resource2...]
 ```
Invoke the CLI with one or more resource names to stop. Note that this CLI can only stop local resources with a currently running `serve_cmd`.
