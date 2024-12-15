package cert

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

// GenerateCertAndKey 为客户端生成证书与私钥，并使用CA签发
func GenerateCertAndKey(caCert *x509.Certificate, caKey *rsa.PrivateKey) (clientCertPem, clientKeyPem []byte, err error) {
	// 生成客户端私钥
	clientKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// 准备客户端证书模板
	serialNumber, _ := rand.Int(rand.Reader, new(big.Int).SetInt64(999999999))
	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:   "ClientUser",
			Organization: []string{"ClientOrg"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:  x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},
	}

	// 使用CA证书为客户端证书签名
	clientCertDER, err := x509.CreateCertificate(rand.Reader, template, caCert, &clientKey.PublicKey, caKey)
	if err != nil {
		return nil, nil, err
	}

	// 编码客户端证书
	clientCertPemBuffer := new(bytes.Buffer)
	_ = pem.Encode(clientCertPemBuffer, &pem.Block{Type: "CERTIFICATE", Bytes: clientCertDER})
	clientCertPem = clientCertPemBuffer.Bytes()

	// 编码客户端私钥
	clientKeyPemBuffer := new(bytes.Buffer)
	_ = pem.Encode(clientKeyPemBuffer, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(clientKey)})
	clientKeyPem = clientKeyPemBuffer.Bytes()

	return clientCertPem, clientKeyPem, nil
}
