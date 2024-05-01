package app

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const URL = "https://raw.githubusercontent.com/zcjswx/code_task/main/misc/graph.xml"

func TestDownloadFileToTmp(t *testing.T) {
	filePath, err := DownloadFileToTmp(URL)
	assert.NoError(t, err)
	info, err := os.Stat(filePath)
	assert.NotEqual(t, err, os.ErrNotExist)
	assert.Greater(t, info.Size(), int64(0))
}
