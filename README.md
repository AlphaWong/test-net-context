# Objective
Learning how to use context in GO

# Run
Start server
```bash
go run ./main.go
```

Send get request
```bash
curl 127.0.0.1:8080/get?q=1234
```

# Result
```sh
===Message===
Incoming q = 1234
ctx q = 1234
ctx inside original request q = 1234
===End===
```

# Reference
1. https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39
1. https://golang.org/pkg/context/#example_WithValue
1. https://stackoverflow.com/questions/32164223/how-can-i-create-cloud-context-context-from-appengine-context
