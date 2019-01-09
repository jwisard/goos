// +build integration

package goos

import (
	"fmt"
	"testing"

	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OSClientTestSuite struct {
	suite.Suite
}

func TestOSClientSuite(t *testing.T) {
	suite.Run(t, new(OSClientTestSuite))
}

// This initializes the test suite. In this case, we're only going to ensure logging is turned
// completely off.
func (s *OSClientTestSuite) SetupSuite() {
	logger.SetLevel(logger.FatalLevel)
}

func (s *OSClientTestSuite) TestCreateOSClient() {

	assert := assert.New(s.T())

	// The AuthConfig created here is assumed to define the properties needed for a successful admin account
	// login.  If that's not the case, the tests below could fail.
	validCfg := getAdminAuthConfig()

	_, err := CreateOSClient(validCfg)

	if !assert.Nil(err, "Valid auth configuration should have produced no errors with call to CreateOSClient but did") {
		fmt.Println(err.Error())
	}

	invalidCfg := validCfg

	// change the password to something offensive
	invalidCfg.Password = "boogers!"

	_, err = CreateOSClient(invalidCfg)

	assert.NotNil(err, "Invalid auth configuration should have produced an error with call to CreateOSClient but didn't: "+err.Error())
}
