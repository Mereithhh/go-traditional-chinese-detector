# go 语言繁体字探测器

只探测纯繁体中文，繁体中文和简体中文字一样的不会通过检测。

## 开始使用

### 安装

```bash
go get github.com/mereithhh/go-traditional-chinese-detector
```

### 使用（内置字典）

```go
package main

import (
	"fmt"
	detector "github.com/mereithhh/go-traditional-chinese-detector"
)

func main() {

	str := "你好，世界！彠嗄"
	r := detector.DetectTraditionalChinese(str)
	if r {
		fmt.Println("繁体")
		return
	}
	fmt.Println("简体")

}

```

### 自定义探测字典

```go
package main

import (
	"fmt"
	"os"
	"time"

	detector "github.com/mereithhh/go-traditional-chinese-detector"
)

// 获取自定义字典的函数
func getCustomDict() ([]string, error) {
	fmt.Println("getCustomDict")
	return []string{"你", "好"}, nil
}

func main() {

	// 初始化探测器，使用自定义的字典获取函数，字典更新间隔为60秒
	d, err := detector.NewCustomDetector(getCustomDict, 60)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 检测文本
	isTraditionalChinese := d.Detect("你好")
	fmt.Println(isTraditionalChinese)
	// 延迟推出，查看获取器是否正常工作
	time.Sleep(600 * time.Second)

}

```
