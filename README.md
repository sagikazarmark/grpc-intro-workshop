# gRPC intro workshop

Record (HUN): https://www.youtube.com/watch?v=ahb4qqKlaaA

## Prerequisites

1. Git, Make, etc.
2. Make sure you have the latest [Go](https://golang.org/)
3. Install the following tools:
  - [Protobuf](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation)
  - Protobuf plugins
    - [Go](https://protobuf.dev/reference/go/go-generated/)
    - [Go gRPC](https://grpc.io/docs/languages/go/quickstart/)
  - [Buf](https://buf.build/docs/installation/)
  - [grpcurl](https://github.com/fullstorydev/grpcurl#installation)
  - [Task] (https://taskfile.dev/) _(optional)_

Alternatively, install [nix](https://nixos.org) and [direnv](https://direnv.net), then run `direnv allow` once you checked out the repository.

## Usage

1. Checkout this repository
1. Go through the examples
1. Check out the solutions if you need inspiration

## Development

Make sure [nix](https://nixos.org) and [direnv](https://direnv.net) are installed, then run `direnv allow`.

To work on the slides, run `task slides`.
It will open a browser window and automatically refresh the page when you make changes to the slides.
