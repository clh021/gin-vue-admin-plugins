package study_plugin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/tour/wc"
)

func s1(context *gin.Context) {
	//实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	keys := strings.Fields(s)
	fmt.Printf("keys: (%T) %q", keys, keys)
	length := len(keys)
	var res = make(map[string]int)
	fmt.Println("")
	fmt.Printf("keys: (%T) %q", length, length)
	fmt.Println("")
	fmt.Println("")
	for _, v := range keys {
		res[v] = 1
		//res = append(res[v], 1)
		fmt.Printf("%q \n", v)
	}
	fmt.Println("")
	return res
}
