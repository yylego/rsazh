package rsa15zh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	"github.com/yylego/erero"
)

// R随机PKCS8私钥 generates RSA private components as PKCS#8 DER bytes.
// Takes bit count as input (common sizes: 2048 and 4096 bits)
// Returns cryptographic components as PKCS#8 format bytes on success, otherwise returns an exception
//
// R随机PKCS8私钥 生成指定位数的新 RSA 私钥，返回 PKCS#8 DER 字节。
// 接收位数作为输入（常用大小：2048 和 4096 位）
// 成功时返回 PKCS#8 格式的密钥组件字节，否则返回异常
func R随机PKCS8私钥(n位数 int) ([]byte, error) {
	// Generate RSA private key with specified bit size (e.g., 2048 bits)
	// 通过指定 RSA 密钥的长度，例如 2048 位，生成 RSA 私钥
	pri, err := rsa.GenerateKey(rand.Reader, n位数)
	if err != nil {
		return nil, erero.Wro(err)
	}

	// Encode private key to PKCS#8 format byte slice
	// 将私钥编码为 PKCS#8 格式的字节切片
	priBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return nil, erero.Wro(err)
	}

	return priBytes, nil
}

// R获得PKIX公钥 extracts PKIX DER public components from PKCS#8 DER private bytes.
// Takes PKCS#8 format private components bytes as input and derives public components
// Returns public components as PKIX format bytes on success, otherwise returns an exception
//
// R获得PKIX公钥 从 PKCS#8 DER 私钥字节提取 PKIX DER 公钥字节。
// 接收 PKCS#8 格式的私钥组件字节作为输入并推导公钥组件
// 成功时返回 PKIX 格式的公钥组件字节，否则返回异常
func R获得PKIX公钥(privateKeyBytes []byte) ([]byte, error) {
	// Parse PKCS#8 format private key byte slice
	// 解析 PKCS#8 格式的私钥字节切片
	prk, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, erero.Wro(err)
	}

	// Convert private key to *rsa.PrivateKey type
	// 将私钥转换为 *rsa.PrivateKey 类型
	pri, ok := prk.(*rsa.PrivateKey)
	if !ok {
		return nil, erero.New("转换失败")
	}

	// Extract RSA public key
	// 提取 RSA 公钥
	pub := pri.Public()

	// Encode public key to PKIX format byte slice
	// 将公钥编码为 PKIX 格式的字节切片
	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, erero.Wro(err)
	}

	return pubBytes, nil
}

// F装载PKCS8私钥 loads private cryptographic components from PKCS#8 DER bytes.
// Takes PKCS#8 encoded bytes as input, parses and validates them
// Returns wrapped Rsa私钥 instance on success, otherwise returns an exception
//
// F装载PKCS8私钥 从 PKCS#8 DER 字节加载私钥组件。
// 接收 PKCS#8 编码字节作为输入，解析并验证
// 成功时返回封装的 Rsa私钥 实例，否则返回异常
func F装载PKCS8私钥(v私钥 []byte) (*Rsa私钥, error) {
	prk, err := x509.ParsePKCS8PrivateKey(v私钥)
	if err != nil {
		return nil, erero.Wro(err)
	}
	pri, ok := prk.(*rsa.PrivateKey)
	if !ok {
		return nil, erero.New("转换失败")
	}
	return &Rsa私钥{pri: pri}, nil
}

// F装载PKIX公钥 loads public cryptographic components from PKIX DER bytes.
// Takes PKIX encoded bytes as input, parses and validates them
// Returns wrapped Rsa公钥 instance on success, otherwise returns an exception
//
// F装载PKIX公钥 从 PKIX DER 字节加载公钥组件。
// 接收 PKIX 编码字节作为输入，解析并验证
// 成功时返回封装的 Rsa公钥 实例，否则返回异常
func F装载PKIX公钥(v公钥 []byte) (*Rsa公钥, error) {
	puk, err := x509.ParsePKIXPublicKey(v公钥)
	if err != nil {
		return nil, erero.Wro(err)
	}
	pub, ok := puk.(*rsa.PublicKey)
	if !ok {
		return nil, erero.New("转换失败")
	}
	return &Rsa公钥{pub: pub}, nil
}
