//go:build arch

package arch_test

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDomainHasNoFrameworkImports(t *testing.T) {
	root := "../modules"
	if _, err := os.Stat(root); os.IsNotExist(err) {
		root = "./modules"
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		if !strings.Contains(path, string(os.PathSeparator)+"domain"+string(os.PathSeparator)) {
			return nil
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			t.Errorf("parse %s: %v", path, err)
			return nil
		}
		for _, imp := range f.Imports {
			p := strings.Trim(imp.Path.Value, "\"")
			if strings.HasPrefix(p, "net/http") || strings.Contains(p, "chi") ||
				strings.Contains(p, "pgx") || strings.HasPrefix(p, "database/sql") ||
				strings.Contains(p, "/infrastructure/") || strings.Contains(p, "/interfaces/") ||
				strings.Contains(p, "/application/") {
				t.Errorf("domain %s imports forbidden package %s", path, p)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
