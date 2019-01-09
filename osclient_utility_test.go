package goos

import (
	"errors"

	"github.com/spf13/viper"
)

type authConfigCreator interface {
	Create() (*AuthConfig, error)
}

type envConfigurator struct {
}

func (c *envConfigurator) Create() (*AuthConfig, error) {

	// We'll use Viper for easy extraction of ENV vars.
	// Note that Viper maintains its data globally, but that makes testing
	// difficult.  Instead, we'll create a new Viper instance and maintain that.
	v := viper.New()

	v.BindEnv("OS_USERNAME")
	v.BindEnv("OS_PASSWORD")
	v.BindEnv("OS_USER_DOMAIN_NAME")
	v.BindEnv("OS_PROJECT_NAME")
	v.BindEnv("OS_AUTH_URL")

	user := v.GetString("OS_USERNAME")

	if user == "" {
		return &AuthConfig{}, errors.New("No OS_USERNAME was provided in the environment")
	}

	pass := v.GetString("OS_PASSWORD")

	if pass == "" {
		return &AuthConfig{}, errors.New("No OS_PASSWORD was provided in the environment")
	}

	authDomain := v.GetString("OS_USER_DOMAIN_NAME")

	if authDomain == "" {
		return &AuthConfig{}, errors.New("No OS_USER_DOMAIN_NAME was provided in the environment")
	}

	authURL := v.GetString("OS_AUTH_URL")

	if authURL == "" {
		return &AuthConfig{}, errors.New("No OS_AUTH_URL was provided in the environment")
	}

	ten := v.GetString("OS_PROJECT_NAME")

	if ten == "" {
		return &AuthConfig{}, errors.New("No OS_PROJECT_NAME was provided in the environment")
	}

	return &AuthConfig{
		User:       user,
		Password:   pass,
		AuthURL:    authURL,
		AuthDomain: authDomain,
		TenantName: ten,
	}, nil
}

func getAdminAuthConfig() *AuthConfig {

	ac, err := getAuthConfigCreator().Create()

	if err != nil {
		panic(err)
	}

	return ac
}

// For now, we'll simply pull the AuthConfig from the local environment.  This punts the problem
// to the execution environment of the test suite.  However, it would be better if this could be
// more under the control of the testing code itself.  e.g. Perhaps pull the password for an automated
// testing account from Vault but pull all of the rest of the environment from code.  This would allow
// secure, authenticated execution of deterministic tests and put the control where we need it in
// the test suite.  Additionally, it would allow the ability to create multiple AuthConfigs that have
// different user accounts to test different sets of authorizations (e.g. run tests to create an instance
// with an account that doesn't have rights to do that).
func getAuthConfigCreator() authConfigCreator {
	return &envConfigurator{}
}
