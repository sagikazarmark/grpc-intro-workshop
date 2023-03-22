# Example 01

Generating Go code with `protoc`:

```shell
mkdir -p generated/go
protoc --go_out=./generated/go/ person.proto
```

Generating Go code with `protoc` (relative source path):

```shell
mkdir -p generated/go
protoc --go_out=./generated/go/ --go_opt=paths=source_relative person.proto
```

Generating Java code with `protoc`:

```shell
mkdir -p generated/java
protoc --java_out=./generated/java/ person.proto
```

Generating code with `buf`:

```shell
buf generate
```
