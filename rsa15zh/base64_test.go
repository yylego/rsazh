package rsa15zh_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/rsazh/rsa15zh"
)

func TestPKCS1Base64(t *testing.T) {
	pri, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	encoded := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(pri))
	ciphertext, err := rsa15zh.New公钥(&pri.PublicKey).M加密([]byte("hello 世界"))
	require.NoError(t, err)
	plain, err := rsa15zh.M解密PKCS1Base64(encoded, base64.StdEncoding.EncodeToString(ciphertext))
	require.NoError(t, err)
	require.Equal(t, "hello 世界", string(plain))
	_, err = rsa15zh.M解密PKCS1Base64("!", "!")
	require.Error(t, err)
	_, err = rsa15zh.M解密PKCS1Base64(encoded, "!")
	require.Error(t, err)
	_, err = rsa15zh.M解密PKCS1Base64(encoded, base64.StdEncoding.EncodeToString([]byte("bad")))
	require.Error(t, err)
	pkcs8, err := x509.MarshalPKCS8PrivateKey(pri)
	require.NoError(t, err)
	_, err = rsa15zh.F装载PKCS1私钥Base64(base64.StdEncoding.EncodeToString(pkcs8))
	require.Error(t, err)
}

func TestBase64往返(t *testing.T) {
	pri, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	v私钥 := rsa15zh.New私钥(pri)
	s私钥 := v私钥.B导出PKCS1Base64()
	require.Equal(t, base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(pri)), s私钥)
	v装载私钥, err := rsa15zh.F装载PKCS1私钥Base64(s私钥)
	require.NoError(t, err)

	a明文 := []byte("hello 世界")
	s密文, err := v私钥.P公钥().M加密Base64(a明文)
	require.NoError(t, err)
	a解密, err := v装载私钥.M解密Base64(s密文)
	require.NoError(t, err)
	require.Equal(t, a明文, a解密)

	// Verify that the existing byte API accepts the same ciphertext.
	// 确认已有字节接口也能处理相同密文。
	a密文, err := base64.StdEncoding.DecodeString(s密文)
	require.NoError(t, err)
	a解密, err = v私钥.M解密(a密文)
	require.NoError(t, err)
	require.Equal(t, a明文, a解密)

	a解密, err = rsa15zh.M解密PKCS1Base64(s私钥, s密文)
	require.NoError(t, err)
	require.Equal(t, a明文, a解密)
}
