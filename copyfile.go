// Package copyfile provides useful routines for copying files and directories.
package copyfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Configure the Copy process.
type Copy struct {
	KeepLinks bool // don't follow symlinks, instead copy symlink itself
	Force     bool // remove destination and retry if possible when necessary
	Clobber   bool // overwrite existing files
	Recursive bool // recursively copy files in directories
}

// Copy a single file, creating necessary target directories.
func (c *Copy) Single(dst, src string) (int64, error) {
	srcStat, err := os.Stat(src)
	if err != nil {
		return 0, fmt.Errorf("error os.Stat src %s: %s", src, err)
	}
	if srcStat.IsDir() {
		return 0, fmt.Errorf("error src is a directory: %s", src)
	}
	dstParent := filepath.Dir(dst)
	dstParentStat, _ := os.Stat(dstParent)
	if dstParentStat != nil && !dstParentStat.IsDir() {
		return 0, fmt.Errorf("error dst parent is not a directory %s", dstParent)
	} else {
		// TODO: transfer dirmode
		err = os.MkdirAll(dstParent, os.FileMode(0755))
		if err != nil {
			return 0, fmt.Errorf("error creating parent %s: %s", dstParent, err)
		}
	}
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0666)
	if err != nil {
		return 0, fmt.Errorf("error opening src file %s: %s", src, err)
	}
	defer srcFile.Close()
	// TODO: transfer FileMode
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, srcStat.Mode())
	if err != nil {
		return 0, fmt.Errorf("error opening dst file %s: %s", dst, err)
	}
	defer dstFile.Close()
	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return 0, fmt.Errorf("error in copy from %s to %s: %s", src, dst, err)
	}
	return n, nil
}
