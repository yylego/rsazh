<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/rsazh/release.yml?branch=main&label=BUILD)](https://github.com/yylego/rsazh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/rsazh)](https://pkg.go.dev/github.com/yylego/rsazh)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/rsazh/main.svg)](https://coveralls.io/github/yylego/rsazh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.26%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/rsazh.svg)](https://github.com/yylego/rsazh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/rsazh)](https://goreportcard.com/report/github.com/yylego/rsazh)
<!-- TEMPLATE (EN) CLOSE: BADGES -->

# rsazh

Chinese-named package providing RSA PKCS#1 v1.5 encryption operations

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

---

## DISCLAIMER

Writing Go code in Chinese is a viable technique, but something to avoid in production engineering. This approach should not be used in serious and business settings. Teams and companies that embrace it could face contempt from peers and negative judgment across the profession. In business companies, this practice is even more prone to becoming a target of public criticism. This project is dedicated to research and academic studies. Do not use this approach in production.

---

## Main Features

- 🔐 **RSA Encryption**: PKCS#1 v1.5 encryption and decryption with Chinese function names
- 🖋️ **Signatures**: SHA256-based signing and verification operations
- 🔑 **RSA Components**: Generate, load, export RSA components in PKCS#8/PKIX formats
- 📦 **Simple API**: Intuitive Chinese-named methods wrapping Go crypto/rsa package
- 🛡️ **Type Protection**: Separate types eliminating private/public confusion

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
	v私钥, err := rsa15zh.R随机PKCS8私钥(2048)
	must.Done(err)
	fmt.Println("Generated private key:", len(v私钥), "bytes")

	// Extract public key from private key
	v公钥, err := rsa15zh.R获得PKIX公钥(v私钥)
	must.Done(err)
	fmt.Println("Extracted public key:", len(v公钥), "bytes")

	// Load keys
	r私钥, err := rsa15zh.F装载PKCS8私钥(v私钥)
	must.Done(err)
	r公钥, err := rsa15zh.F装载PKIX公钥(v公钥)
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
	exportedPrivate, err := r私钥.B导出PKCS8()
	must.Done(err)
	exportedPublic, err := r公钥.B导出PKIX()
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
	v私钥, err := rsa15zh.R随机PKCS8私钥(2048)
	must.Done(err)

	r私钥, err := rsa15zh.F装载PKCS8私钥(v私钥)
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

| Function                                | Description (EN)                                           | 描述 (ZH)                            |
| --------------------------------------- | ---------------------------------------------------------- | ------------------------------------ |
| `R随机PKCS8私钥(n位数 int)`             | Generates PKCS#8 DER private bytes                         | 生成 PKCS#8 DER 私钥字节             |
| `R获得PKIX公钥(privateKeyBytes []byte)` | Converts PKCS#8 DER private bytes to PKIX DER public bytes | 从 PKCS#8 DER 私钥提取 PKIX DER 公钥 |
| `F装载PKCS8私钥(v私钥 []byte)`          | Loads private components from PKCS#8 DER bytes             | 从 PKCS#8 DER 字节加载私钥           |
| `F装载PKIX公钥(v公钥 []byte)`           | Loads public components from PKIX DER bytes                | 从 PKIX DER 字节加载公钥             |

### Private Key Methods (Rsa私钥)

`F装载PKCS1私钥Base64` loads standard Base64 encoded PKCS#1 DER private keys. It is distinct from `F装载PKCS8私钥`, which accepts PKCS#8 DER bytes. `M解密PKCS1Base64(privateBase64, ciphertextBase64)` combines loading and decryption; an existing private instance can use `M解密Base64(ciphertextBase64)`. Both use the existing PKCS#1 v1.5 scheme and return decoding / parsing / decryption errors without logging private data.

| Method                      | Description (EN)                                                 | 描述 (ZH)                               |
| --------------------------- | ---------------------------------------------------------------- | --------------------------------------- |
| `M签名(v明文 []byte)`       | Signs plaintext using SHA256                                     | 使用 SHA256 对明文签名                  |
| `M解密(v密文 []byte)`       | Decrypts ciphertext                                              | 解密密文                                |
| `M解密Base64(s密文 string)` | Decodes standard Base64 ciphertext and decrypts to bytes         | 解码标准 Base64 密文并解密为字节        |
| `B导出PKCS1Base64()`        | Exports PKCS#1 DER as standard Base64 for `F装载PKCS1私钥Base64` | 导出与装载函数配对的 PKCS#1 Base64 私钥 |
| `B导出PKCS8()`              | Exports private components as PKCS#8 DER bytes                   | 导出私钥为 PKCS#8 DER 字节              |
| `P公钥()`                   | Extracts public key from private key                             | 从私钥中提取公钥                        |

### Public Key Methods (Rsa公钥)

| Method                              | Description (EN)                                      | 描述 (ZH)                          |
| ----------------------------------- | ----------------------------------------------------- | ---------------------------------- |
| `M加密(v明文 []byte)`               | Encrypts plaintext                                    | 加密明文                           |
| `M加密Base64(v明文 []byte)`         | Encrypts bytes and returns standard Base64 ciphertext | 加密明文字节并返回标准 Base64 密文 |
| `M验签(v明文 []byte, v签名 []byte)` | Verifies signature using SHA256                       | 使用 SHA256 验证签名               |
| `B导出PKIX()`                       | Exports public components as PKIX DER bytes           | 导出公钥为 PKIX DER 字节           |

## Examples

### Complete Workflow with Key Persistence

**Generate and save keys:**

```go
v私钥bytes, err := rsa15zh.R随机PKCS8私钥(2048)
must.Done(err)
私钥String := base64.StdEncoding.EncodeToString(v私钥bytes)
// Save 私钥String to database/file
```

**Load and use keys:**

```go
v私钥restored, err := base64.StdEncoding.DecodeString(私钥String)
must.Done(err)
r私钥, err := rsa15zh.F装载PKCS8私钥(v私钥restored)
must.Done(err)
// Use r私钥 to sign or decrypt
```

**Extract public key:**

```go
r公钥 := r私钥.P公钥()
v导出, err := r公钥.B导出PKIX()
must.Done(err)
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
- **Output**: Raw signature bytes; examples encode them as Base64 for display

## Naming Conventions

- `R` prefix: Generation and derivation functions (R随机PKCS8私钥, R获得PKIX公钥)
- `F` prefix: Loading functions (F装载PKCS8私钥, F装载PKIX公钥)
- `M` prefix: Main operation methods (M加密, M解密, M签名, M验签)
- `B` prefix: Export methods (B导出PKCS8, B导出PKIX, B导出PKCS1Base64)
- `P` prefix: Extraction methods (P公钥)

PKCS8 / PKCS1 / PKIX names identify the DER format. A Base64 suffix identifies standard Base64 text; it does not mean PEM. Byte encryption and decryption methods retain `M加密` / `M解密`.

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

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
7. **Documentation**: Update documentation to support client-facing changes
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

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->

---

<!-- TEMPLATE (EN) BEGIN: GITHUB STARS -->

## GitHub Stars

[![Stargazers](https://starchart.cc/yylego/rsazh.svg?variant=adaptive)](https://starchart.cc/yylego/rsazh)
<!-- TEMPLATE (EN) CLOSE: GITHUB STARS -->
