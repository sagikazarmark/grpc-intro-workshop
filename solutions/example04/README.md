# Example 4

Calling the service from `grpcurl`:

```shell
grpcurl --plaintext -d '{"name": "HWSW"}' localhost:8080 example04.Greeter/SayHello
```
