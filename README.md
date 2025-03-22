# Query Extract


# Test

```shell
go test ./... -coverprofile=coverage.out
```

```shell
go tool cover -html=coverage.out
```

```shell
go test -v ./... -cover
```


# Husky
https://github.com/automation-co/husky


- install
```shell
go install github.com/automation-co/husky@latest
```

- initialize at the source folder where the `.git` directory resides:
```shell
husky init
```

- add command
```shell
husky add pre-commit "
  go build -v ./... 
  go test -v ./...
"
```

