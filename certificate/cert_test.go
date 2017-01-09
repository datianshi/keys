package certificate_test

import (
	"github.com/datianshi/keys/certificate"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
	var test_cert = `-----BEGIN CERTIFICATE-----
MIID2jCCAsKgAwIBAgIJALZoanV5O6duMA0GCSqGSIb3DQEBBQUAMHAxCzAJBgNV
BAYTAlVTMQswCQYDVQQIDAJWQTERMA8GA1UEBwwIU29tZUNpdHkxEjAQBgNVBAoM
CU15Q29tcGFueTETMBEGA1UECwwKTXlEaXZpc2lvbjEYMBYGA1UEAwwPd3d3LmNv
bXBhbnkuY29tMB4XDTE3MDEwNjE2NTYwN1oXDTI3MDEwNDE2NTYwN1owcDELMAkG
A1UEBhMCVVMxCzAJBgNVBAgMAlZBMREwDwYDVQQHDAhTb21lQ2l0eTESMBAGA1UE
CgwJTXlDb21wYW55MRMwEQYDVQQLDApNeURpdmlzaW9uMRgwFgYDVQQDDA93d3cu
Y29tcGFueS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDDB9X2
PZrPA5hAXFxc+MiW/LBB8MWdGhhx2cyL5De+84tBKvhmXQ+FMy7dLNEQXMF2+Ep4
/bG/kpEQCrU8SJoqfG9laSg8JHocGSZimzSYAwmgyKPMrawgo/R791ezXzFAaCgD
3b/IU3ZejPoETPFqcj1qsj9Cm0+qm3IrE4KvTGz18DY45MWROzPbAYe6Mgfu7opY
7smuOWLe2zvY12mzZPSYF9gNWu6rezuIj9+1jnIxI9f9DgHE5i7CeRVv+AE5wUob
EcLT3+xk9ndLlCG7japxZQnM6DDm0lSvhUXHBY4cr2WmQ1z/ZdrkeukJfryKTPYW
lknp66+/NJlj5N2JAgMBAAGjdzB1MAsGA1UdDwQEAwIEMDATBgNVHSUEDDAKBggr
BgEFBQcDATBRBgNVHREESjBIgg93d3cuY29tcGFueS5jb22CC2NvbXBhbnkuY29t
gg93d3cuY29tcGFueS5uZXSCC2NvbXBhbnkubmV0hwQKCgoKhwQKCgoLMA0GCSqG
SIb3DQEBBQUAA4IBAQAapctUAF+5H+EyxtSos8O6fZB8Ilea555FNio+Yank517A
i9wRyP+3s3uPKmf7R317W9xTnW6jKIEGDAXxinjQIIqMNdOJA+D7HxtISh/mSTLR
ItMAkLopdSSyl6I4jcPZtv9XxQPQ4e+XMIJK/9viZfi9xI63+mzFDNr1crrzugIC
wYDOycT7BxPsoeqd6BNc0pzktaqgeFXZKQNyn2JO1E+boM1YI16VDGXYgfX9Tout
adpuh9hNT0jD6JouRL+n6Bg+FyyN+FzOYyvK6HY9o0uXmrQfeF9Al2sL35hbTPIh
JooCK18FVuTz1q/BjfCJEWRxcVFO4tn6JdpElIJB
-----END CERTIFICATE-----`

	var _ = `
-----BEGIN CERTIFICATE REQUEST-----
MIIDPTCCAiUCAQAwcDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAlZBMREwDwYDVQQH
DAhTb21lQ2l0eTESMBAGA1UECgwJTXlDb21wYW55MRMwEQYDVQQLDApNeURpdmlz
aW9uMRgwFgYDVQQDDA93d3cuY29tcGFueS5jb20wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDDB9X2PZrPA5hAXFxc+MiW/LBB8MWdGhhx2cyL5De+84tB
KvhmXQ+FMy7dLNEQXMF2+Ep4/bG/kpEQCrU8SJoqfG9laSg8JHocGSZimzSYAwmg
yKPMrawgo/R791ezXzFAaCgD3b/IU3ZejPoETPFqcj1qsj9Cm0+qm3IrE4KvTGz1
8DY45MWROzPbAYe6Mgfu7opY7smuOWLe2zvY12mzZPSYF9gNWu6rezuIj9+1jnIx
I9f9DgHE5i7CeRVv+AE5wUobEcLT3+xk9ndLlCG7japxZQnM6DDm0lSvhUXHBY4c
r2WmQ1z/ZdrkeukJfryKTPYWlknp66+/NJlj5N2JAgMBAAGggYcwgYQGCSqGSIb3
DQEJDjF3MHUwCwYDVR0PBAQDAgQwMBMGA1UdJQQMMAoGCCsGAQUFBwMBMFEGA1Ud
EQRKMEiCD3d3dy5jb21wYW55LmNvbYILY29tcGFueS5jb22CD3d3dy5jb21wYW55
Lm5ldIILY29tcGFueS5uZXSHBAoKCgqHBAoKCgswDQYJKoZIhvcNAQELBQADggEB
ADf2hwuyAEPvrf6lt+k4M7EEeCVoWPJ+ph5PVNqhgHfq6MivMxJxcF8D97BJ2V1b
Qi8kot/+vdpQeG+50PpDyS+UHHAybtYjRJn+OOOM8Yc8HScaUv/KCAoYTty7Lfkq
2+IN7RmB56xz/iSy6HT5J3ciuh0eD6t0L+FPwBnR4o+lcps4cvwQ2RAK/GKelCtN
AaUwGofAYpkCV9OYhiITXLZZGrvLrer/1DEqsZVTN4e323ou4geWXFphwD8HuIqg
L+8TQicpkj/ME+MXbNc9lriyRAUVLkQ7LH2fs6CtuR0JCjrp+Qq1OnfOjqL5n6rX
/ABvIFVvkSA6MSmkw+FGiOM=
-----END CERTIFICATE REQUEST-----
`
	var _ = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAwwfV9j2azwOYQFxcXPjIlvywQfDFnRoYcdnMi+Q3vvOLQSr4
Zl0PhTMu3SzREFzBdvhKeP2xv5KREAq1PEiaKnxvZWkoPCR6HBkmYps0mAMJoMij
zK2sIKP0e/dXs18xQGgoA92/yFN2Xoz6BEzxanI9arI/QptPqptyKxOCr0xs9fA2
OOTFkTsz2wGHujIH7u6KWO7Jrjli3ts72Ndps2T0mBfYDVruq3s7iI/ftY5yMSPX
/Q4BxOYuwnkVb/gBOcFKGxHC09/sZPZ3S5Qhu42qcWUJzOgw5tJUr4VFxwWOHK9l
pkNc/2Xa5HrpCX68ikz2FpZJ6euvvzSZY+TdiQIDAQABAoIBAQC15qgL67do3zxw
wvQujqdgE7w9NFjLx1wHTBF4ZFoOEk7OcEqQBroEDRHnnfY8IZcb5EZphIz1VblQ
0fvphvmrKLYpSqQJPChb7qUbkRDDuwmO2bPUoiMJqoHL4Yz9sopYPQ0dVgSaziV3
2FsUCcOc+YII8skH0Z3DTQp13NOltKx3YYGV7BoHYSVA3CP2gRa3rONa45wTmq5f
2FdPTVhm571EdLrXeeDivLk9cErLa9Qek3T7inLMgTZzM9Tr6ZrdH1Ik2CXyvLXv
ROVbTTc4b9OJu/Y7iH3XFlSqT6RG4MK5c4NeQEhgU5prr1c77VBaDK7tLfXWHKBY
IJXuh+1RAoGBAOUIW6pQZ8TSqOuPFFShyL+WjazDOVBYPOo2Tn3FH09fvI1q2u/s
BXrTn3usFJpwPY4n83A0IsqLb/qpdFcfLFeZ5g95+/Sjk52YcCi9TMRfdYujVHCX
76zFBRFudQ6b5dX+gn358QD72MYYjkrUHkm7X8BmfCMzeq5F4bOK5H5NAoGBANn+
kUWDGmbZZ3KZ2GGPYAjFMBlHfttJhOZpGtdZiXCY4x2nMlBPLro7DUqCSJWyYYHN
KHdvmLh0MSQrIp7/bi0HSBxz+CVZN7ZADzGNVYOkCsTGPDBSPzxJSXAZGl+bYOFD
9KNZTCmceyg1F/lVOapGRYvJydoaWPjpA26NXlItAoGBAN12joSKiTy14qccTDFl
2yL/JnF2uNXq7MQgzRcqZBMxZLZuK9QxKj9RoHZKuAV6INtz4B6QG1A1Mot0Dc4s
k4kFhINUlgBBp/yMqW+LVE7Sa2wh/AV/tT8fp73qTABbWQvqtS2xYIxt+nfayMSs
hbkQGQ+HSeQT5D3dztf4pudFAoGAP80m3W5iksl8YmkukFczJz33Nt7a/PlGgV0v
sumMtTy59BrUOfcC2RShbjY1YkLEEQZnmpU54Hp5S3mod1qQ93LTIyi57sxu5tNW
3wzbg1zGnhLa2NWplk5oxPtD4cmLEEMgXZKsHN2kmf74FkxgRTqt7hzQDBp8AdXI
4FlVyr0CgYBpymKwlQry4XSnSWGbY7LU1fvRLtgHTkf/9WOW9Ljjs9QFs/pNoHut
fWe+IBbrHPFgO4G2m2TDXoy6ZyUQ5M259W6VzESFf3yKQNCQ8+u0kmBU6mH0GpmA
T2ZS4mwmSan8SmpGuOR5hnhcHlaVFn9vOPugxtCz3S64XI0rxX9K9w==
-----END RSA PRIVATE KEY-----
`

	Context("Load Correct Root Certificate", func() {
		var cert *certificate.Cert
		cert, err := certificate.LoadCertificate(test_cert)
		It("Should Not return error when certficate is valid", func() {
			Ω(err).Should(BeNil())
		})
		It("Should Return 4 HostNames", func() {
			hostName := cert.GetHostNames()
			Ω(len(hostName)).Should(Equal(4))
		})
		It("Should Return subject", func() {
			cn := cert.GetCommonName()
			Ω(cn).Should(Equal("www.company.com"))
		})
		It("Should be the root certificate", func() {
			isRoot, err := cert.IsRootCert()
			Ω(err).Should(BeNil())
			Ω(isRoot).Should(Equal(true))
		})

	})

	Context("Load InCorrect Certificate", func() {
		_, err := certificate.LoadCertificate("Invalid Cert")
		It("Should return error when certficate is invalid", func() {
			Ω(err).ShouldNot(BeNil())
		})
	})

})
