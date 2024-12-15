package cert

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net/http"
)

func Load() (*x509.Certificate, *rsa.PrivateKey, error) {
	block, _ := pem.Decode(Public)
	if block == nil {
		return nil, nil, http.ErrAbortHandler
	}
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, err
	}

	block, _ = pem.Decode(Private)
	if block == nil {
		return nil, nil, http.ErrAbortHandler
	}
	caKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, nil, err
	}
	return caCert, caKey, nil
}
