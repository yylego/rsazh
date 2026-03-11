// Package main demonstrates RSA digital signature and verification operations
// Shows signing process, signature validation, and tampered message detection
//
// main 演示 RSA 数字签名和验证操作
// 展示签名过程、签名验证和篡改消息检测
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// Demo: RSA digital signature and verification (RSA 数字签名和验证演示)

	// Generate keys (生成密钥)
	v私钥, err := rsa15zh.R随机私钥(2048)
	must.Done(err)

	r私钥, err := rsa15zh.F装载私钥(v私钥)
	must.Done(err)

	// Sign message (签名消息)
	message := "Important document"
	fmt.Println("Message to sign (待签名消息):", message)

	v签名, err := r私钥.M签名([]byte(message))
	must.Done(err)
	fmt.Println("Signature (签名):", base64.StdEncoding.EncodeToString(v签名)[:50]+"...")

	// Extract public key from private key (从私钥提取公钥)
	r公钥 := r私钥.P公钥()

	// Verify signature (验证签名)
	err = r公钥.M验签([]byte(message), v签名)
	if err != nil {
		fmt.Println("Verification failed (验证失败):", err)
	} else {
		fmt.Println("Verification succeeded (验证成功): signature is authentic")
	}

	// Test with tampered message (测试篡改消息)
	tamperedMessage := "Important document!"
	fmt.Println("\nTampered message (篡改消息):", tamperedMessage)
	err = r公钥.M验签([]byte(tamperedMessage), v签名)
	if err != nil {
		fmt.Println("Verification failed as expected (验证失败，符合预期):", err)
	} else {
		fmt.Println("Verification succeeded (验证成功): signature is authentic")
	}

	// Export keys test (导出密钥测试)
	v导出私钥, err := r私钥.B导出()
	must.Done(err)
	v导出公钥, err := r公钥.B导出()
	must.Done(err)

	fmt.Println("\nExported keys (导出密钥):")
	fmt.Println("Private key size (私钥大小):", len(v导出私钥), "bytes")
	fmt.Println("Public key size (公钥大小):", len(v导出公钥), "bytes")
}
