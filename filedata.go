package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

/*
   文件读取存储
*/

// FInfo 文件信息
type FInfo struct {
	Path  string
	Finfo os.FileInfo
}

// GetAllFileData 读取文件内容
func GetAllFileData(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	var n int64

	if fi, err2 := f.Stat(); err2 == nil {
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	} else {
		return nil
	}
	buf := bytes.NewBuffer(make([]byte, 0, n+bytes.MinRead))
	defer buf.Reset()
	_, err = buf.ReadFrom(f)
	f.Close()
	if err != nil {
		return nil
	}
	return buf.Bytes()
	//上面为了便于控制 打开文件读取完成后立即 关闭 下面会有延迟
}

//SaveReplaceFile 存储并替换文件内容
func SaveReplaceFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return err
	}
	f.Write(data)
	f.Close()
	return nil
}

// SaveDataToFile SaveDataToFile
func SaveDataToFile(file string, bdat []byte) error {
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	fd.Write(bdat)
	fd.Close()
	return nil
}

//IsFile 判断文件是否存在
func IsFile(filepath string) error {
	_, err := os.Stat(filepath)
	return err
}

//DelFile 删除文件
func DelFile(path string) error {
	return os.Remove(path)
}

//GetFileMd5 计算文件内容的md5
func GetFileMd5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()

	io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil))
}

// ReadDirFileName 遍历获得文件夹中文件名和文件信息
func ReadDirFileName(dirpath string, cfileinfo chan<- FInfo, readend chan<- bool) error {
	err := filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		var fi FInfo
		fi.Path = path
		fi.Finfo = f
		cfileinfo <- fi
		return nil
	})
	readend <- true
	return err
}

// ReadFileLine 按行读取文件
func ReadFileLine(filename string, linstr chan<- string, readend chan<- bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linstr <- scanner.Text()
	}
	readend <- true
	return nil
}

// MkAlldir 建立文件夹 包含子路径
func MkAlldir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// DelDir 删除文件夹
func DelDir(path string) error {
	return os.RemoveAll(path)
}

// CheckDir CheckDir
func CheckDir(dir string) bool {
	finfo, err := os.Stat(dir)
	if err != nil {
		return false
	}
	if finfo.IsDir() {
		return true
	}
	return false
}

func GetProName() string {
	_, name := filepath.Split(os.Args[0])
	return name
}

func GetUserInfo() (name, hpath string) {
	user, err := user.Current()
	if err != nil {
		return
	}
	return user.Name, user.HomeDir
}
