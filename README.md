# 🚀 LZCONVX — 一个比官方 `strconv` 更高效、更快速的整数与浮点解析库

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

---

## 📊 基准测试（Apple M1 Max）

以下基准测试在 **Apple M1 Max** 上运行：

<img width="1212" height="580" alt="benchmark" src="https://github.com/user-attachments/assets/6a536c0a-025d-4283-81f7-3a194afbc6f1" />

👉 LZCONVX 在多组测试下比 `strconv.ParseInt` 快 **约 1.7×**。


# ✨ 特性亮点

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

---

# 🧰 支持的整数解析 API

```go
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

🧮 浮点解析 API

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

⸻

🧪 使用示例

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


⸻

🧱 项目结构

lzconvx/
 ├── conv/
 │     ├── int.go        # 整型解析核心
 │     ├── float.go      # 浮点解析核心
 │     ├── trim.go       # 空白处理
 │     ├── errors.go     # ErrSyntax / ErrRange
 │     ├── fast.go       # Fast-Path 逻辑
 │     ├── bench_test.go # Benchmark
 ├── README.md
 ├── README_CN.md
 └── LICENSE


⸻

🧭 Roadmap
	•	添加无符号整型（uint32/uint64）
	•	添加更快的浮点 Fast-Path 优化
	•	针对 amd64 添加 SSE/AVX 优化
	•	添加 WebAssembly（WASM）版本
	•	覆盖更多 fuzz 测试（与 strconv 行为对齐）
	•	提交 Go 官方 Proposal（标准库 Fast-Path 优化）

⸻

🤝 参与贡献

欢迎提交：
	•	Bug 报告
	•	行为与官方对齐的测试用例
	•	性能优化建议
	•	新的基准测试或实现思路

你的贡献将帮助我们进一步提升性能与兼容性。

⸻

📄 License

MIT License.
可安全用于商业项目。

⸻

✨ 作者

Liang Zhanbo（梁展波）
Creator of LZCONVX
GitHub: https://github.com/JohnLyonX

如你有 Go 性能优化、编译器、解析器相关问题，欢迎交流！
### 🔥 写 Go 官方提案草稿（Proposal Draft）
