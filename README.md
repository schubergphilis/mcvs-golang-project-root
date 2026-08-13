# mcvs-golang-project-root

[![GitHub release](https://img.shields.io/github/v/release/schubergphilis/mcvs-golang-project-root)](https://github.com/schubergphilis/mcvs-golang-project-root/releases)
[![License](https://img.shields.io/github/license/schubergphilis/mcvs-golang-project-root)](LICENSE)

<img src="./assets/logos/mcvs-golang-project-root.png" width="250"></a>

A tiny Go library that resolves the root directory of a project by walking up
from the current working directory until it finds a `go.mod`.

This is useful whenever code needs to locate files relative to the repository
root rather than relative to the process working directory, for example test
fixtures, configuration files or templates. Tests run with the working
directory set to the package under test, so a hard coded relative path like
`../../testdata` breaks as soon as a package is moved.

## Installation

```zsh
go get github.com/schubergphilis/mcvs-golang-project-root
```

## Usage

```go
package main

import (
	"github.com/schubergphilis/mcvs-golang-project-root/pkg/projectroot"
	log "github.com/sirupsen/logrus"
)

func main() {
	projectRoot, err := projectroot.FindProjectRoot()
	if err != nil {
		log.WithError(err).Fatal("unable to find project root")
	}

	log.Infof("project root found at: %s", projectRoot)
}
```

A runnable example is located in [examples](./examples). To run it, issue:

```zsh
go run examples/main.go
```

Resolving a path relative to the root:

```go
root, err := projectroot.FindProjectRoot()
if err != nil {
	return err
}

fixture := filepath.Join(root, "testdata", "input.json")
```

## API

### `func FindProjectRoot() (string, error)`

Returns the absolute path of the first directory, starting at the current
working directory and walking up through its parents, that contains a `go.mod`
file. It returns an error if the working directory cannot be determined, or the
error `project root not found` if the filesystem root is reached without
encountering a `go.mod`.

Note that in a repository with nested modules the nearest enclosing module
directory is returned, not the top level of the repository.

## Development

This project uses [Task](https://taskfile.dev) with the shared build tasks from
[mcvs-golang-action](https://github.com/schubergphilis/mcvs-golang-action). The
version of those remote tasks is pinned in [Taskfile.yml](Taskfile.yml) and kept
in sync with the version used by the GitHub Actions workflow.

```zsh
export TASK_X_REMOTE_TASKFILES=1
task --list-all
```

Some common tasks:

```zsh
task remote:format
task remote:lint
task remote:test
task remote:coverage
```

The same checks run in CI via [.github/workflows/golang.yml](.github/workflows/golang.yml),
which executes the unit, component and integration test suites, linting and the
security scans (`golang-modules`, `grype` and `trivy`).

## License

[MIT](LICENSE)
