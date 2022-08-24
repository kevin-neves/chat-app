# Real time Golang chat application using Websockets

This is a fullstack real time chat application using Go templates and connecting via Gorilla Websockets. [Gorilla](https://www.gorillatoolkit.org/) offers some useful tools for Golang, including Websockets Api.

### Folder structure
```
.
└── chat-app/
    ├── cmd/
    │   └── web/
    │       ├── main.go
    │       └── routes.go
    ├── html/
    │   └── html.jet
    ├── internal/
    │   └── handlers/
    │       └── handlers.go
    ├── test/
    │   └── main.go
    ├── go.mod
    ├── go.sum
    └── README.md
```
 
### Packages used

* [Pat](github.com/bmizerany/pat)
* [Jet](https://github.com/CloudyKit/jet)
* [Websockets](github.com/gorilla/websocket)

### Prerequisites

* [Go 1.19](https://go.dev/doc/install)

### Running the project

First, install the dependencies from go.mod:
```
go mod download
```

To run the application:
```
go run cmd/web/*.go
```
