package defines

import (
	"os"
	"path/filepath"
)

var APP_DATA_DIR = filepath.Join(os.Getenv("LOCALAPPDATA"), "fiz")
