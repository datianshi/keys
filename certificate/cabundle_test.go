package certificate_test

import (
	. "github.com/datianshi/keys/certificate"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cabundle", func() {
	Context("Read from the bundle files", func() {
		bundle, testErr := LoadCerts("fixtures/cabundle.crt")
		It("Should read files without error", func() {
			Ω(testErr).ShouldNot(HaveOccurred())
		})
		It("Have two certs", func() {
			Ω(len(*bundle)).Should(Equal(2))
		})
		It("Should Have a valid chain", func() {
			Ω(bundle.VerifyChain()).Should(BeTrue())
		})
	})
	Context("Read from the bad bundle files", func() {
		bundle, testErr := LoadCerts("fixtures/bad_cabundle.crt")
		It("Should read files without error", func() {
			Ω(testErr).ShouldNot(HaveOccurred())
		})
		It("Should Not Have a valid chain", func() {
			Ω(bundle.VerifyChain()).Should(BeFalse())
		})
	})

})
