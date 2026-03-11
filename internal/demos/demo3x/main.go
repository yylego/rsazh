// Package main demonstrates complete RSA workflow with key persistence
// Shows key generation, storage, loading, encryption, decryption, signing and verification
//
// main 演示完整的 RSA 工作流程和密钥持久化
// 展示密钥生成、存储、加载、加密、解密、签名和验证
package main

import (
	"encoding/base64"
	"fmt"

	"github.com/yylego/rsazh/rsa15zh"
	"github.com/yylego/must"
)

func main() {
	// Demo: Complete RSA workflow with key persistence (完整 RSA 工作流和密钥持久化演示)

	// Step 1: Generate and export keys (第1步：生成并导出密钥)
	fmt.Println("=== Key Generation (密钥生成) ===")
	v私钥bytes, err := rsa15zh.R随机私钥(2048)
	must.Done(err)

	v公钥bytes, err := rsa15zh.R获得公钥(v私钥bytes)
	must.Done(err)

	// Encode to base64 strings (编码为 base64 字符串)
	私钥String := base64.StdEncoding.EncodeToString(v私钥bytes)
	公钥String := base64.StdEncoding.EncodeToString(v公钥bytes)

	fmt.Println("Private key (base64) (私钥):", 私钥String[:60]+"...")
	fmt.Println("Public key (base64) (公钥):", 公钥String[:60]+"...")

	// Step 2: Load keys from base64 strings (第2步：从 base64 字符串加载密钥)
	fmt.Println("\n=== Key Loading (密钥加载) ===")
	v私钥restored, err := base64.StdEncoding.DecodeString(私钥String)
	must.Done(err)
	v公钥restored, err := base64.StdEncoding.DecodeString(公钥String)
	must.Done(err)

	r私钥, err := rsa15zh.F装载私钥(v私钥restored)
	must.Done(err)
	r公钥, err := rsa15zh.F装载公钥(v公钥restored)
	must.Done(err)

	// Step 3: Encrypt and decrypt (第3步：加密和解密)
	fmt.Println("\n=== Encryption/Decryption (加密/解密) ===")
	plaintext := "Confidential data 机密数据"
	fmt.Println("Plaintext (明文):", plaintext)

	ciphertext, err := r公钥.M加密([]byte(plaintext))
	must.Done(err)
	fmt.Println("Ciphertext (密文):", base64.StdEncoding.EncodeToString(ciphertext)[:60]+"...")

	decrypted, err := r私钥.M解密(ciphertext)
	must.Done(err)
	fmt.Println("Decrypted (解密):", string(decrypted))

	// Step 4: Sign and verify (第4步：签名和验证)
	fmt.Println("\n=== Signing/Verification (签名/验证) ===")
	document := "Contract version 1.0"
	fmt.Println("Document (文档):", document)

	signature, err := r私钥.M签名([]byte(document))
	must.Done(err)
	fmt.Println("Signature (签名):", base64.StdEncoding.EncodeToString(signature)[:60]+"...")

	err = r公钥.M验签([]byte(document), signature)
	if err == nil {
		fmt.Println("Verification (验证): ✓ Signature is authentic (签名真实有效)")
	} else {
		fmt.Println("Verification (验证): ✗ Signature is invalid (签名无效)")
	}

	// Step 5: Extract public key from private key (第5步：从私钥提取公钥)
	fmt.Println("\n=== Public Key Extraction (公钥提取) ===")
	r公钥FromPrivate := r私钥.P公钥()
	err = r公钥FromPrivate.M验签([]byte(document), signature)
	if err == nil {
		fmt.Println("Extracted public key works (提取的公钥有效): ✓")
	}
}
