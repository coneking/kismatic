package integration_tests

import (
	"os"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("kismatic install validate tests", func() {
	BeforeEach(func() {
		dir := setupTestWorkingDir()
		os.Chdir(dir)
	})

	Describe("Running validation with package installation disabled", func() {
		Context("using a Minikube Ubuntu layout", func() {
			ItOnAWS("should succeed if and only if all packages are installed on the node", func(aws infrastructureProvisioner) {
				validateMiniPkgInstallationDisabled(aws, Ubuntu1604LTS)
			})
		})

		Context("using a Minikube CentOS 7 layout", func() {
			ItOnAWS("should succeed if and only if all packages are installed on the node", func(aws infrastructureProvisioner) {
				validateMiniPkgInstallationDisabled(aws, CentOS7)
			})
		})

		Context("using a Minikube RedHat 7 layout", func() {
			ItOnAWS("should succeed if and only if all packages are installed on the node", func(aws infrastructureProvisioner) {
				validateMiniPkgInstallationDisabled(aws, RedHat7)
			})
		})
	})

	Describe("Running validation with bad SSH key", func() {
		Context("Using Ubuntu 16.04", func() {
			ItOnAWS("should result in an ssh validation error", func(aws infrastructureProvisioner) {
				WithMiniInfrastructure(Ubuntu1604LTS, aws, func(node NodeDeets, sshKey string) {
					badSSHKey, err := getBadSSHKeyFile()
					if err != nil {
						Fail("Unexpected error generating fake SSH key: %v")
					}
					ValidateKismaticMiniWithBadSSH(node, node.SSHUser, badSSHKey)
				})
			})
		})
	})
})
