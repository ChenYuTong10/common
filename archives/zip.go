package archives

import (
	"archive/zip"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Unzip decompresses zip under the path and destination directory under the same path.
func Unzip(path string) error {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer func() {
		if err = reader.Close(); err != nil {
			panic(err)
		}
	}()

	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return err
	}

	decoder := simplifiedchinese.GBK.NewDecoder()

	for _, file := range reader.File {
		fileName, err := decoder.String(file.Name)
		if err != nil {
			return err
		}
		filePath := filepath.Join(dir, fileName)
		if !strings.HasPrefix(filePath,
			filepath.Clean(dir)+string(os.PathSeparator),
		) {
			return fmt.Errorf("uncorrect file path %s", filePath)
		}
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
		} else {
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return err
			}
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
