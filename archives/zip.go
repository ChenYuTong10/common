package archives

import (
	"archive/zip"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"path/filepath"
)

// Unzip decompresses zip under the path and destination directory under the same path.
//
// TODO: Unzip function is not passing the test because of `test-chinese.zip`.
func Unzip(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	dir := filepath.Dir(absPath)

	reader, err := zip.OpenReader(absPath)
	if err != nil {
		return err
	}
	defer func() {
		if err = reader.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := simplifiedchinese.GBK.NewDecoder()

	for _, file := range reader.File {
		fileName, err := decoder.String(file.Name)
		if err != nil {
			return err
		}
		filePath := filepath.Join(dir, fileName)
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
		} else {
			unzipFile, err := os.Create(filePath)
			if err != nil {
				return err
			}
			zipFile, err := file.Open()
			if err != nil {
				return err
			}
			_, err = io.Copy(unzipFile, zipFile)
			if err != nil {
				return err
			}
			err = zipFile.Close()
			if err != nil {
				return err
			}
			err = unzipFile.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
