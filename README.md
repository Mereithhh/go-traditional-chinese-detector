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
	detector "github.com/mereithhh/go-traditional-chinese-detector"
)

func main() {


}
```
