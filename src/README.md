# StackHead Project validator

## Development

The schema is expected to be in `schema/project-definition.schema.json`.  
The schema folder is expected to be parallel to the `bin` directory.

### Run Tests

```shell script
go test
```

### Build binary

```shell script
go build -o ../bin/project-validator
```