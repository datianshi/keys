package main

import (
	"fmt"
	"os"

	"github.com/datianshi/keys/certificate"
	. "github.com/datianshi/keys/cli"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	var (
		VERSION = ""
	)
	app := &cli.App{
		Name:     "keys",
		HelpName: "keys",
		Version:  VERSION,
	}
	app.Commands = []*cli.Command{
		{
			Name:  "cert",
			Usage: "cert",
			Subcommands: []*cli.Command{
				{
					Name:  "sans",
					Usage: "subject alternative names",
					Action: CertFunc(func(cert *certificate.Cert) error {
						fmt.Println(cert.GetHostNames())
						return nil
					}),
					Flags: []cli.Flag{CertFile},
				},
				{
					Name:  "cn",
					Usage: "common name",
					Action: CertFunc(func(cert *certificate.Cert) error {
						fmt.Println(cert.GetCommonName())
						return nil
					}),
					Flags: []cli.Flag{CertFile},
				},
			},
		},
	}
	app.Run(os.Args)
}
