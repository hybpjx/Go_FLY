/*
@Time : 2023/3/29 14:25
@Author : zic
@File : generateDir01_test
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

// 分隔符
var stSeparator string

// 当前目录
var stWorkDir string

// 根目录
var stRootDir string

// Json内容
var iJsonData map[string]any

// json文件名
const stJsonFileName = "dir.json"

func loadJson() {
	//分隔符
	stSeparator = string(filepath.Separator)

	//当前目录
	stWorkDir, _ = os.Getwd()

	/*获取根目录
	stringsDirList := strings.Split(stWorkDir, stSeparator)
	stRootDir := strings.Join(stringsDirList[:len(stringsDirList)-1], stSeparator)
	*/

	//获取根目录  strings.LastIndex(stWorkDir, stSeparator) 指最后一次出现 /的时候
	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]
	stJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)
	err := json.Unmarshal(stJsonBytes, &iJsonData)
	if err != nil {
		panic(err.Error())
	}

}

func parseMap(mapData map[string]any, stParentDir string) {
	for _, v := range mapData {
		switch v.(type) {
		case string:
			{
				path := v.(string)
				if path == "" {
					continue
				}
				if stParentDir != "" {
					path = stParentDir + stSeparator + path
					stParentDir = path
				} else {
					stParentDir = path
				}
				createDir(path)
			}
		case []any:
			{
				parseArray(v.([]any), stParentDir)
			}
		}
	}
}

func parseArray(listJsonData []any, stParentDir string) {
	for _, v := range listJsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, stParentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	fmt.Println(path)
	err := os.MkdirAll(stRootDir+stSeparator+path, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
}

func TestGenerateDir01(t *testing.T) {
	loadJson()
	parseMap(iJsonData, "")
}
