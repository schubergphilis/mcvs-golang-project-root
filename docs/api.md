# API

## `func FindProjectRoot() (string, error)`

Returns the absolute path of the first directory, starting at the current
working directory and walking up through its parents, that contains a `go.mod`
file. It returns an error if the working directory cannot be determined, or the
error `project root not found` if the filesystem root is reached without
encountering a `go.mod`.

Note that in a repository with nested modules the nearest enclosing module
directory is returned, not the top level of the repository.
