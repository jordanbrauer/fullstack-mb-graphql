# From REST to GraphQL

_Building GraphQL APIs for Microservice Architectures_

> Note that this project is solely used for demonstration purposes!


## Setup

Follow the steps below to get the repository setup for development!

1. Clone the repository
    ```shell
    git@github.com:jordanbrauer/fullstack-mb-graphql.git \
      && cd fullstack-mb-graphql;
    ```
1. Install modules
    ```shell
    make mod
    ```
1. Generate types
    ```shell
    make types
    ```
1. Start the dev server
    ```shell
    make serve
    ```

## Usage

List all available commands using `make` or

```shell
make help
```

## Tests

Tests are written using built-in Go testing library.

```
make test
```
