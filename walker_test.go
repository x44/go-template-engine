package temple

import (
	"os"
	"path/filepath"
	"testing"
)

func createTmp(t *testing.T) (root string, nDirs, nFiles int) {
	root = filepath.Join(os.TempDir(), "temple")

	subs := []string{
		"dir1",
		"dir2",
	}
	fns := []string{
		"file1.txt",
		"file2.txt",
	}

	dirs := []string{
		root,
	}

	files := []string{}

	for _, f := range fns {
		files = append(files, filepath.Join(root, f))
	}
	for _, s := range subs {
		dirs = append(dirs, filepath.Join(root, s))
		for _, f := range fns {
			files = append(files, filepath.Join(root, s, f))
		}
		for _, t := range subs {
			dirs = append(dirs, filepath.Join(root, s, t))
			for _, f := range fns {
				files = append(files, filepath.Join(root, s, t, f))
			}
		}
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			t.Fatalf("failed to create dir %s %v", dir, err)
		}
	}
	for _, file := range files {
		_, err := os.Create(file)
		if err != nil {
			t.Fatalf("failed to create file %s %v", file, err)
		}
		s := "line_with_crlf\r\nline_with_lf\nline_no_eol"
		os.WriteFile(file, []byte(s), os.ModePerm)
	}

	nDirs = len(dirs) - 1 // not counting root dir
	nFiles = len(files)
	return root, nDirs, nFiles
}

func deleteTmp() {
	root := filepath.Join(os.TempDir(), "temple")
	os.RemoveAll(root)
}

func getList(t *testing.T, walker *Walker) []string {
	list, err := walker.List()
	if err != nil {
		deleteTmp()
		t.Fatalf("unexpected error\n%v", err)
	}
	return list
}

func TestRecursiveNoFilter(t *testing.T) {
	root, nDirs, nFiles := createTmp(t)
	var list []string

	walker := NewWalker(root)

	list = getList(t, walker)
	if len(list) != nFiles {
		t.Errorf("expected files: %d got: %d", nFiles, len(list))
	}

	walker.IncludeDirs(true)
	walker.IncludeFiles(true)
	list = getList(t, walker)
	if len(list) != nDirs+nFiles {
		t.Errorf("expected dirs+files: %d got: %d", nDirs+nFiles, len(list))
	}

	walker.IncludeDirs(true)
	walker.IncludeFiles(false)
	list = getList(t, walker)
	if len(list) != nDirs {
		t.Errorf("expected dirs: %d got: %d", nDirs, len(list))
	}

	deleteTmp()
}

func TestNonRecursiveNoFilter(t *testing.T) {
	root, _, _ := createTmp(t)
	var list []string

	walker := NewWalker(root)
	walker.Recursive(false)

	list = getList(t, walker)
	if len(list) != 2 {
		t.Errorf("expected files: %d got: %d", 2, len(list))
	}

	walker.IncludeDirs(true)
	walker.IncludeFiles(true)
	list = getList(t, walker)
	if len(list) != 4 {
		t.Errorf("expected dirs+files: %d got: %d", 4, len(list))
	}

	walker.IncludeDirs(true)
	walker.IncludeFiles(false)
	list = getList(t, walker)
	if len(list) != 2 {
		t.Errorf("expected dirs: %d got: %d", 2, len(list))
	}

	deleteTmp()
}

func TestRecursiveFilter(t *testing.T) {
	root, nDirs, _ := createTmp(t)
	var list []string

	walker := NewWalker(root)

	walker.Filter(func(dir, name string, isFile bool) bool {
		if !isFile {
			return true
		}
		return name == "file1.txt"
	})

	list = getList(t, walker)
	if len(list) != 7 {
		t.Errorf("expected files: %d got: %d", 7, len(list))
	}

	walker.IncludeDirs(true)
	walker.IncludeFiles(true)
	list = getList(t, walker)
	if len(list) != 13 {
		t.Errorf("expected dirs+files: %d got: %d", 13, len(list))
	}

	walker.IncludeDirs(true)
	walker.IncludeFiles(false)
	list = getList(t, walker)
	if len(list) != nDirs {
		t.Errorf("expected dirs: %d got: %d", nDirs, len(list))
	}

	deleteTmp()
}

func TestAbsoluteFilename(t *testing.T) {
	root, _, _ := createTmp(t)
	var list []string

	walker := NewWalker(root)
	walker.Recursive(false)
	walker.Filter(func(dir, name string, isFile bool) bool {
		return name == "file1.txt"
	})

	list = getList(t, walker)
	if len(list) != 1 {
		t.Errorf("expected files: %d got: %d", 1, len(list))
	}

	abs := filepath.Join(root, "file1.txt")
	if list[0] != abs {
		t.Errorf("expected filename:\n%s\ngot:\n%s", abs, list[0])
	}
	deleteTmp()
}
