[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/rsazh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/rsazh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/rsazh)](https://pkg.go.dev/github.com/yylego/rsazh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/rsazh/main.svg)](https://coveralls.io/github/yylego/rsazh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/rsazh.svg)](https://github.com/yylego/rsazh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/rsazh)](https://goreportcard.com/report/github.com/yylego/rsazh)

# rsazh

使用中文命名的 RSA PKCS#1 v1.5 加密操作库

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 免责声明

使用中文编写 Go 代码在技术上可行，但在实际工程中极不推荐。这种做法不适用于任何正式或商业场景，采用它的团队或公司可能面临同行的轻视和行业的负面评价，在商业公司中尤其可能成为被业内议论的对象。该项目仅供学习和研究，请勿在正式工程中采用。

---

## 主要特性

🔐 **RSA 加密**: 使用中文函数名的 PKCS#1 v1.5 加密和解密
🖋️ **数字签名**: 基于 SHA256 的签名和验签操作
🔑 **密钥管理**: 生成、加载、导出 PKCS#8/PKIX 格式的 RSA 密钥组件
📦 **简单 API**: 直观的中文命名方法封装 Go crypto/rsa 包
🛡️ **类型保护**: 独立的类型避免公私混淆

## 安装

```bash
go get github.com/yylego/rsazh
```

## 使用方法

### 基础加密和解密

此示例演示生成 RSA 密钥、加密消息和解密密文。

```go
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// 生成 2048 位 RSA 私钥
	v私钥, err := rsa15zh.R随机私钥(2048)
	must.Done(err)
	fmt.Println("生成私钥:", len(v私钥), "字节")

	// 从私钥提取公钥
	v公钥, err := rsa15zh.R获得公钥(v私钥)
	must.Done(err)
	fmt.Println("提取公钥:", len(v公钥), "字节")

	// 加载密钥
	r私钥, err := rsa15zh.F装载私钥(v私钥)
	must.Done(err)
	r公钥, err := rsa15zh.F装载公钥(v公钥)
	must.Done(err)

	// 加密测试
	message := "你好 RSA!"
	fmt.Println("\n原始消息:", message)

	v密文, err := r公钥.M加密([]byte(message))
	must.Done(err)
	fmt.Println("已加密:", base64.StdEncoding.EncodeToString(v密文)[:50]+"...")

	// 解密测试
	v明文, err := r私钥.M解密(v密文)
	must.Done(err)
	fmt.Println("已解密:", string(v明文))

	// 导出密钥
	exportedPrivate, err := r私钥.B导出()
	must.Done(err)
	exportedPublic, err := r公钥.B导出()
	must.Done(err)

	fmt.Println("\n导出私钥:", len(exportedPrivate), "字节")
	fmt.Println("导出公钥:", len(exportedPublic), "字节")
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### 数字签名和验证

此示例展示如何使用 RSA 密钥组件对文档进行签名和验签。

```go
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// 生成密钥
	v私钥, err := rsa15zh.R随机私钥(2048)
	must.Done(err)

	r私钥, err := rsa15zh.F装载私钥(v私钥)
	must.Done(err)

	// 签名消息
	message := "重要文档"
	fmt.Println("待签名消息:", message)

	v签名, err := r私钥.M签名([]byte(message))
	must.Done(err)
	fmt.Println("签名:", base64.StdEncoding.EncodeToString(v签名)[:50]+"...")

	// 从私钥提取公钥
	r公钥 := r私钥.P公钥()

	// 验证签名
	err = r公钥.M验签([]byte(message), v签名)
	if err != nil {
		fmt.Println("验证失败:", err)
	} else {
		fmt.Println("验证成功: 签名真实有效")
	}

	// 测试篡改消息
	tamperedMessage := "重要文档!"
	fmt.Println("\n篡改消息:", tamperedMessage)
	err = r公钥.M验签([]byte(tamperedMessage), v签名)
	if err != nil {
		fmt.Println("验证失败，符合预期:", err)
	} else {
		fmt.Println("验证成功: 签名真实有效")
	}
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

## API 参考

### 密钥生成函数

| 函数 | 描述 (ZH) | Description (EN) |
|------|-----------|-----------------|
| `R随机私钥(n位数 int)` | 生成新的 RSA 私钥 | Generates new RSA private key |
| `R获得公钥(privateKeyBytes []byte)` | 从私钥字节中提取公钥 | Extracts public key from private key bytes |
| `F装载私钥(v私钥 []byte)` | 从 PKCS#8 字节加载私钥 | Loads private key from PKCS#8 bytes |
| `F装载公钥(v公钥 []byte)` | 从 PKIX 字节加载公钥 | Loads public key from PKIX bytes |

### 私钥方法 (Rsa私钥)

| 方法 | 描述 (ZH) | Description (EN) |
|------|-----------|-----------------|
| `M签名(v明文 []byte)` | 使用 SHA256 对明文签名 | Signs plaintext using SHA256 |
| `M解密(v密文 []byte)` | 解密密文 | Decrypts ciphertext |
| `B导出()` | 导出私钥为 PKCS#8 字节 | Exports private key as PKCS#8 bytes |
| `P公钥()` | 从私钥中提取公钥 | Extracts public key from private key |

### 公钥方法 (Rsa公钥)

| 方法 | 描述 (ZH) | Description (EN) |
|------|-----------|-----------------|
| `M加密(v明文 []byte)` | 加密明文 | Encrypts plaintext |
| `M验签(v明文 []byte, v签名 []byte)` | 使用 SHA256 验证签名 | Verifies signature using SHA256 |
| `B导出()` | 导出公钥为 PKIX 字节 | Exports public key as PKIX bytes |

## 示例

### 完整工作流和密钥持久化

**生成并保存密钥:**
```go
v私钥bytes, err := rsa15zh.R随机私钥(2048)
私钥String := base64.StdEncoding.EncodeToString(v私钥bytes)
// 将 私钥String 保存到数据库/文件
```

**加载并使用密钥:**
```go
v私钥restored, _ := base64.StdEncoding.DecodeString(私钥String)
r私钥, _ := rsa15zh.F装载私钥(v私钥restored)
// 使用 r私钥 进行签名或解密
```

**提取公钥:**
```go
r公钥 := r私钥.P公钥()
v导出, _ := r公钥.B导出()
// 将 v导出 分享给他人
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

## 实现细节

### 加密方案
- **算法**: RSA with PKCS#1 v1.5 填充
- **长度**: 支持 2048, 3072, 4096 位（推荐 2048）
- **格式**: PKCS#8 (私钥组件), PKIX (公钥组件)

### 签名方案
- **哈希函数**: SHA256
- **签名算法**: RSA PKCS#1 v1.5 签名
- **输出**: Base64 编码的签名字节

## 命名规范

- `R` 前缀: 随机生成函数 (R随机私钥, R获得公钥)
- `F` 前缀: 加载/初始化函数 (F装载私钥, F装载公钥)
- `M` 前缀: 主要操作方法 (M加密, M解密, M签名, M验签)
- `B` 前缀: 字节导出方法 (B导出)
- `P` 前缀: 提取方法 (P公钥)

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
