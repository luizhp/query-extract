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

# Sql Server volume folder

```shell
/opt/mssql/bin/sqlservr: Error: The system directory [/.system] could not be created. File: LinuxDirectory.cpp:420 [Status: 0xC0000022 Access Denied errno = 0xD(13) Permission denied]
```

```shell
chown 10001:10001 <<mount folder of /var/opt/mssql>>
```
