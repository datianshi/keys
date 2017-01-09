package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

//LoadCertificate Parse a certificate to x509 Certificate project
func LoadCertificate(pemCert string) (*Cert, error) {
	p, _ := pem.Decode([]byte(pemCert))
	if p == nil {
		return nil, errors.New("The certificate is not in the correct pem format")
	}
	c, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		return nil, err
	}
	return &Cert{certificate: c}, nil
}

//Cert A wrapper of x509 certificate
type Cert struct {
	certificate *x509.Certificate
}

//GetHostNames
func (c *Cert) GetHostNames() []string {
	return c.certificate.DNSNames
}

//GetCommonName
func (c *Cert) GetCommonName() string {
	return c.certificate.Subject.CommonName
}

//IsRootCert By validating it was signed by itself
func (c *Cert) IsRootCert() (bool, error) {
	err := c.certificate.CheckSignature(c.certificate.SignatureAlgorithm, c.certificate.RawTBSCertificate, c.certificate.Signature)
	if err != nil {
		return false, err
	}
	return true, err
}
