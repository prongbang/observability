FROM rust:1.69 as builder

# Check if we are doing cross-compilation, if so we need to add in some more dependencies and run rustup
RUN apt-get update && apt-get install --no-install-recommends -y g++-aarch64-linux-gnu libc6-dev-arm64-cross libprotobuf-dev protobuf-compiler ca-certificates && \
    rustup target add aarch64-unknown-linux-gnu && \
    rustup toolchain install stable-aarch64-unknown-linux-gnu;

WORKDIR /app/

COPY . .

# Compile or crosscompile
RUN cargo build --release --bin user-server

FROM debian:bullseye-slim as release

WORKDIR /app
COPY --from=builder /app/target/release/user-server /app/user-server

EXPOSE 50051
ENTRYPOINT ["/app/user-server"]