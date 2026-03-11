[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/rsazh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/rsazh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/rsazh)](https://pkg.go.dev/github.com/yylego/rsazh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/rsazh/main.svg)](https://coveralls.io/github/yylego/rsazh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/rsazh.svg)](https://github.com/yylego/rsazh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/rsazh)](https://goreportcard.com/report/github.com/yylego/rsazh)

# rsazh

Chinese-named package providing RSA PKCS#1 v1.5 encryption operations

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice is even more prone to becoming a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

---

## Main Features

🔐 **RSA Encryption**: PKCS#1 v1.5 encryption and decryption with Chinese function names
🖋️ **Digital Signatures**: SHA256-based signing and verification operations
🔑 **Cryptographic Material Management**: Generate, load, export RSA components in PKCS#8/PKIX formats
📦 **Simple API**: Intuitive Chinese-named methods wrapping Go crypto/rsa package
🛡️ **Type Protection**: Separate types eliminating private/public confusion

## Installation

```bash
go get github.com/yylego/rsazh
```

## Usage

### Basic Encryption and Decryption

This example demonstrates generating RSA keys, encrypting messages and decrypting ciphertext.

```go
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// Generate 2048-bit RSA private key
	v私钥, err := rsa15zh.R随机私钥(2048)
	must.Done(err)
	fmt.Println("Generated private key:", len(v私钥), "bytes")

	// Extract public key from private key
	v公钥, err := rsa15zh.R获得公钥(v私钥)
	must.Done(err)
	fmt.Println("Extracted public key:", len(v公钥), "bytes")

	// Load keys
	r私钥, err := rsa15zh.F装载私钥(v私钥)
	must.Done(err)
	r公钥, err := rsa15zh.F装载公钥(v公钥)
	must.Done(err)

	// Encryption test
	message := "Hello RSA!"
	fmt.Println("\nOriginal message:", message)

	v密文, err := r公钥.M加密([]byte(message))
	must.Done(err)
	fmt.Println("Encrypted:", base64.StdEncoding.EncodeToString(v密文)[:50]+"...")

	// Decryption test
	v明文, err := r私钥.M解密(v密文)
	must.Done(err)
	fmt.Println("Decrypted:", string(v明文))

	// Export keys
	exportedPrivate, err := r私钥.B导出()
	must.Done(err)
	exportedPublic, err := r公钥.B导出()
	must.Done(err)

	fmt.Println("\nExported private key:", len(exportedPrivate), "bytes")
	fmt.Println("Exported public key:", len(exportedPublic), "bytes")
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### Digital Signatures and Validation

This example shows how to sign documents and validate signatures using RSA cryptographic components.

```go
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// Generate keys
	v私钥, err := rsa15zh.R随机私钥(2048)
	must.Done(err)

	r私钥, err := rsa15zh.F装载私钥(v私钥)
	must.Done(err)

	// Sign message
	message := "Important document"
	fmt.Println("Message to sign:", message)

	v签名, err := r私钥.M签名([]byte(message))
	must.Done(err)
	fmt.Println("Signature:", base64.StdEncoding.EncodeToString(v签名)[:50]+"...")

	// Extract public key from private key
	r公钥 := r私钥.P公钥()

	// Verify signature
	err = r公钥.M验签([]byte(message), v签名)
	if err != nil {
		fmt.Println("Verification failed:", err)
	} else {
		fmt.Println("Verification succeeded: signature is authentic")
	}

	// Test with tampered message
	tamperedMessage := "Important document!"
	fmt.Println("\nTampered message:", tamperedMessage)
	err = r公钥.M验签([]byte(tamperedMessage), v签名)
	if err != nil {
		fmt.Println("Verification failed as expected:", err)
	} else {
		fmt.Println("Verification succeeded: signature is authentic")
	}
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

## API Reference

### Key Generation Functions

| Function | Description (EN) | 描述 (ZH) |
|----------|-----------------|-----------|
| `R随机私钥(n位数 int)` | Generates new RSA private key | 生成新的 RSA 私钥 |
| `R获得公钥(privateKeyBytes []byte)` | Extracts public key from private key bytes | 从私钥字节中提取公钥 |
| `F装载私钥(v私钥 []byte)` | Loads private key from PKCS#8 bytes | 从 PKCS#8 字节加载私钥 |
| `F装载公钥(v公钥 []byte)` | Loads public key from PKIX bytes | 从 PKIX 字节加载公钥 |

### Private Key Methods (Rsa私钥)

| Method | Description (EN) | 描述 (ZH) |
|--------|-----------------|-----------|
| `M签名(v明文 []byte)` | Signs plaintext using SHA256 | 使用 SHA256 对明文签名 |
| `M解密(v密文 []byte)` | Decrypts ciphertext | 解密密文 |
| `B导出()` | Exports private key as PKCS#8 bytes | 导出私钥为 PKCS#8 字节 |
| `P公钥()` | Extracts public key from private key | 从私钥中提取公钥 |

### Public Key Methods (Rsa公钥)

| Method | Description (EN) | 描述 (ZH) |
|--------|-----------------|-----------|
| `M加密(v明文 []byte)` | Encrypts plaintext | 加密明文 |
| `M验签(v明文 []byte, v签名 []byte)` | Verifies signature using SHA256 | 使用 SHA256 验证签名 |
| `B导出()` | Exports public key as PKIX bytes | 导出公钥为 PKIX 字节 |

## Examples

### Complete Workflow with Key Persistence

**Generate and save keys:**
```go
v私钥bytes, err := rsa15zh.R随机私钥(2048)
私钥String := base64.StdEncoding.EncodeToString(v私钥bytes)
// Save 私钥String to database/file
```

**Load and use keys:**
```go
v私钥restored, _ := base64.StdEncoding.DecodeString(私钥String)
r私钥, _ := rsa15zh.F装载私钥(v私钥restored)
// Use r私钥 to sign or decrypt
```

**Extract public key:**
```go
r公钥 := r私钥.P公钥()
v导出, _ := r公钥.B导出()
// Share v导出 with others
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

## Implementation Details

### Encryption Scheme
- **Algorithm**: RSA with PKCS#1 v1.5 padding
- **Sizes**: Supports 2048, 3072, 4096 bits (2048 recommended)
- **Formats**: PKCS#8 (private components), PKIX (public components)

### Signature Scheme
- **Hash Function**: SHA256
- **Signature Algorithm**: RSA PKCS#1 v1.5 signature
- **Output**: Base64-encoded signature bytes

## Naming Conventions

- `R` prefix: Random generation functions (R随机私钥, R获得公钥)
- `F` prefix: Loading/initialization functions (F装载私钥, F装载公钥)
- `M` prefix: Main operation methods (M加密, M解密, M签名, M验签)
- `B` prefix: Bytes export methods (B导出)
- `P` prefix: Extraction methods (P公钥)

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->
