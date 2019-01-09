// +build integration

package goos

import (
	"fmt"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FlavorsTestSuite struct {
	suite.Suite
	adminOsClient OSClient
	seedFlavors   []flavors.Flavor
}

func TestFlavorsSuite(t *testing.T) {
	suite.Run(t, new(FlavorsTestSuite))
}

// This initializes the test suite.  e.g. we need a valid, authorized OSClient for each test,
// so this method creates that once for re-use in subsequent tests.
func (s *FlavorsTestSuite) SetupSuite() {

	assert := assert.New(s.T())

	// The AuthConfig created here is assumed to define the properties needed for a successful admin account
	// login.  If that's not the case, the tests below could fail.
	validCfg := getAdminAuthConfig()
	osClient, err := CreateOSClient(validCfg)
	if !assert.Nil(err, "Valid auth configuration should have produced no errors with call to CreateOSClient but did") {
		fmt.Println(err.Error())
	}

	s.adminOsClient = osClient

	// Since we need the flavor slice for most of the other tests, we'll go ahead
	// and test the RetrieveFlavors method here and cache the seedFlavors.
	// That will save some test execution time.
	// BTW, this may not be good testing behavior. :P
	flavors, err := s.adminOsClient.RetrieveFlavors()

	assert.Nil(err, "RetrieveFlavors should not have returned an exception but did")

	assert.True(len(flavors) > 1, "Flavor slice from RetrieveFlavors should have had more than one Flavor but didn't")

	s.seedFlavors = flavors
}

func (s *FlavorsTestSuite) TestRetrieveFlavorByID() {

	assert := assert.New(s.T())

	// Grab a real flavor
	seedFlavor := s.seedFlavors[0]

	flavorByID, err := s.adminOsClient.RetrieveFlavorByID(seedFlavor.ID)

	if !assert.Nil(err, "RetrieveFlavorByID should not have returned an error but did") {
		fmt.Println(err.Error())
	}

	assert.NotNil(flavorByID, "RetrieveFlavorByID returned a nil Flavor but it should not have done")

	assert.Equal(seedFlavor, *flavorByID, "RetrieveFlavorByID did not return a Flavor equal to the expected Flavor")

}

func (s *FlavorsTestSuite) TestRetrieveFlavorByName() {

	assert := assert.New(s.T())

	// Grab a real flavor
	seedFlavor := s.seedFlavors[0]

	flavorByName, err := s.adminOsClient.RetrieveFlavorByName(seedFlavor.Name)

	if !assert.Nil(err, "RetrieveFlavorByName should not have returned an error but did") {
		fmt.Println(err.Error())
	}

	assert.NotNil(flavorByName, "RetrieveFlavorByName returned a nil Flavor but it should not have done")

	assert.Equal(seedFlavor, *flavorByName, "RetrieveFlavorByName did not return a Flavor equal to the expected Flavor")
}
