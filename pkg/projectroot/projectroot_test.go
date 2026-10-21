package projectroot

import (
	"os"
	"path/filepath"
	"testing"
)

// requireNoGoModAbove skips the test when a go.mod exists in dir or any of its
// ancestors, as FindProjectRoot would then legitimately find it.
func requireNoGoModAbove(t *testing.T, dir string) {
	t.Helper()

	for {
		_, err := os.Stat(filepath.Join(dir, "go.mod"))
		if err == nil {
			t.Skipf("go.mod found in %s, outside the test directory", dir)
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return
		}

		dir = parent
	}
}

//nolint:paralleltest // t.Chdir cannot be used in parallel tests.
func TestFindProjectRoot(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(t *testing.T, root string) (cwd, want string)
		wantErr string
	}{
		{
			name: "go.mod in a parent directory",
			setup: func(t *testing.T, root string) (string, string) {
				t.Helper()

				writeFile(t, filepath.Join(root, "go.mod"))

				return mkdirAll(t, filepath.Join(root, "a", "b")), root
			},
			wantErr: "",
		},
		{
			name: "go.mod in the working directory",
			setup: func(t *testing.T, root string) (string, string) {
				t.Helper()

				writeFile(t, filepath.Join(root, "go.mod"))

				return root, root
			},
			wantErr: "",
		},
		{
			name: "no go.mod up to the filesystem root",
			setup: func(t *testing.T, root string) (string, string) {
				t.Helper()

				requireNoGoModAbove(t, root)

				return mkdirAll(t, filepath.Join(root, "a")), ""
			},
			wantErr: "project root not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// EvalSymlinks: on macOS t.TempDir() lives under a symlinked /var.
			root, err := filepath.EvalSymlinks(t.TempDir())
			if err != nil {
				t.Fatal(err)
			}

			cwd, want := test.setup(t, root)
			t.Chdir(cwd)

			got, err := FindProjectRoot()
			if test.wantErr != "" {
				if err == nil || err.Error() != test.wantErr {
					t.Fatalf("FindProjectRoot() error = %v, want %q", err, test.wantErr)
				}

				return
			}

			if err != nil {
				t.Fatalf("FindProjectRoot() unexpected error: %v", err)
			}

			if got != want {
				t.Errorf("FindProjectRoot() = %q, want %q", got, want)
			}
		})
	}
}

func mkdirAll(t *testing.T, path string) string {
	t.Helper()

	err := os.MkdirAll(path, 0o750)
	if err != nil {
		t.Fatal(err)
	}

	return path
}

func writeFile(t *testing.T, path string) {
	t.Helper()

	err := os.WriteFile(path, []byte("module example.com/test\n"), 0o600)
	if err != nil {
		t.Fatal(err)
	}
}
