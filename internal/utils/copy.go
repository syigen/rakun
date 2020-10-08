package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFiles(src, dst string, bufferSize int64) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	srcDir, _ := filepath.Split(src)
	CreateDir(srcDir, false)

	dstDir, _ := filepath.Split(dst)
	CreateDir(dstDir, false)

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
