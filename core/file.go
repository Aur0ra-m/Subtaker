package core

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetSubNames(namesFilePath string) []string {
	names, err := readLine(namesFilePath)
	if err != nil {
		panic("Dict file read false")
		os.Exit(0)
	}
	return names
}

func readLine(filePath string) ([]string, error) {
	//打开文件
	fi, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	buf := bufio.NewScanner(fi)
	// 循环读取
	var lineArr []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		lineArr = append(lineArr, line)
	}
	return lineArr, nil
}

func readFingers(jsonPath string) (fingers []Finger) {
	bytes := readJson(jsonPath)
	json.Unmarshal(bytes, &fingers)
	return
}

func readJson(jsonPath string) []byte {
	fp, err := os.OpenFile("fingers.json", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	return bytes
}
