{
  description = "gRPC intro workshop";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    skm.url = "github:sagikazarmark/nur-packages";
    skm.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils, skm, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;

          overlays = [
            (final: prev: {
              quarto = skm.packages.${system}.quarto;
              decktape = skm.packages.${system}.decktape;
            })
          ];
        };
      in
      rec {
        devShells = {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake
              go-task

              quarto
              decktape

              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
              buf

              go

              grpcurl
            ];

            shellHook = ''
              echo "quarto $(quarto --version)"
              protoc --version
              echo "buf $(buf --version)"
            '';
          };

          solutions = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              go-task

              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
              buf

              go
            ];

            shellHook = ''
              protoc --version
              echo "buf $(buf --version)"
            '';
          };

          slides = pkgs.mkShell {
            buildInputs = with pkgs; [
              quarto
            ];

            shellHook = ''
              echo "quarto $(quarto --version)"
            '';
          };
        };
      }
    );
}
