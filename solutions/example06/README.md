# Example 6

Calling the service from `grpcurl`:

```shell
grpcurl --plaintext -H "token: root" -d '{"name": "HWSW"}' 127.0.0.1:8080 example05.Greeter/SayHello
```
