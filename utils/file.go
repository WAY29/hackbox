package utils

import (
	"io/ioutil"
	"os"
)

var TempFileNames = []string{}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// 写入文件
func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0755)
}

// 删除文件
func RemoveFile(path string) error {
	return os.Remove(path)
}

// 生成临时文件
func TempFile() (*os.File, error) {
	file, err := ioutil.TempFile("", ".hackbox.file.")
	if err != nil {
		return nil, err
	}

	TempFileNames = append(TempFileNames, file.Name())

	return file, nil
}
