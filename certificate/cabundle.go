package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

//CABundle a list of certificates
type CABundle []*Cert

//LoadCerts Load Certs from a file
func LoadCerts(filename string) (*CABundle, error) {
	bundle := make([]*Cert, 0)
	rest, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cert *Cert
	p, rest := pem.Decode(rest)
	for p != nil {
		c, err := x509.ParseCertificate(p.Bytes)
		if err != nil {
			return nil, err
		}
		cert = &Cert{certificate: c}
		p, rest = pem.Decode(rest)
		bundle = append(bundle, cert)
	}
	caBundle := CABundle(bundle)
	return &caBundle, nil
}

//AsSlice Return to slices
func (b CABundle) AsSlice() []*Cert {
	return []*Cert(b)
}

//VerifyChain Verify the chain is valid
func (b CABundle) VerifyChain() bool {
	length := len(b)
	if length <= 1 {
		return true
	}
	for i := range b {
		if (i + 1) != length {
			err := b[i].CheckSignedFrom(b[i+1])
			if err != nil {
				return false
			}
		}
	}
	return true
}
