# lzconvx

一个纯 Go 的字符串转换小库，不依赖 `strconv` 等官方转换函数，手写解析逻辑，支持将字符串转换为整型与浮点型，并提供明确的错误类型。

## 安装

```sh
go get liangzhanbo/lzconvx
```

## 错误类型

- `conv.ErrSyntax`：输入语法错误（空字符串、非数字字符、重复小数点等）。
- `conv.ErrRange`：数值溢出或下溢（超出对应类型范围或指数过大/过小）。

## API 列表与说明

### 整型转换（字符串 -> 有符号整型）

- `conv.StringToInt8(s string) (int8, error)`
- `conv.StringToInt16(s string) (int16, error)`
- `conv.StringToInt32(s string) (int32, error)`
- `conv.StringToInt64(s string) (int64, error)`
- `conv.StringToInt(s string) (int, error)` （根据运行平台自动选择 32/64 位）

行为说明：
- 支持前后空白裁剪、可选正负号。
- 严格检查范围，越界返回 `ErrRange`。
- 只接受十进制数字，出现非数字字符返回 `ErrSyntax`。

### 浮点型转换（字符串 -> 浮点型）

- `conv.StringToFloat64(s string) (float64, error)`
- `conv.StringToFloat32(s string) (float32, error)`

行为说明：
- 支持前后空白裁剪、可选正负号。
- 支持小数点与科学计数法（`e`/`E`，可带正负号）。
- 重复小数点、缺少数字、指数格式错误返回 `ErrSyntax`。
- 指数过大或过小导致溢出/下溢返回 `ErrRange`。

## 使用示例

```go
package main

import (
	"fmt"

	"liangzhanbo/lzconvx/conv"
)

func main() {
	// 整型
	if v, err := conv.StringToInt16(" 32767 "); err == nil {
		fmt.Println("int16:", v)
	}

	// 浮点
	if f, err := conv.StringToFloat64("1.5e2"); err == nil {
		fmt.Println("float64:", f)
	}
}
```

## 行为与实现特点

- 所有解析逻辑手写实现，不使用标准库转换函数。
- 解析过程按位宽检查上/下界，避免溢出。
- 对无效字符、空白、指数格式等都有明确错误反馈。
