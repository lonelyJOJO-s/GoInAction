package search

import (
	"encoding/json"
	_ "encoding/json"
	"os"
	_ "os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) { // 解析元数据
	file, err := os.Open(dataFile) // 返回一个指针指向文件和一个错误信息
	if err != nil {
		return nil, err
	}
	defer file.Close() // defer 让后面的函数调用在函数返回时才执行
	// defer有助于提高可读性。 就算主函数意外崩溃，defer后的函数依然会被执行  -- defer的好处

	var feeds []*Feed                          // 注意是多个指针组成的数组
	err = json.NewDecoder(file).Decode(&feeds) // Decode将地址作为输入

	return feeds, err
}
