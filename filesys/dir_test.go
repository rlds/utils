package filesys_test

import (
	"testing"
	. "../filesys"
)

func TestCanonicalURLPath(t *testing.T) {
	tests := []struct {
		p  string
		wp string
	}{
		{"/a", "/a"},
		{"", "/"},
		{"a", "/a"},
		{"//a", "/a"},
		{"/a/.", "/a"},
		{"/a/..", "/"},
		{"/a/", "/a/"},
		{"/a//", "/a/"},
	}
	for i, tt := range tests {
		if g := CanonicalURLPath(tt.p); g != tt.wp {
			t.Errorf("#%d: canonical path = %s, want %s", i, g, tt.wp)
		}
	}
}

func TestIsDirWriteable(t *testing.T) {
	if err := IsDirWriteable("./testdata"); err != nil {
		t.Error("dir is not writeable error : " + err.Error())
	}
}

func TestReadDir(t *testing.T) {
	names, err := ReadDir("./testdata")
	if err != nil {
		t.Error("read dir error : " + err.Error())
	}
	if len(names) <= 0 {
		t.Error("read dir error")
	}
}

func TestExist(t *testing.T) {
	if ! Exist("./testdata") {
		t.Error("should be true")
	}
	if Exist("./testdatas") {
		t.Error("should be false")
	}
}

func TestGetDirFileNum(t *testing.T) {
	num := GetDirFileNum("./testdata")
	if num != 10 {
		t.Error("should be 10, now :", num)
	}
}

