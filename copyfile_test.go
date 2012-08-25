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
	defaultCp   = &copyfile.Copy{}
	keepLinksCp = &copyfile.Copy{
		KeepLinks: true,
	}
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
	n, err := defaultCp.Single(dst, src)
	if err != nil {
		t.Fatalf("got error in single copy: %s", err)
	}
	if n != 12 {
		t.Fatalf("was expecting length 12 but got %d", n)
	}
}

func TestSingleSymlinkCopy(t *testing.T) {
	dst := testfile("d.txt")
	src := testfile("c.txt")
	defer os.Remove(dst)
	n, err := keepLinksCp.Single(dst, src)
	if err != nil {
		t.Fatalf("got error in single copy: %s", err)
	}
	if n != 12 {
		t.Fatalf("was expecting length 12 but got %d", n)
	}
	dstStat, err := os.Stat(dst)
	if err != nil {
		t.Fatalf("error os.Stat dst %s: %s", dst, err)
	}
	if dstStat.Mode()&os.ModeSymlink == os.ModeSymlink {
		t.Fatalf("destination file is not a symlink as expected")
	}
}
