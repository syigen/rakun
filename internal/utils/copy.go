package utils

import (
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func Copy(src, dst string) error {
	fi, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return Dir(src, dst)
	case mode.IsRegular():
		return File(src, dst)
	}
	return nil
}

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

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

func PkgFile(src, dst string) error {
	var err error
	var dstfd *os.File

	srcfd, err := pkger.Open(src)
	if err != nil {
		return err
	}
	defer srcfd.Close()

	info, err := srcfd.Stat()
	if err != nil {
		return err
	}

	dstDir, _ := filepath.Split(dst)
	CreateDir(dstDir, false)

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	return os.Chmod(dst, info.Mode())
}
