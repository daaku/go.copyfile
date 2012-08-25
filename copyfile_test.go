package copyfile_test

import (
	"github.com/daaku/go.copyfile"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	defaults = &copyfile.Copy{}
)

func testfile(name string) string {
	pkg, err := build.Import("github.com/daaku/go.copyfile", "", build.FindOnly)
	if err != nil {
		log.Fatalf("error in finding import path: %s", err)
	}
	return filepath.Join(pkg.Dir, "_testdata", name)
}

func TestSingleNormalCopy(t *testing.T) {
	dst := testfile("b.txt")
	src := testfile("a.txt")
	defer os.Remove(dst)
	n, err := defaults.Single(dst, src)
	if err != nil {
		t.Fatalf("got error in single copy: %s", err)
	}
	if n != 12 {
		t.Fatalf("was expecting length 12 but got %d", n)
	}
}
