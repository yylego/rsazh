package rsa15zh

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestR随机PKCS8私钥 tests generating RSA private components in PKCS#8.
// TestR随机PKCS8私钥 测试生成 PKCS#8 私钥。
func TestR随机PKCS8私钥(t *testing.T) {
	v私钥, err := R随机PKCS8私钥(2048)
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v私钥))
}

// TestR获得PKIX公钥 tests extracting PKIX public components from PKCS#8.
// TestR获得PKIX公钥 测试从 PKCS#8 私钥提取 PKIX 公钥。
func TestR获得PKIX公钥(t *testing.T) {
	v私钥, err := R随机PKCS8私钥(4096)
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v私钥))

	v公钥, err := R获得PKIX公钥(v私钥)
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v公钥))
}
