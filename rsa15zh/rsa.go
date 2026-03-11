// Package rsa15zh: RSA PKCS#1 v1.5 encryption operations with Chinese-named functions
// Implements complete RSA cryptographic operations: generation, encryption, decryption, signing and verification
// Applies SHA256 hash algorithm and PKCS#1 v1.5 padding scheme with standard PKCS#8/PKIX encoding
// Provides intuitive Chinese method names wrapping Go crypto/rsa package
//
// rsa15zh: 使用中文命名的 RSA PKCS#1 v1.5 加密操作包
// 实现完整的 RSA 密码学操作：生成、加密、解密、签名和验签
// 应用 SHA256 哈希算法和 PKCS#1 v1.5 填充方案，使用标准 PKCS#8/PKIX 编码
// 提供直观的中文方法名称，封装 Go crypto/rsa 包
package rsa15zh

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	"github.com/yylego/erero"
)

// Rsa私钥 wraps RSA private cryptographic components with Chinese-named methods
// Provides signing and decryption operations
//
// Rsa私钥 封装 RSA 私钥组件，提供中文命名的方法
// 提供签名和解密操作
type Rsa私钥 struct {
	pri *rsa.PrivateKey
}

// New私钥 creates a new Rsa私钥 instance from rsa.PrivateKey
// Returns the wrapped private cryptographic components instance
//
// New私钥 从 rsa.PrivateKey 创建新的 Rsa私钥 实例
// 返回封装后的私钥组件实例
func New私钥(pri *rsa.PrivateKey) *Rsa私钥 {
	return &Rsa私钥{pri: pri}
}

// M签名 signs the plaintext using SHA256 hash and PKCS#1 v1.5 signature scheme
// Takes plaintext bytes as input and generates digital signature
// Returns signature bytes on success, otherwise returns an exception
//
// M签名 使用 SHA256 哈希和 PKCS#1 v1.5 签名方案对明文进行签名
// 接收明文字节作为输入并生成数字签名
// 成功时返回签名字节，否则返回异常
func (r *Rsa私钥) M签名(v明文 []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(v明文)
	return rsa.SignPKCS1v15(rand.Reader, r.pri, crypto.SHA256, hash.Sum(nil))
}

// M解密 decrypts ciphertext using PKCS#1 v1.5 decryption scheme
// Takes encrypted ciphertext bytes as input and decrypts them
// Returns decrypted plaintext bytes on success, otherwise returns an exception
//
// M解密 使用 PKCS#1 v1.5 解密方案解密密文
// 接收加密的密文字节作为输入并解密
// 成功时返回解密后的明文字节，否则返回异常
func (r *Rsa私钥) M解密(v密文 []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, r.pri, v密文)
}

// B导出 exports private cryptographic components as PKCS#8 format bytes
// Serializes the private components to standard PKCS#8 encoding
// Returns serialized bytes on success, otherwise returns an exception
//
// B导出 将私钥组件导出为 PKCS#8 格式字节
// 将私钥组件序列化为标准 PKCS#8 编码
// 成功时返回序列化字节，否则返回异常
func (r *Rsa私钥) B导出() ([]byte, error) {
	// Export to PKCS#8 format (more cross-platform compatible than PKCS#1)
	// 导出为 PKCS#8 格式（比 PKCS#1 跨平台兼容性更佳）
	priBytes, err := x509.MarshalPKCS8PrivateKey(r.pri)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return priBytes, nil
}

// P公钥 extracts the public cryptographic components from this private one
// Returns the wrapped public components instance
//
// P公钥 从此私钥中提取公钥组件
// 返回封装后的公钥组件实例
func (r *Rsa私钥) P公钥() *Rsa公钥 {
	return New公钥(&r.pri.PublicKey)
}

// Rsa公钥 wraps RSA public cryptographic components with Chinese-named methods
// Provides encryption and signature verification operations
//
// Rsa公钥 封装 RSA 公钥组件，提供中文命名的方法
// 提供加密和验签操作
type Rsa公钥 struct {
	pub *rsa.PublicKey
}

// New公钥 creates a new Rsa公钥 instance from rsa.PublicKey
// Returns the wrapped public cryptographic components instance
//
// New公钥 从 rsa.PublicKey 创建新的 Rsa公钥 实例
// 返回封装后的公钥组件实例
func New公钥(puk *rsa.PublicKey) *Rsa公钥 {
	return &Rsa公钥{pub: puk}
}

// M加密 encrypts plaintext using PKCS#1 v1.5 encryption scheme
// Takes plaintext bytes as input and encrypts them with public components
// Returns encrypted ciphertext bytes on success, otherwise returns an exception
//
// M加密 使用 PKCS#1 v1.5 加密方案加密明文
// 接收明文字节作为输入并使用公钥组件加密
// 成功时返回加密后的密文字节，否则返回异常
func (r *Rsa公钥) M加密(v明文 []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, r.pub, v明文)
}

// M验签 verifies signature using SHA256 hash and PKCS#1 v1.5 verification scheme
// Takes plaintext and signature bytes as inputs and validates authenticity
// Returns nothing when signature is authentic, otherwise returns an exception
//
// M验签 使用 SHA256 哈希和 PKCS#1 v1.5 验签方案验证签名
// 接收明文和签名字节作为输入并验证真实性
// 签名有效时返回空值，否则返回异常
func (r *Rsa公钥) M验签(v明文 []byte, v签名 []byte) error {
	hash := sha256.New()
	hash.Write(v明文)
	return rsa.VerifyPKCS1v15(r.pub, crypto.SHA256, hash.Sum(nil), v签名)
}

// B导出 exports public cryptographic components as PKIX format bytes
// Serializes the public components to standard PKIX encoding
// Returns serialized bytes on success, otherwise returns an exception
//
// B导出 将公钥组件导出为 PKIX 格式字节
// 将公钥组件序列化为标准 PKIX 编码
// 成功时返回序列化字节，否则返回异常
func (r *Rsa公钥) B导出() ([]byte, error) {
	// Encode public cryptographic components to PKIX format byte slice
	// 将公钥组件编码为 PKIX 格式的字节切片
	pubBytes, err := x509.MarshalPKIXPublicKey(r.pub)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return pubBytes, nil
}
