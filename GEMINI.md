## Coding Rules
 - in `internal/agent/agent.go`, always keep `RepositoryDirectory` and `Debug` to build or instantiate each new tool


## Git-related operations
 - Never automatically commit any change

## Testing

 - Whenever possible, when creating a function in the `internal` directory, try to write the test
   in a corresponding file. So if you create `internal/foo.go`, try to create `internal/foo_test.go`.

## Documentation

 - When using a new module, update the list of 3rd-party libraries used in the file `README.md`
 - When adding or removing a command-line option, update `README.md`