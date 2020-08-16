# Gravity

*3D Physics Engine Written in Go*

https://gravitydocker.herokuapp.com/

## Installation

```
go build github.com/kwlo/gravity
```

## Start

```
go run github.com/kwlo/gravity
```

## Run tests

```
CGO_ENABLED=0 go test ./... -coverprofile cp.out

/* Generate HTML Coverage Report after running test */
go tool cover -html=cp.out -o coverage.html
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
