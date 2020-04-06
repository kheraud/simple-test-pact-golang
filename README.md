# Simple test pact golang

Very simple test using pact.io written in go 1.14.

The tested application is an API exposing '/' endpoint for GET requests.

## How to run :

1. Install [pact-ruby-standalone](https://github.com/pact-foundation/pact-ruby-standalone/releases) and have its bin in your `PATH`
    - `cd somewhere`
    - `curl -SLO https://github.com/pact-foundation/pact-ruby-standalone/releases/download/v1.82.2/pact-1.82.2-linux-x86_64.tar.gz`
    - `tar xzf pact-1.82.2-linux-x86_64.tar.gz`
    - `cd pact/bin`
    - `export PATH=$(pwd):$PATH`
2. Launch consumer tests (creates a pacts/myconsumer-myprovider.json file representing what consumer is expecting) : `go test -v -run TestConsumer`
3. Launch producter tests (rely on pacts/myconsumer-myprovider.json file to test on real server) : `go test -v -run TestProvider`

