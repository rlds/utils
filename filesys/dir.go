package filesys

import (
	"os/exec"
	"runtime"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
)

func CanonicalURLPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root,
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

func IsDirWriteable(dir string) error {
	f := path.Join(dir, ".touch")
	if err := ioutil.WriteFile(f, []byte(""), 0600); err != nil {
		return err
	}
	return os.Remove(f)
}

// ReadDir returns the filenames in the given directory in sorted order.
func ReadDir(dirpath string) ([]string, error) {
	dir, err := os.Open(dirpath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

func Exist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func GetDirFileNum(dirpath string) int {
	num := 0
	filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		num++
		return nil
	})
	return num
}

func MoveDir(src, des string) bool {
	if src == "" || des == "" || src == des {
		return false
	}
	if runtime.GOOS == "linux" {
		mvCmd := exec.Command("mv", src, des)
		mvCmd.Run()
	}
	if runtime.GOOS == "windows" {
		mvCmd := exec.Command("cmd.exe", "/c", "move", src, des)
		mvCmd.Run()
	}
	return true
}

