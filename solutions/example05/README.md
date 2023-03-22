# Example 5

Calling the service from `grpcurl`:

```shell
grpcurl --plaintext -H "token: root" -d '{"name": "HWSW"}' 127.0.0.1:8080 example05.Greeter/SayHello
```

To get an unauthenticated error:

```shell
grpcurl --plaintext -d '{"name": "HWSW"}' 127.0.0.1:8080 example05.Greeter/SayHello
```

To get custom error:

```shell
grpcurl --plaintext -H "token: root" -d '{"name": "noone"}' 127.0.0.1:8080 example05.Greeter/SayHello
```
