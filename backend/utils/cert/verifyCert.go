package cert

import (
	"crypto/x509"
	"encoding/pem"
	"net/http"
)

// 验证客户端上传的证书是否由本CA签发
func verifyCert(caCert *x509.Certificate, clientCertPem []byte) error {
	block, _ := pem.Decode(clientCertPem)
	if block == nil {
		return http.ErrAbortHandler
	}
	clientCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	roots := x509.NewCertPool()
	roots.AddCert(caCert)
	opts := x509.VerifyOptions{
		Roots: roots,
	}

	_, err = clientCert.Verify(opts)
	return err
}
