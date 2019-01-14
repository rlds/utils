package filesys_test

import (
	"testing"
	. "../filesys"
)

var fileTest = []struct {
	filename string
	filetype string
}{
	{"./testdata/re2-search.txt", ""},
	{"./testdata/re2-search.txt.gz", "gz"},
	{"./testdata/re2-exhaustive.txt.bz2", "bz2"},
}

func TestNewLinesReaderFromFile(t *testing.T) {
	for _, v := range fileTest {
		_, err := NewLinesReaderFromFile(v.filename)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestReadLines(t *testing.T) {
	for _, v := range fileTest {
		if v.filetype != "" {
			continue
		}
		r, err := NewLinesReaderFromFile(v.filename)
		if err != nil {
			t.Errorf(err.Error())
		}
		for line := range r.ReadLines() {
			t.Logf("%s\n", line)
		}
	}
}

func TestReadLinesBytes(t *testing.T) {
	for _, v := range fileTest {
		if v.filetype != "" {
			continue
		}
		r, err := NewLinesReaderFromFile(v.filename)
		if err != nil {
			t.Errorf(err.Error())
		}
		for line := range r.ReadLinesBytes() {
			t.Logf("%v\n", line)
		}
	}
}

func TestBreak(t *testing.T) {
	for _, v := range fileTest {
		if v.filetype != "" {
			continue
		}
		r, err := NewLinesReaderFromFile(v.filename)
		if err != nil {
			t.Errorf(err.Error())
		}
		count := 0
		for line := range r.ReadLines() {
			if count == 10 {
				t.Logf("break line %d, %s\n", count, line)
				r.Break()
			}
			count++
		}
	}
}

