# tern

> The missing ternary shell operator for humans

Just need some ternary support in shell automation.

```shell
# shell mode
tern thruthy_expr yes no
yes

# boolean mode (-b)
tern -b 0 yes no
no

tern -b false yes no
no

tern "" yes no
no
```

## Tests

### Unit

```shell
go test -v -cover ./...
```

### Functional

```shell
bats --tap tests
```

## Build

```shell
go build -o ~/Bin/tern cmd/tern/main.go
```
