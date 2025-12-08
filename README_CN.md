# 🚀 LZCONVX — 一个比官方 `strconv` 更高效、更快速的整数与浮点解析库

[English Version](./README.md) | [中文文档](./README_CN.md)

LZCONVX 是一个经过深度优化、以 **性能与可靠性为核心** 的数字解析库。  
在保持与 Go 官方 `strconv` 行为完全一致的前提下，通过手写 Fast-Path、零分配（Zero-Alloc）、  
更优的分支布局，实现了 **显著快于官方标准库的解析速度**。

适用于：

- 高并发后端
- RPC 框架
- 监控/埋点系统
- 日志平台
- 网络协议解析
- 游戏服务器
- 性能敏感的底层组件


## 📊 性能基准（Apple M1 Max · Go 1.25）

以下基准测试在 **Apple M1 Max** 上运行：

<img width="1212" height="580" alt="benchmark" src="https://github.com/user-attachments/assets/6a536c0a-025d-4283-81f7-3a194afbc6f1" />

以下基准测试数据来源于实际在 **Apple M1 Max** 上运行的结果（见上图），
对比了 LZCONVX 与 Go 官方 `strconv` 的整数解析性能：

| 基准项 | LZCONVX | strconv（官方） | 性能差异 |
|-------|---------|----------------|----------|
| Int8  | **7.117 ns/op** | 8.455 ns/op | ⚡ 提升约 19.8% |
| Int16 | **8.628 ns/op** | 10.37 ns/op | ⚡ 提升约 17% |
| Int32 | **7.933 ns/op** | 14.10 ns/op | ⚡ 提升约 43.7% |
| Int64 | **12.15 ns/op** | 21.33 ns/op | ⚡ 提升约 43.0% |
| Atoi  | **7.106 ns/op** | 6.917 ns/op | ⚠ 约慢 0.2ns（解释见下） |

### 性能总结

👉 **LZCONVX 在 Int8/Int16/Int32/Int64 多项测试中相比 `strconv.ParseInt` 可获得约 17% ~ 43% 的性能提升。**

👉 其中 `LzInt32`、`LzInt64` 提升最明显，几乎达到 **1.7× 加速**。

### 关于 Atoi 的 1ns 差异

在 `Atoi` 基准项中：

- LZCONVX：**7.106 ns/op**
- strconv.Atoi：**6.917 ns/op**

两者相差约 **0.19ns**（不到 0.2ns）。

这是因为：

> LZCONVX 内部会进行一次 fast-path / slow-path 判断，  
> 以保证行为与官方 `Atoi` 完全一致，并保持对错误格式的严格校验。

尽管多了这一道判断，整体性能依然几乎持平，并且在绝大多数实际场景中仍然表现更稳定。


## ✨ 特性亮点

### 🚀 超高性能（比官方快）
Fast-Path 下零分配、分支优化，使解析速度大幅提升。

### 🧩 行为完全对齐 `strconv`
错误类型、空白规则、格式规则全部一致，可安全替换。

### ⚡ 零内存分配
所有 Fast-Path 都不产生 GC 压力。

### 🛡 严格的边界检查
- 溢出 → `ErrRange`
- 格式错误 → `ErrSyntax`
- 前后空白自动裁剪
- 支持可选符号 `+` / `-`

### 📦 单文件即可使用
可直接复制到任何项目，无需额外依赖。



## 🧰 支持的整数解析 API

LzInt8
LzInt16
LzInt32
LzInt64
LzAtoi

整数解析行为说明
•	支持前后空白裁剪
•	支持 + / -
•	只接受十进制数字
•	数字越界 → ErrRange
•	非数字字符 → ErrSyntax
•	无内存分配（Zero-Alloc Fast Path）

## 🧮 浮点解析 API

LzFloat32
LzFloat64

浮点解析行为说明
•	支持前后空白裁剪
•	支持 + / -
•	支持小数点
•	支持科学计数法（e / E）
•	科学计数法指数支持正负号
•	多个小数点、格式错误 → ErrSyntax
•	指数过大/过小 → ErrRange

已广泛验证与官方 strconv.ParseFloat 行为一致。


## 🧪 使用示例
```go
package main

import (
	"fmt"
	"github.com/JohnLyonX/lzconvx/conv"
)

func main() {
	// 整型解析
	if v, err := conv.LzInt16(" 32767 "); err == nil {
		fmt.Println("int16:", v)
	}

	// 浮点解析
	if f, err := conv.LzFloat64("1.5e2"); err == nil {
		fmt.Println("float64:", f)
	}
}

```


## 🧱 项目结构
```bash
lzconvx/
 ├── conv/
 ├── int.go        # 整型解析核心
 ├── float.go      # 浮点解析核心
 ├── conv_test/
 │     ├── official_or_lzconv_float_test.go        # 官方与Lzconvx浮点数基准测试
 │     ├── official_or_lzconv_test.go      # 官方与Lzconvx整数数基准测试
 ├── README.md
 ├── README_CN.md
 └── LICENSE
 └── go.mod
```


## 🧭 Roadmap
	•	添加无符号整型（uint32/uint64）
	•	添加更快的浮点 Fast-Path 优化
	•	针对 amd64 添加 SSE/AVX 优化
	•	添加 WebAssembly（WASM）版本
	•	覆盖更多 fuzz 测试（与 strconv 行为对齐）
	•	提交 Go 官方 Proposal（标准库 Fast-Path 优化）


## 🤝 参与贡献

欢迎提交：
•	Bug 报告
•	行为与官方对齐的测试用例
•	性能优化建议
•	新的基准测试或实现思路

你的贡献将帮助我们进一步提升性能与兼容性。


## 📄 License
LZCONVX is released under the MIT License.  
Copyright © 2025 Liang Zhanbo.


## ✨ 作者

Liang Zhanbo（梁展波）
Creator of LZCONVX
GitHub: https://github.com/JohnLyonX

如你有 Go 性能优化、编译器、解析器相关问题，欢迎交流！
