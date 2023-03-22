# Example 2

Listing available methods with `grpcurl`:

```shell
grpcurl --plaintext 127.0.0.1:8080 list
```

Calling the service from `grpcurl`:

```shell
grpcurl --plaintext -d '{"name": "HWSW"}' 127.0.0.1:8080 example02.Greeter/SayHello
```
