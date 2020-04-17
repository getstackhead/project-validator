# Project validator

This tool validates your Mackerel project configuration files.

The schema is provided as [JSON Schema](https://json-schema.org/).

## How to use

```shell script
./bin/project-validator path/to/configuration.yml
```

Find valid and invalid files in examples/ folder.
(Note that examples/invalid/docker_double-expose.yml currently validates... :/)

### Planned packages

* pypi (Python)
* Composer (PHP)
* NPM (NodeJS)

## Development

Build binary:

```shell script
go build -o ../bin/project-validator src/main.go
```