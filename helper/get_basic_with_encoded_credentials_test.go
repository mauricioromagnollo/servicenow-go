package helper

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetBasicWithEncodedCredentials", func() {
	It("should return an error if username is empty", func() {
		_, err := GetBasicWithEncodedCredentials("", "password")
		Expect(err).To(MatchError("The params 'username' and 'password' are required and cannot be empty"))
	})

	It("should return an error if password is empty", func() {
		_, err := GetBasicWithEncodedCredentials("username", "")
		Expect(err).To(MatchError("The params 'username' and 'password' are required and cannot be empty"))
	})

	It("should return the basic authentication header with the encoded credentials", func() {
		username := "username"
		password := "password"

		expectedBasicAuth := "Basic dXNlcm5hbWU6cGFzc3dvcmQ="

		basicAuth, err := GetBasicWithEncodedCredentials(username, password)

		Expect(err).To(BeNil())
		Expect(basicAuth).To(Equal(expectedBasicAuth))
	})
})
