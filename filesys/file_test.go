package filesys_test

import (
	"path"
	"os"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
	. "../filesys"
)

var noExistedFile = "./testdata/tmp.txt"

func TestSelfPath(t *testing.T) {
	path := SelfPath()
	if path == "" {
		t.Error("path cannot be empty")
	}
	t.Logf("SelfPath: %s", path)
}

func TestSelfDir(t *testing.T) {
	dir := SelfDir()
	t.Logf("SelfDir: %s", dir)
}

func TestFileExists(t *testing.T) {
	if !FileExists("./file.go") {
		t.Errorf("./file.go should exists, but it didn't")
	}

	if FileExists(noExistedFile) {
		t.Errorf("Wierd, how could this file exists: %s", noExistedFile)
	}
}

func TestSearchFile(t *testing.T) {
	path, err := SearchFile(filepath.Base(SelfPath()), SelfDir())
	if err != nil {
		t.Error(err)
	}
	t.Log(path)

	path, err = SearchFile(noExistedFile, ".")
	if err == nil {
		t.Errorf("err shouldnot be nil, got path: %s", SelfDir())
	}
}

func TestGrepFile(t *testing.T) {
	_, err := GrepFile("", noExistedFile)
	if err == nil {
		t.Error("expect file-not-existed error, but got nothing")
	}

	path := filepath.Join(".", "testdata", "grepe.test")
	lines, err := GrepFile(`^\s*[^#]+`, path)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(lines, []string{"hello", "world"}) {
		t.Errorf("expect [hello world], but receive %v", lines)
	}
}

func TestSameFile(t *testing.T) {
	file1, file2, file3 := "./testdata/basic.dat", "./testdata/basic.dat", "./testdata/grepe.test"
	if ! SameFile(file1, file2) {
		t.Error("should be true")
	}
	if SameFile(file1, file3) {
		t.Error("should be false")
	}
}

func TestReadFileAll(t *testing.T) {
	filename := "./testdata/basic.dat"
	if _, err := ReadFileAll(filename); err != nil {
		t.Error("read file all error : " + err.Error())
	}
}

func TestWriteFileAll(t *testing.T) {
	filename := "./testdata/basic.dat1"
	if err := WriteFileAll(filename, "123"); err != nil {
		t.Error("write file all error : " + err.Error())
	}
}

func TestCopyFile(t *testing.T) {
	tempFolder, err := ioutil.TempDir("./testdata", "test1")
	defer os.RemoveAll(tempFolder)
	if err != nil {
		t.Fatal(err)
	}
	src := path.Join(tempFolder, "src")
	dest := path.Join(tempFolder, "dest")
	ioutil.WriteFile(src, []byte("content"), 0777)
	ioutil.WriteFile(dest, []byte("destContent"), 0777)
	bytes, err := CopyFile(src, dest)
	if err != nil {
		t.Fatal(err)
	}
	if bytes != 7 {
		t.Fatalf("Should have written %d bytes but wrote %d", 7, bytes)
	}
	actual, err := ioutil.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if string(actual) != "content" {
		t.Fatalf("Dest content was '%s', expected '%s'", string(actual), "content")
	}
}

func TestCreateIfNotExistsDir(t *testing.T) {
	tempFolder, err := ioutil.TempDir("./testdata", "test2")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	folderToCreate := filepath.Join(tempFolder, "tocreate")

	if err := CreateIfNotExists(folderToCreate, true); err != nil {
		t.Fatal(err)
	}
	fileinfo, err := os.Stat(folderToCreate)
	if err != nil {
		t.Fatalf("Should have create a folder, got %v", err)
	}

	if !fileinfo.IsDir() {
		t.Fatalf("Should have been a dir, seems it's not")
	}
}

func TestCreateIfNotExistsFile(t *testing.T) {
	tempFolder, err := ioutil.TempDir("./testdata", "test3")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempFolder)

	fileToCreate := filepath.Join(tempFolder, "file/to/create")

	if err := CreateIfNotExists(fileToCreate, false); err != nil {
		t.Fatal(err)
	}
	fileinfo, err := os.Stat(fileToCreate)
	if err != nil {
		t.Fatalf("Should have create a file, got %v", err)
	}

	if fileinfo.IsDir() {
		t.Fatalf("Should have been a file, seems it's not")
	}
}

