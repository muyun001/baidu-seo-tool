package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 打开或者创建文件
func OpenOrCreateFile(fileName string) (*os.File, error) {
	var osFile *os.File
	if isExist, err := PathExists(fileName); err != nil {
		return osFile, err
	} else {
		if isExist {
			osFile, err = os.OpenFile(fileName, os.O_RDWR, 0666)
			if err != nil {
				return osFile, err
			}
		} else {
			osFile, err = os.Create(fileName)
			if err != nil {
				return osFile, err
			}
		}
	}

	return osFile, nil
}

// 判断文件或者文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 将数据写入文件
func WriteResultTxt(file *os.File, content string) {
	bufWriter := bufio.NewWriter(file)
	_, err := bufWriter.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bufWriter.Flush()
	if err != nil {
		fmt.Println(err)
	}
}

// 将数据写入文件
func WriteResultMap(file *os.File, dataMap map[string]int) {
	bufWriter := bufio.NewWriter(file)
	for k, v := range dataMap {
		content := fmt.Sprintf("%s, %d\n", k, v)
		_, err := bufWriter.WriteString(content)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = bufWriter.Flush()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

// 读取所有行
func ReadLines(fileName string) []string {
	var allLines = []string{}
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				continue
			}
		}
		allLines = append(allLines, line)
	}

	return allLines
}

// 追加数据
// 每次openFile不会移除之前的数据
func WriteTxt(file *os.File, content string) {
	n, _ := file.Seek(0, 2)
	_, err := file.WriteAt([]byte(content), n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
