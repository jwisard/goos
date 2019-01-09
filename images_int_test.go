// +build integration

package goos

import (
	"fmt"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ImagesTestSuite struct {
	suite.Suite
	adminOsClient OSClient
	seedImages    []images.Image
}

func TestImagesSuite(t *testing.T) {
	suite.Run(t, new(ImagesTestSuite))
}

// This initializes the test suite.  e.g. we need a valid, authorized OSClient for each test,
// so this method creates that once for re-use in subsequent tests.
func (s *ImagesTestSuite) SetupSuite() {

	assert := assert.New(s.T())

	// The AuthConfig created here is assumed to define the properties needed for a successful admin account
	// login.  If that's not the case, the tests below could fail.
	validCfg := getAdminAuthConfig()
	osClient, err := CreateOSClient(validCfg)
	if !assert.Nil(err, "Valid auth configuration should have produced no errors with call to CreateOSClient but did") {
		fmt.Println(err.Error())
	}

	s.adminOsClient = osClient

	// Since we need the image slice for most of the other tests, we'll go ahead
	// and test the RetrieveImages method here and cache the seedImages.
	// That will save some test execution time.
	// BTW, this may not be good testing behavior. :P
	Images, err := s.adminOsClient.RetrieveImages()

	assert.Nil(err, "RetrieveImages should not have returned an exception but did")

	assert.True(len(Images) > 1, "Image slice from RetrieveImages should have had more than one image but didn't")

	s.seedImages = Images
}

func (s *ImagesTestSuite) TestRetrieveImageByID() {

	assert := assert.New(s.T())

	// Grab a real Image
	seedImage := s.seedImages[0]

	imageByID, err := s.adminOsClient.RetrieveImageByID(seedImage.ID)

	if !assert.Nil(err, "RetrieveImageByID should not have returned an error but did") {
		fmt.Println(err.Error())
	}

	assert.NotNil(imageByID, "RetrieveImageByID returned a nil image but it should not have done")

	assert.Equal(seedImage, *imageByID, "RetrieveImageByID did not return a image equal to the expected image")

}

func (s *ImagesTestSuite) TestRetrieveImageByName() {

	assert := assert.New(s.T())

	// Grab a real image
	seedImage := s.seedImages[0]

	imageByName, err := s.adminOsClient.RetrieveImageByName(seedImage.Name)

	if !assert.Nil(err, "RetrieveImageByName should not have returned an error but did") {
		fmt.Println(err.Error())
	}

	assert.NotNil(imageByName, "RetrieveImageByName returned a nil image but it should not have done")

	assert.Equal(seedImage, *imageByName, "RetrieveImageByName did not return a image equal to the expected image")
}
