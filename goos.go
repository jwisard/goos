package goos

import (
	"errors"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	logger "github.com/sirupsen/logrus"
)

// OSClient provides methods for interacting with OpenStack
type OSClient interface {
	RetrieveFlavors() ([]flavors.Flavor, error)
	RetrieveFlavorByID(flavorID string) (*flavors.Flavor, error)
	RetrieveFlavorByName(flavorName string) (*flavors.Flavor, error)
	RetrieveImages() ([]images.Image, error)
	RetrieveImageByID(imageID string) (*images.Image, error)
	RetrieveImageByName(imageName string) (*images.Image, error)
}

type provider struct {
	providerClient *gophercloud.ProviderClient
}

// AuthConfig provides values required for authenticating to an OpenStack cloud/domain/tenant
type AuthConfig struct {
	User       string
	Password   string
	AuthURL    string
	AuthDomain string
	TenantName string
}

// CreateOSClient is a factory for creating a authenticated OSClient instances
func CreateOSClient(authConfig *AuthConfig) (OSClient, error) {

	if err := validateAuthConfig(authConfig); err != nil {
		return nil, err
	}

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: authConfig.AuthURL,
		Username:         authConfig.User,
		Password:         authConfig.Password,
		TenantName:       authConfig.TenantName,
		DomainName:       authConfig.AuthDomain,
	}

	authClient, err := openstack.AuthenticatedClient(authOpts)

	if err != nil {
		logger.Error("Failed to establish an authenticated OpenStack client: " + err.Error())
		return nil, err
	}

	return &provider{providerClient: authClient}, err
}

func validateAuthConfig(authConfig *AuthConfig) error {

	if authConfig.User == "" {
		return errors.New("No goos.AuthConfig.User was provided")
	}

	if authConfig.Password == "" {
		return errors.New("No goos.AuthConfig.Password was provided")
	}

	if authConfig.AuthURL == "" {
		return errors.New("No goos.AuthConfig.AuthURL was provided")
	}

	if authConfig.AuthDomain == "" {
		return errors.New("No goos.AuthConfig.AuthDomain was provided")
	}

	if authConfig.TenantName == "" {
		return errors.New("No goos.AuthConfig.TenantName was provided")
	}

	return nil
}

// RetrieveFlavorByID returns a Flavor from its ID. If no flavor has the given flavorID then
// and error will occur.
func (p *provider) RetrieveFlavorByID(flavorID string) (*flavors.Flavor, error) {

	compute, err := openstack.NewComputeV2(p.providerClient, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return &flavors.Flavor{}, err
	}

	return flavors.Get(compute, flavorID).Extract()
}

// RetrieveFlavorByName returns a Flavor from its name. If more than one flavor shares the name or
// if no flavor matches the name exactly (case sensitive), then an error will occur.
func (p *provider) RetrieveFlavorByName(flavorName string) (*flavors.Flavor, error) {

	compute, err := openstack.NewComputeV2(p.providerClient, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return &flavors.Flavor{}, err
	}

	id, err := flavors.IDFromName(compute, flavorName)

	if err != nil {
		return &flavors.Flavor{}, err
	}

	return p.RetrieveFlavorByID(id)
}

// RetrieveFlavors lists all of the flavors of the project currently indicated by the given ProviderClient
func (p *provider) RetrieveFlavors() ([]flavors.Flavor, error) {

	compute, err := openstack.NewComputeV2(p.providerClient, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return nil, err
	}

	listOpts := flavors.ListOpts{
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

// RetrieveImageByID returns an Image from its ID. If no image has the given imageID then
// and error will occur.
func (p *provider) RetrieveImageByID(imageID string) (*images.Image, error) {

	compute, err := openstack.NewComputeV2(p.providerClient, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return &images.Image{}, err
	}

	return images.Get(compute, imageID).Extract()
}

// RetrieveImageByName returns a Image from its name. If more than one image shares the name or
// if no image matches the name exactly (case sensitive), then an error will occur.
func (p *provider) RetrieveImageByName(imageName string) (*images.Image, error) {

	compute, err := openstack.NewComputeV2(p.providerClient, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return &images.Image{}, err
	}

	id, err := images.IDFromName(compute, imageName)

	if err != nil {
		return &images.Image{}, err
	}

	return p.RetrieveImageByID(id)
}

// RetrieveImages lists all of the images of the project currently indicated by the given ProviderClient
func (p *provider) RetrieveImages() ([]images.Image, error) {

	compute, err := openstack.NewComputeV2(p.providerClient, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	if err != nil {
		return nil, err
	}

	listOpts := images.ListOpts{}

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
