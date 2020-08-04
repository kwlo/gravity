# Gravity

*3D Physics Engine Written in Go*

## Installation

```
go get -u github.com/go-delve/delve/cmd/dlv
go get -u go.uber.org/zap
```

## Start

```
go run github.com/kwlo/gravity
```

## Run tests

```
CGO_ENABLED=0 go test ./...
```

## Configurations / Env var

```
PORT=8080
```

## Docker

### Docker Build

```
docker build --tag gravity .
```

### Docker Run

```
docker run --rm --detach --publish 8080:8080 --name gravity gravity
```

## Todo backlogs

- Define Interfaces (Shapes)
- Define Sphere
- Define RigidBody
- Simulations
