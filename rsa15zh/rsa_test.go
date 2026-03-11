package rsa15zh

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/must"
)

var case私钥 string
var case公钥 string

// TestMain generates RSA cryptographic pairs before executing tests
// Creates 2048-bit private components and derives public components
//
// TestMain 在执行测试前生成 RSA 密钥对
// 创建 2048 位私钥组件并导出公钥组件
func TestMain(m *testing.M) {
	v私钥, err := R随机私钥(2048)
	must.Done(err)
	v公钥, err := R获得公钥(v私钥)
	must.Done(err)

	case私钥 = base64.StdEncoding.EncodeToString(v私钥)
	case公钥 = base64.StdEncoding.EncodeToString(v公钥)
	m.Run()
}

// TestF装载公钥 tests loading public components and encryption operations
// TestF装载公钥 测试加载公钥组件和加密操作
func TestF装载公钥(t *testing.T) {
	r公钥 := must装载公钥(t)
	v密文, err := r公钥.M加密([]byte("abc"))
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v密文))
}

// must装载公钥 loads public cryptographic components from base64 encoded text
// must装载公钥 从 base64 编码文本加载公钥组件
func must装载公钥(t *testing.T) *Rsa公钥 {
	v公钥, err := base64.StdEncoding.DecodeString(case公钥)
	require.NoError(t, err)
	r公钥, err := F装载公钥(v公钥)
	require.NoError(t, err)
	return r公钥
}

// TestF装载私钥 tests loading private components and signing operations
// TestF装载私钥 测试加载私钥组件和签名操作
func TestF装载私钥(t *testing.T) {
	r私钥 := must装载私钥(t)
	v密文, err := r私钥.M签名([]byte("xyz"))
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v密文))
}

// must装载私钥 loads private cryptographic components from base64 encoded text
// must装载私钥 从 base64 编码文本加载私钥组件
func must装载私钥(t *testing.T) *Rsa私钥 {
	v私钥, err := base64.StdEncoding.DecodeString(case私钥)
	require.NoError(t, err)
	r私钥, err := F装载私钥(v私钥)
	require.NoError(t, err)
	return r私钥
}

// TestRsa公钥_M加密 tests encryption and decryption workflow
// TestRsa公钥_M加密 测试加密和解密工作流
func TestRsa公钥_M加密(t *testing.T) {
	r公钥 := must装载公钥(t)
	v密文, err := r公钥.M加密([]byte("abc"))
	require.NoError(t, err)

	r私钥 := must装载私钥(t)
	v明文, err := r私钥.M解密(v密文)
	require.NoError(t, err)

	require.Equal(t, "abc", string(v明文))
}

// TestRsa私钥_M签名 tests signing and verification workflow
// TestRsa私钥_M签名 测试签名和验签工作流
func TestRsa私钥_M签名(t *testing.T) {
	r私钥 := must装载私钥(t)
	v密文, err := r私钥.M签名([]byte("xyz"))
	require.NoError(t, err)

	r公钥 := must装载公钥(t)
	err = r公钥.M验签([]byte("xyz"), v密文)
	require.NoError(t, err)
}

// TestRsa私钥_P公钥 tests extracting public components from private components
// TestRsa私钥_P公钥 测试从私钥组件提取公钥组件
func TestRsa私钥_P公钥(t *testing.T) {
	r私钥 := must装载私钥(t)
	v密文, err := r私钥.M签名([]byte("xyz"))
	require.NoError(t, err)

	r公钥 := r私钥.P公钥()
	err = r公钥.M验签([]byte("xyz"), v密文)
	require.NoError(t, err)
}

// TestRsa私钥_B导出 tests exporting private components to bytes
// TestRsa私钥_B导出 测试将私钥组件导出为字节
func TestRsa私钥_B导出(t *testing.T) {
	r私钥 := must装载私钥(t)
	data, err := r私钥.B导出()
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(data))
}

// TestRsa公钥_B导出 tests exporting public components to bytes
// TestRsa公钥_B导出 测试将公钥组件导出为字节
func TestRsa公钥_B导出(t *testing.T) {
	r公钥 := must装载公钥(t)
	data, err := r公钥.B导出()
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(data))
}
