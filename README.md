franz-go-examples
=================

This repository contains simple examples of how to use the [franz-go](https://github.com/twmb/franz-go) library. The repository accompanies the [blog post](https://aran.dev/posts/getting-started-with-golang-and-kafka). 

## Running the examples

You'll first have to run the `docker-compose` file to start the Kafka compatible broker [Redpanda](https://vectorized.io/redpanda). 

```bash
❯ docker-compose up
```

Then you can run the examples using the following commands:

To run the consumer:

```bash
❯ gor ./consumer/main.go
```

To run the producer: 

```bash
❯ gor ./producer/main.go
```
