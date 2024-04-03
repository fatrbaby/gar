package shortcut

import (
	"path"
	"runtime"
)

func CurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}
