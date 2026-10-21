# Usage

This is useful whenever code needs to locate files relative to the repository
root rather than relative to the process working directory, for example test
fixtures, configuration files or templates. Tests run with the working
directory set to the package under test, so a hard coded relative path like
`../../testdata` breaks as soon as a package is moved.

Resolving a path relative to the root:

```go
root, err := projectroot.FindProjectRoot()
if err != nil {
	return err
}

fixture := filepath.Join(root, "testdata", "input.json")
```
