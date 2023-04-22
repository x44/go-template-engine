package temple

import (
	"os"
	"path/filepath"
)

type Walker struct {
	dir          string
	recursive    bool
	includeFiles bool
	includeDirs  bool
	filter       Filter
}

// File/Directory filter function
//
// dir is a directory name relative to the root directory
//
// name is a file name or directory name
type Filter = func(dir, name string, isFile bool) bool

// File/Directory handler function
//
// path is an absolute path
type Handler = func(path string, isFile bool)

// Creates a Walker instance which is recursive, accepts all files and directories and includes files only.
func NewWalker(dir string) *Walker {
	return &Walker{
		dir:          dir,
		recursive:    true,
		includeFiles: true,
		includeDirs:  false,
		filter:       nil,
	}
}

func (w *Walker) Recursive(recursive bool) *Walker {
	w.recursive = recursive
	return w
}

func (w *Walker) IncludeFiles(includeFiles bool) *Walker {
	w.includeFiles = includeFiles
	return w
}

func (w *Walker) IncludeDirs(includeDirs bool) *Walker {
	w.includeDirs = includeDirs
	return w
}

func (w *Walker) Filter(filter Filter) *Walker {
	w.filter = filter
	return w
}

// Returns absolute directories and/or files
func (w *Walker) List() ([]string, error) {
	out := []string{}

	err := w.Walk(func(fn string, isFile bool) {
		out = append(out, fn)
	})

	return out, err
}

func (w *Walker) Walk(handler Handler) error {
	rootDir, err := filepath.Abs(w.dir)
	if err != nil {
		return err
	}
	return w.walk(rootDir, "", handler)
}

func (w *Walker) walk(rootDir, subDir string, handler Handler) error {
	entries, err := os.ReadDir(filepath.Join(rootDir, subDir))
	if err != nil {
		return err
	}
	for _, entry := range entries {

		if !w.accept(subDir, entry.Name(), !entry.IsDir()) {
			continue
		}

		rel := filepath.Join(subDir, entry.Name())

		if entry.IsDir() {
			if w.includeDirs {
				abs, err := filepath.Abs(filepath.Join(rootDir, rel))
				if err != nil {
					return err
				}
				handler(abs, false)
			}
			if w.recursive {
				err := w.walk(rootDir, rel, handler)
				if err != nil {
					return err
				}
			}
		} else {
			if w.includeFiles {
				abs, err := filepath.Abs(filepath.Join(rootDir, rel))
				if err != nil {
					return err
				}
				handler(abs, true)
			}
		}
	}
	return nil
}

func (w *Walker) accept(dir, name string, isFile bool) bool {
	if w.filter != nil {
		return w.filter(dir, name, isFile)
	}
	return true
}
