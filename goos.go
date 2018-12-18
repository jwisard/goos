package goos

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	logger "github.com/sirupsen/logrus"
)

// AuthConfig provides values required for authenticating to an OpenStack cloud/domain/tenant
type AuthConfig struct {
	user       string
	password   string
	authURL    string
	authDomain string
	tenantName string
}

// CreateProviderClient creates an authenticated ProviderClient for use by Gopher services
func CreateProviderClient(authConfig *AuthConfig) (*gophercloud.ProviderClient, error) {

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: authConfig.authURL,
		Username:         authConfig.user,
		Password:         authConfig.password,
		TenantName:       authConfig.tenantName,
		DomainName:       authConfig.authDomain,
	}

	authClient, err := openstack.AuthenticatedClient(authOpts)

	if err != nil {
		logger.Fatal("Failed to establish an authenticated OpenStack client " + err.Error())
	}

	return authClient, err
}
