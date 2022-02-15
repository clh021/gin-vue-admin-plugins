package study_plugin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func s3(context *gin.Context) {
	//实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

}

func fibonacci() func(x int) int {
	return func(x int) int {
		return x
	}
}
