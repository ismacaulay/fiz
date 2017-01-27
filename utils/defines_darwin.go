package utils

import (
	"os"
	"path/filepath"
)

var APP_DATA_DIR = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "fiz")
