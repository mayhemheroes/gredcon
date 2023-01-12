FROM golang:1.18 as builder

COPY . /gredcon
WORKDIR /gredcon/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang


RUN go-fuzz-build -libfuzzer -o fuzzgredcon.a && \
    clang -fsanitize=fuzzer fuzzgredcon.a -o fuzzgredcon.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /gredcon/mayhem/fuzzgredcon.libfuzzer /