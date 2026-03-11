// Package main demonstrates basic RSA encryption and decryption operations
// Shows complete workflow: key generation, encryption, decryption, and export
//
// main 演示基础 RSA 加密和解密操作
// 展示完整工作流程：密钥生成、加密、解密和导出
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// Demo: Basic RSA encryption and decryption (基础 RSA 加密解密演示)

	// Generate 2048-bit RSA private key (生成 2048 位 RSA 私钥)
	v私钥, err := rsa15zh.R随机私钥(2048)
	must.Done(err)
	fmt.Println("Generated private key (生成私钥):", len(v私钥), "bytes")

	// Extract public key from private key (从私钥提取公钥)
	v公钥, err := rsa15zh.R获得公钥(v私钥)
	must.Done(err)
	fmt.Println("Extracted public key (提取公钥):", len(v公钥), "bytes")

	// Load keys (加载密钥)
	r私钥, err := rsa15zh.F装载私钥(v私钥)
	must.Done(err)
	r公钥, err := rsa15zh.F装载公钥(v公钥)
	must.Done(err)

	// Encryption test (加密测试)
	message := "Hello RSA!"
	fmt.Println("\nOriginal message (原始消息):", message)

	v密文, err := r公钥.M加密([]byte(message))
	must.Done(err)
	fmt.Println("Encrypted (已加密):", base64.StdEncoding.EncodeToString(v密文)[:50]+"...")

	// Decryption test (解密测试)
	v明文, err := r私钥.M解密(v密文)
	must.Done(err)
	fmt.Println("Decrypted (已解密):", string(v明文))

	// Export keys (导出密钥)
	exportedPrivate, err := r私钥.B导出()
	must.Done(err)
	exportedPublic, err := r公钥.B导出()
	must.Done(err)

	fmt.Println("\nExported private key (导出私钥):", len(exportedPrivate), "bytes")
	fmt.Println("Exported public key (导出公钥):", len(exportedPublic), "bytes")
}
