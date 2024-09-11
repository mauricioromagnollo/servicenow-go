package api

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("getBasicWithEncodedCredentials", func() {
	It("should return an error if username is empty", func() {
		_, err := getBasicWithEncodedCredentials("", "password")
		Expect(err).To(MatchError("The params 'username' and 'password' are required and cannot be empty"))
	})

	It("should return an error if password is empty", func() {
		_, err := getBasicWithEncodedCredentials("username", "")
		Expect(err).To(MatchError("The params 'username' and 'password' are required and cannot be emptys"))
	})
})
