package certificate_test

import (
	"github.com/datianshi/keys/certificate"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ReadFile(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(b), err
}

var _ = Describe("Cert", func() {
	/*
	   Certificate Request:
	       Data:
	           Version: 0 (0x0)
	           Subject: C=US, ST=VA, L=SomeCity, O=MyCompany, OU=MyDivision, CN=www.company.com
	           Subject Public Key Info:
	               Public Key Algorithm: rsaEncryption
	               RSA Public Key: (2048 bit)
	                   Modulus (2048 bit):
	                       00:c3:07:d5:f6:3d:9a:cf:03:98:40:5c:5c:5c:f8:
	                       c8:96:fc:b0:41:f0:c5:9d:1a:18:71:d9:cc:8b:e4:
	                       37:be:f3:8b:41:2a:f8:66:5d:0f:85:33:2e:dd:2c:
	                       d1:10:5c:c1:76:f8:4a:78:fd:b1:bf:92:91:10:0a:
	                       b5:3c:48:9a:2a:7c:6f:65:69:28:3c:24:7a:1c:19:
	                       26:62:9b:34:98:03:09:a0:c8:a3:cc:ad:ac:20:a3:
	                       f4:7b:f7:57:b3:5f:31:40:68:28:03:dd:bf:c8:53:
	                       76:5e:8c:fa:04:4c:f1:6a:72:3d:6a:b2:3f:42:9b:
	                       4f:aa:9b:72:2b:13:82:af:4c:6c:f5:f0:36:38:e4:
	                       c5:91:3b:33:db:01:87:ba:32:07:ee:ee:8a:58:ee:
	                       c9:ae:39:62:de:db:3b:d8:d7:69:b3:64:f4:98:17:
	                       d8:0d:5a:ee:ab:7b:3b:88:8f:df:b5:8e:72:31:23:
	                       d7:fd:0e:01:c4:e6:2e:c2:79:15:6f:f8:01:39:c1:
	                       4a:1b:11:c2:d3:df:ec:64:f6:77:4b:94:21:bb:8d:
	                       aa:71:65:09:cc:e8:30:e6:d2:54:af:85:45:c7:05:
	                       8e:1c:af:65:a6:43:5c:ff:65:da:e4:7a:e9:09:7e:
	                       bc:8a:4c:f6:16:96:49:e9:eb:af:bf:34:99:63:e4:
	                       dd:89
	                   Exponent: 65537 (0x10001)
	           Attributes:
	           Requested Extensions:
	               X509v3 Key Usage:
	                   Key Encipherment, Data Encipherment
	               X509v3 Extended Key Usage:
	                   TLS Web Server Authentication
	               X509v3 Subject Alternative Name:
	                   DNS:www.company.com, DNS:company.com, DNS:www.company.net, DNS:company.net, IP Address:10.10.10.10, IP Address:10.10.10.11
	       Signature Algorithm: sha256WithRSAEncryption
	           37:f6:87:0b:b2:00:43:ef:ad:fe:a5:b7:e9:38:33:b1:04:78:
	           25:68:58:f2:7e:a6:1e:4f:54:da:a1:80:77:ea:e8:c8:af:33:
	           12:71:70:5f:03:f7:b0:49:d9:5d:5b:42:2f:24:a2:df:fe:bd:
	           da:50:78:6f:b9:d0:fa:43:c9:2f:94:1c:70:32:6e:d6:23:44:
	           99:fe:38:e3:8c:f1:87:3c:1d:27:1a:52:ff:ca:08:0a:18:4e:
	           dc:bb:2d:f9:2a:db:e2:0d:ed:19:81:e7:ac:73:fe:24:b2:e8:
	           74:f9:27:77:22:ba:1d:1e:0f:ab:74:2f:e1:4f:c0:19:d1:e2:
	           8f:a5:72:9b:38:72:fc:10:d9:10:0a:fc:62:9e:94:2b:4d:01:
	           a5:30:1a:87:c0:62:99:02:57:d3:98:86:22:13:5c:b6:59:1a:
	           bb:cb:ad:ea:ff:d4:31:2a:b1:95:53:37:87:b7:db:7a:2e:e2:
	           07:96:5c:5a:61:c0:3f:07:b8:8a:a0:2f:ef:13:42:27:29:92:
	           3f:cc:13:e3:17:6c:d7:3d:96:b8:b2:44:05:15:2e:44:3b:2c:
	           7d:9f:b3:a0:ad:b9:1d:09:0a:3a:e9:f9:0a:b5:3a:77:ce:8e:
	           a2:f9:9f:aa:d7:fc:00:6f:20:55:6f:91:20:3a:31:29:a4:c3:
	           e1:46:88:e3
	*/
	root_cert, _ := ReadFile("fixtures/root.crt")
	intermidiate_cert, _ := ReadFile("fixtures/intermidiate.crt")
	user_cert, _ := ReadFile("fixtures/user.crt")
	var root *certificate.Cert
	var user *certificate.Cert
	var intermidiate *certificate.Cert

	Context("Load Correct Root Certificate", func() {
		var err error
		root, err = certificate.LoadCertificate(root_cert)
		It("Should Not return error when certficate is valid", func() {
			Ω(err).Should(BeNil())
		})
		It("Should Return subject", func() {
			cn := root.GetCommonName()
			Ω(cn).Should(Equal("www.company.com"))
		})
		It("Should be the root certificate", func() {
			isRoot, err := root.IsRootCert()
			Ω(err).Should(BeNil())
			Ω(isRoot).Should(Equal(true))
		})

	})

	Context("Load Correct Intermidiate Certificate", func() {
		var err error
		intermidiate, err = certificate.LoadCertificate(intermidiate_cert)
		It("Should Not return error when certficate is valid", func() {
			Ω(err).Should(BeNil())
		})
		It("Should be signed by the root certificate", func() {
			Ω(intermidiate.CheckSignedFrom(root)).Should(Not(HaveOccurred()))
		})
		It("Should be the CA Cert", func() {
			isCA := intermidiate.IsCACert()
			Ω(isCA).Should(Equal(true))
		})
	})

	Context("Load Correct User Certificate", func() {
		var err error
		user, err = certificate.LoadCertificate(user_cert)
		It("Should Not return error when certficate is valid", func() {
			Ω(err).Should(BeNil())
		})
		It("Should Not be the CA Cert", func() {
			isCA := user.IsCACert()
			Ω(isCA).Should(Equal(false))
		})
		It("Should Be signed from the intermidiate", func() {
			Ω(user.CheckSignedFrom(intermidiate)).Should(Not(HaveOccurred()))
		})
		It("Should Not be signed from the root", func() {
			Ω(user.CheckSignedFrom(root)).Should(HaveOccurred())
		})
		It("Should Return 4 HostNames", func() {
			hostName := user.GetHostNames()
			Ω(len(hostName)).Should(Equal(4))
		})

	})

	Context("Load InCorrect Certificate", func() {
		_, err := certificate.LoadCertificate("Invalid Cert")
		It("Should return error when certficate is invalid", func() {
			Ω(err).ShouldNot(BeNil())
		})
	})

})
