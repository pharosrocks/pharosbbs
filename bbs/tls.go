package bbs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"os"
	"time"
)

func GetCertificate() tls.Certificate {
	now := time.Now()
	template := &x509.Certificate{
		SerialNumber: big.NewInt(now.Unix()),
		Subject: pkix.Name{
			CommonName:         os.Getenv("BBS_FQDN"),
			Country:            []string{"Taiwan"},
			Organization:       []string{"Organization"},
			OrganizationalUnit: []string{"OrganizationalUnit"},
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(10, 0, 0),
		SubjectKeyId:          []byte(os.Getenv("BBS_FQDN")),
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage: x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	cert, _ := x509.CreateCertificate(rand.Reader, template, template, priv.Public(), priv)

	outCert := tls.Certificate{}
	outCert.Certificate = append(outCert.Certificate, cert)
	outCert.PrivateKey = priv

	return outCert
}
