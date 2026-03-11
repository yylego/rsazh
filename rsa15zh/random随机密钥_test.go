package rsa15zh

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestR随机私钥 tests generating random RSA private cryptographic components
// TestR随机私钥 测试生成随机 RSA 私钥组件
func TestR随机私钥(t *testing.T) {
	v私钥, err := R随机私钥(2048)
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v私钥))
}

// TestR获得公钥 tests extracting public components from private components
// TestR获得公钥 测试从私钥组件提取公钥组件
func TestR获得公钥(t *testing.T) {
	v私钥, err := R随机私钥(4096)
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v私钥))

	v公钥, err := R获得公钥(v私钥)
	require.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(v公钥))
}
