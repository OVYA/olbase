package compress

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToGZip(t *testing.T) {
	assert := assert.New(t)
	path := "dataTest/file*.txt"

	assert.Nil(FileToGZip(path))

	filesOk := []string{"dataTest/file1.txt.gz", "dataTest/file2.txt.gz"}

	for i := range filesOk {
		if _, err := os.Stat(filesOk[i]); os.IsNotExist(err) {
			t.Errorf("The file '%s' is not generated", filesOk[i])
		} else {
			assert.NoError(os.Remove(filesOk[i]))
		}
	}

	filesKo := []string{"dataTest/file3.notxt.gz"}

	for i := range filesKo {
		if _, err := os.Stat(filesKo[i]); err == nil {
			t.Errorf("The file '%s' is generated but should not", filesKo[i])
			assert.NoError(os.Remove(filesKo[i]))
		}
	}
}
