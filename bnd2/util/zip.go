package util

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Unzip(zipFilePath, destination string) error {
	r, err := zip.OpenReader(zipFilePath)

	if nil != err {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		err = cloneZipItem(f, destination)
		if nil != err {
			return err
		}
	}

	return nil
}

func cloneZipItem(f *zip.File, dest string) error {
	// create full directory path
	fileName := f.Name

	if !utf8.ValidString(fileName) {
		data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(fileName)), simplifiedchinese.GB18030.NewDecoder()))
		if nil == err {
			fileName = string(data)
		}
	}

	path := filepath.Join(dest, fileName)

	err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
	if nil != err {
		return err
	}

	if f.FileInfo().IsDir() {
		err = os.Mkdir(path, os.ModeDir|os.ModePerm)
		if nil != err {
			return err
		}

		return nil
	}

	// clone if item is a file

	rc, err := f.Open()
	if nil != err {
		return err
	}

	defer rc.Close()

	// use os.Create() since Zip don't store file permissions
	fileCopy, err := os.Create(path)
	if nil != err {
		return err
	}

	defer fileCopy.Close()

	_, err = io.Copy(fileCopy, rc)
	if nil != err {
		return err
	}

	return nil
}
