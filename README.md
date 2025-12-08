# 🚀 LZCONVX — A Faster & More Efficient Integer and Float Parsing Library for Go

[English Version](./README.md) | [中文文档](./README_CN.md)

LZCONVX is a highly optimized numeric parsing library designed with **performance and reliability** in mind.  
While maintaining full behavioral compatibility with Go’s official `strconv`, it provides significantly faster  
parsing through hand-crafted fast paths, zero allocations, and optimized branch layouts.

Suitable for:

- High-concurrency backend services
- RPC frameworks
- Monitoring and telemetry systems
- Logging platforms
- Network protocol parsers
- Game servers
- Performance-critical low-level components


## 📊 Benchmark (Apple M1 Max · Go 1.25)

The following benchmark was executed on **Apple M1 Max**:

<img width="1212" height="580" alt="benchmark" src="https://github.com/user-attachments/assets/6a536c0a-025d-4283-81f7-3a194afbc6f1" />

Below is the exact benchmark result comparing LZCONVX with Go's official `strconv`:

| Benchmark | LZCONVX | strconv (official) | Difference |
|-----------|---------|--------------------|------------|
| Int8  | **7.117 ns/op** | 8.455 ns/op | ⚡ ~19.8% faster |
| Int16 | **8.628 ns/op** | 10.37 ns/op | ⚡ ~17% faster |
| Int32 | **7.933 ns/op** | 14.10 ns/op | ⚡ ~43.7% faster |
| Int64 | **12.15 ns/op** | 21.33 ns/op | ⚡ ~43.0% faster |
| Atoi  | **7.106 ns/op** | 6.917 ns/op | ⚠ ~0.2ns slower (explained below) |

### Performance Summary

👉 **LZCONVX achieves ~17% to ~43% performance improvements compared to `strconv.ParseInt`.**

👉 `LzInt32` & `LzInt64` show the strongest acceleration, nearly **1.7× faster**.

### Why Atoi Is ~0.2ns Slower

- LZCONVX: **7.106 ns/op**
- strconv.Atoi: **6.917 ns/op**

The difference is only **0.19ns**.

Reason:

> LZCONVX performs an internal fast-path / slow-path check  
> to ensure strict behavioral consistency with `strconv.Atoi`,  
> especially for malformed or edge-case inputs.

Despite this additional check, overall performance remains nearly identical,  
and still extremely stable in real-world environments.


## ✨ Key Features

### 🚀 High Performance (Faster Than Official)
Zero-allocation fast paths and optimized branching significantly boost performance.

### 🧩 Fully Compatible With `strconv`
Error types, whitespace rules, and formatting behavior all match the official implementation.

### ⚡ Zero Allocations
All fast paths generate **no heap allocations**, minimizing GC overhead.

### 🛡 Strict Boundary Checking
- Overflow → `ErrRange`
- Invalid format → `ErrSyntax`
- Automatically trims surrounding whitespace
- Supports optional sign `+` / `-`

### 📦 Single-File Usage
Can be copied directly into any project, no dependencies.


## 🧰 Integer Parsing API

```
LzInt8
LzInt16
LzInt32
LzInt64
LzAtoi
```

Integer parsing behavior:
- Trims whitespace
- Supports `+` / `-`
- Decimal digits only
- Overflow → `ErrRange`
- Non-digit → `ErrSyntax`
- Zero-allocation fast path


## 🧮 Floating-Point Parsing API

```
LzFloat32
LzFloat64
```

Float parsing behavior:
- Trims whitespace
- Supports `+` / `-`
- Supports decimal points
- Supports scientific notation (`e` / `E`)
- Exponent supports sign
- Multiple dots or invalid exponent → `ErrSyntax`
- Exponent too large/small → `ErrRange`

Fully validated against `strconv.ParseFloat`.


## 🧪 Usage Example
```go
package main

import (
    "fmt"
    "github.com/JohnLyonX/lzconvx/conv"
)

func main() {
    // Integer parsing
    if v, err := conv.LzInt16(" 32767 "); err == nil {
        fmt.Println("int16:", v)
    }

    // Float parsing
    if f, err := conv.LzFloat64("1.5e2"); err == nil {
        fmt.Println("float64:", f)
    }
}
```

## 🧱 Project Structure
```bash
lzconvx/
 ├── int.go                   # Core integer parsing
 ├── float.go                 # Core float parsing
 ├── conv_test/
 │     ├── official_or_lzconv_float_test.go  # Float benchmark comparison
 │     ├── official_or_lzconv_test.go        # Integer benchmark comparison
 ├── README.md
 ├── README_CN.md
 ├── LICENSE
 └── go.mod
```


## 🧭 Roadmap
- Add unsigned integers (`uint32`/`uint64`)
- Add faster float parsing fast-path
- Add SSE/AVX optimizations for amd64
- Add WebAssembly (WASM) version
- Expand fuzz testing (fully aligned with `strconv`)
- Submit Go Proposal for standard library fast-path improvements


## 🤝 Contributing

Contributions are welcome:
- Bug reports
- Behavior-matching test cases
- Performance optimization ideas
- New benchmarks or parsing strategies

Your help will further improve performance and compatibility.


## 📄 License
LZCONVX is released under the MIT License.  
Copyright © 2025 Liang Zhanbo.


## ✨ Author

**Liang Zhanbo (JohnLyonX)**  
Creator of LZCONVX  
GitHub: https://github.com/JohnLyonX

If you're interested in Go performance tuning, compiler internals,  
or numeric parser design, feel free to reach out!
