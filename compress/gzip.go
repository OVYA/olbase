package compress

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileToGZip compresses the file(s) within the path
// where path can be a glob expression.
func FileToGZip(path string) error {
	var (
		err       error
		bs        []byte
		filepaths []string
	)

	filepaths, err = filepath.Glob(path)
	if err != nil {
		return err
	}

	if len(filepaths) == 0 {
		return fmt.Errorf("The Glob expression '%s' does not return any file...", path)
	}

	for _, path := range filepaths {
		gf, _err := os.Create(path + ".gz")
		if _err != nil {
			return _err
		}

		defer func() {
			if _err := gf.Close(); _err != nil {
				panic(_err)
			}
		}()

		gz := gzip.NewWriter(gf)
		defer func() {
			if _err := gz.Close(); _err != nil {
				panic(_err)
			}
		}()

		bs, err = ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = gz.Write(bs)
	}

	return err
}
