package cli

import (
	"io/ioutil"

	"github.com/datianshi/keys/certificate"
	"gopkg.in/urfave/cli.v2"
)

//CertInfo Action on certificate object
type CertInfo func(cert *certificate.Cert) error

//CertFunc Generate a Cli action
func CertFunc(ci CertInfo) cli.ActionFunc {
	return func(c *cli.Context) error {
		certContent, err := ioutil.ReadFile(c.String("cert-file"))
		if err != nil {
			return err
		}
		cert, err := certificate.LoadCertificate(string(certContent))
		if err != nil {
			return err
		}
		return ci(cert)
	}
}

//CertFile
var CertFile = &cli.StringFlag{
	Name:    "cert-file",
	Aliases: []string{"c"},
	Usage:   "Cert File",
	EnvVars: []string{"CERT_FILE"},
}
