package goos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAuthConfig(t *testing.T) {

	cases := []struct {
		cfg         *AuthConfig
		msg         string
		isErrorCase bool
	}{
		{
			cfg:         &AuthConfig{},
			msg:         "Empty config should have returned an error",
			isErrorCase: true,
		},
		{
			cfg: &AuthConfig{
				Password:   "boogies",
				AuthURL:    "http://someplace.go",
				AuthDomain: "CAS",
				TenantName: "ooga-booga",
			},
			msg:         "Config with no User should have returned an error",
			isErrorCase: true,
		},
		{
			cfg: &AuthConfig{
				User:       "jeff",
				AuthURL:    "http://someplace.go",
				AuthDomain: "CAS",
				TenantName: "ooga-booga",
			},
			msg:         "Config with no Password should have returned an error",
			isErrorCase: true,
		},
		{
			cfg: &AuthConfig{
				User:       "jeff",
				Password:   "boogies",
				AuthDomain: "CAS",
				TenantName: "ooga-booga",
			},
			msg:         "Config with no AuthURL should have returned an error",
			isErrorCase: true,
		},
		{
			cfg: &AuthConfig{
				User:       "jeff",
				Password:   "boogies",
				AuthURL:    "http://someplace.go",
				TenantName: "ooga-booga",
			},
			msg:         "Config with no AuthDomain should have returned an error",
			isErrorCase: true,
		},
		{
			cfg: &AuthConfig{
				User:       "jeff",
				Password:   "boogies",
				AuthURL:    "http://someplace.go",
				AuthDomain: "CAS",
			},
			msg:         "Config with no TenantName should have returned an error",
			isErrorCase: true,
		},
		{
			cfg: &AuthConfig{
				User:       "jeff",
				Password:   "boogies",
				AuthURL:    "http://someplace.go",
				AuthDomain: "CAS",
				TenantName: "ooga-booga",
			},
			msg:         "Fully populated AuthConfig should not have returned an error",
			isErrorCase: false,
		},
	}

	for _, c := range cases {

		err := validateAuthConfig(c.cfg)

		if c.isErrorCase {
			assert.NotNil(t, err, c.msg)
		} else {
			assert.Nil(t, err, c.msg)
		}
	}
}
