package goos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	err := validateAuthConfig(&AuthConfig{})
	assert.True(t, err != nil, "Empty config should have returned an error")

	err = validateAuthConfig(&AuthConfig{
		Password:   "boogies",
		AuthURL:    "http://someplace.go",
		AuthDomain: "CAS",
		TenantName: "ooga-booga",
	})
	assert.True(t, err != nil, "Config with no User should have returned an error")

	err = validateAuthConfig(&AuthConfig{
		User:       "jeff",
		AuthURL:    "http://someplace.go",
		AuthDomain: "CAS",
		TenantName: "ooga-booga",
	})
	assert.True(t, err != nil, "Config with no Password should have returned an error")

	err = validateAuthConfig(&AuthConfig{
		User:       "jeff",
		Password:   "boogies",
		AuthDomain: "CAS",
		TenantName: "ooga-booga",
	})
	assert.True(t, err != nil, "Config with no AuthURL should have returned an error")

	err = validateAuthConfig(&AuthConfig{
		User:       "jeff",
		Password:   "boogies",
		AuthURL:    "http://someplace.go",
		TenantName: "ooga-booga",
	})
	assert.True(t, err != nil, "Config with no AuthDomain should have returned an error")

	err = validateAuthConfig(&AuthConfig{
		User:       "jeff",
		Password:   "boogies",
		AuthURL:    "http://someplace.go",
		AuthDomain: "CAS",
	})
	assert.True(t, err != nil, "Config with no TenantName should have returned an error")

	err = validateAuthConfig(&AuthConfig{
		User:       "jeff",
		Password:   "boogies",
		AuthURL:    "http://someplace.go",
		AuthDomain: "CAS",
		TenantName: "ooga-booga",
	})
	assert.True(t, err == nil, "Fully populated AuthConfig should not have returned an error")
}
