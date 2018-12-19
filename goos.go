package goos

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	logger "github.com/sirupsen/logrus"
)

// AuthConfig provides values required for authenticating to an OpenStack cloud/domain/tenant
type AuthConfig struct {
	User       string
	Password   string
	AuthURL    string
	AuthDomain string
	TenantName string
}

// CreateProviderClient creates an authenticated ProviderClient for use by Gopher services
func CreateProviderClient(authConfig *AuthConfig) (*gophercloud.ProviderClient, error) {

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: authConfig.AuthURL,
		Username:         authConfig.User,
		Password:         authConfig.Password,
		TenantName:       authConfig.TenantName,
		DomainName:       authConfig.AuthDomain,
	}

	authClient, err := openstack.AuthenticatedClient(authOpts)

	if err != nil {
		logger.Fatal("Failed to establish an authenticated OpenStack client " + err.Error())
	}

	return authClient, err
}

// RetrieveFlavors lists all of the flavors of the project currently indicated by the given ProviderClient
func RetrieveFlavors(provider *gophercloud.ProviderClient) ([]flavors.Flavor, error) {

	compute, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return nil, err
	}

	listOpts := flavors.ListOpts{
		Limit:      20,
		AccessType: flavors.PublicAccess,
	}

	allPages, err := flavors.ListDetail(compute, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		return nil, err
	}

	return allFlavors, nil
}

// RetrieveImages lists all of the images of the project currently indicated by the given ProviderClient
func RetrieveImages(provider *gophercloud.ProviderClient) ([]images.Image, error) {

	compute, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return nil, err
	}

	listOpts := images.ListOpts{
		Limit: 20,
	}

	allPages, err := images.ListDetail(compute, listOpts).AllPages()
	if err != nil {
		return nil, err
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		return nil, err
	}

	return allImages, nil
}
