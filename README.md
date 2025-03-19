# Pear CLI

Pear CLI is a command-line tool designed to initialize and build backends in Go. Currently, only the `pear-cli init` command is functional, but additional features are planned for the future.

## Installation

To install Pear CLI, use the `go install` command:

```sh
go install github.com/ZiplEix/pear-cli@latest
```

You can also clone this repository and build the project manually using Go.

```sh
git clone https://github.com/ZiplEix/pear-cli.git
cd pear-cli
go build
```

## Usage

### Initializing a Project

To initialize a new backend project in Go, use the following command:

```sh
pear-cli init
```

This command will create the basic structure of the project.

#### Flags for `init` Command

- `-n, --name` : Project name (will be appended to the command `go mod init`). **Required**.
- `-p, --path` : Project path. Default is `./`.
- `-f, --force` : Force the creation of the project even if the directory already exists.
- `--air` : Use air to daemonize the server.
- `--docker` : Use Docker to containerize the server.
- `--swagger` : Use Swagger for documentation.
- `framework` : Framework to use. Default is `fiber`.
- `full` : Use all features with their default values.

#### Usable frameworks

- `fiber` : Fiber framework.

#### Generated projet structure (full features)

```
.
├── Dockerfile
├── docs/
├── go.mod
├── go.sum
├── main.go
└── routes
    ├── main.go
    └── version.go
```

Example:

```sh
pear-cli init --name myproject --path ./myproject --docker --swagger
```

### Adding Routes

A future version of Pear CLI will include the command `pear-cli add route <route name>` to add new routes to your project.

## Contributing

Contributions are welcome! Feel free to submit issues and pull requests.
