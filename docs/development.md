# Development

This project uses [Task](https://taskfile.dev) with the shared build tasks from
[mcvs-golang-action](https://github.com/schubergphilis/mcvs-golang-action). The
version of those remote tasks is pinned in [Taskfile.yml](../Taskfile.yml) and kept
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

The same checks run in CI via [.github/workflows/golang.yml](../.github/workflows/golang.yml),
which executes the unit, component and integration test suites, linting and the
security scans (`golang-modules`, `grype` and `trivy`).
