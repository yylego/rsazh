package rsa15zh

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// F装载PKCS1私钥Base64 decodes standard Base64 containing PKCS#1 DER bytes.
// F装载PKCS1私钥Base64 接收标准 Base64 编码的 PKCS#1 DER 私钥，不接收 PEM 或 PKCS#8。
func F装载PKCS1私钥Base64(s私钥 string) (*Rsa私钥, error) {
	v私钥, err := base64.StdEncoding.DecodeString(s私钥)
	if err != nil {
		return nil, fmt.Errorf("decode PKCS1 base64: %w", err)
	}
	pri, err := x509.ParsePKCS1PrivateKey(v私钥)
	if err != nil {
		return nil, fmt.Errorf("parse PKCS1: %w", err)
	}
	return New私钥(pri), nil
}

// B导出PKCS1Base64 encodes PKCS#1 DER bytes as standard Base64.
// B导出PKCS1Base64 导出标准 Base64 编码的 PKCS#1 DER 私钥，与 F装载PKCS1私钥Base64 配对。
func (r *Rsa私钥) B导出PKCS1Base64() string {
	return base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(r.pri))
}

// M加密Base64 encrypts bytes with PKCS#1 v1.5 and returns standard Base64.
// M加密Base64 加密明文字节，返回标准 Base64 密文，与 M解密Base64 配对。
func (r *Rsa公钥) M加密Base64(v明文 []byte) (string, error) {
	v密文, err := r.M加密(v明文)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(v密文), nil
}

// M解密Base64 decodes standard Base64 and decrypts with PKCS#1 v1.5.
// M解密Base64 解码标准 Base64 密文，再按已有 PKCS#1 v1.5 方案解密。
func (r *Rsa私钥) M解密Base64(s密文 string) ([]byte, error) {
	v密文, err := base64.StdEncoding.DecodeString(s密文)
	if err != nil {
		return nil, fmt.Errorf("decode ciphertext base64: %w", err)
	}
	return r.M解密(v密文)
}

// M解密PKCS1Base64 combines PKCS#1 private DER decoding and ciphertext decryption.
// M解密PKCS1Base64 用标准 Base64 编码的 PKCS#1 DER 私钥解密标准 Base64 密文。
func M解密PKCS1Base64(s私钥, s密文 string) ([]byte, error) {
	pri, err := F装载PKCS1私钥Base64(s私钥)
	if err != nil {
		return nil, err
	}
	return pri.M解密Base64(s密文)
}
