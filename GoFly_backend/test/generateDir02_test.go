/*
@Time : 2023/3/29 17:18
@Author : zic
@File : generateDir02_test
@Software: GoLand
@blog : https://www.cnblogs.com/zichliang
*/
package test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var rootDir string
var separator string
var rootNode Node

const jsonFileName = "dir.json"

// 通过结构体来接收 JSON中的值
type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

func loadJson02() {
	separator = string(filepath.Separator)
	workDir, _ := os.Getwd()
	rootDir = workDir[:strings.LastIndex(workDir, separator)]
	jsonFileBytes, _ := os.ReadFile(workDir + separator + jsonFileName)
	err := json.Unmarshal(jsonFileBytes, &rootNode)
	if err != nil {
		panic(err)
	}
	fmt.Println(rootNode)
}

func parseNode(iNode Node, parentDir string) {
	if iNode.Text != "" {
		createDir02(iNode, parentDir)
	}

	if parentDir != "" {
		parentDir += separator
	}

	if iNode.Text != "" {
		parentDir += iNode.Text
	}

	for _, v := range iNode.Children {
		parseNode(v, parentDir)
	}
}

func createDir02(iNode Node, parentDir string) {
	dirPath := rootDir + separator
	if parentDir != "" {
		dirPath += parentDir + separator
	}
	dirPath += iNode.Text
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestGenerate02(T *testing.T) {
	loadJson02()
	parseNode(rootNode, "")
}
