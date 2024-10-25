package config

import (
	"bootstrap/efi/common"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	var testCases = []struct {
		caseName       string
		efivars        string
		expectedResult *Config
		expectedErr    error
	}{
		{
			caseName:    "not exists efivars",
			efivars:     "fixtures/not_exists",
			expectedErr: os.ErrNotExist,
		},
		{
			caseName: "valid efivars with dynamic DHCP for IPv4 and IPv6",
			efivars:  "fixtures/valid_dhcp",
			expectedResult: &[]Config{{
				MAC:  "3c:ec:ef:c4:45:80",
				VLAN: 1,
				IPv4: IPv4{
					Static:  false,
					Address: "0.0.0.0/0",
					Gateway: "0.0.0.0",
				},
				IPv6: IPv6{
					Static:    false,
					ForceDHCP: false,
					Address:   "::/0",
					Gateway:   "::",
				},
				DNS: []string{
					"192.168.0.1",
					"192.168.0.2",
				},
				URI: "http://www.google.com/",
			}}[0],
		},
		{
			caseName: "valid efivars with dynamic DHCP for IPv4 and IPv6 (force IPv6 DHCP)",
			efivars:  "fixtures/valid_dhcp_force6",
			expectedResult: &[]Config{{
				MAC:  "3c:ec:ef:c4:45:80",
				VLAN: 1,
				IPv4: IPv4{
					Static:  false,
					Address: "0.0.0.0/0",
					Gateway: "0.0.0.0",
				},
				IPv6: IPv6{
					Static:    false,
					ForceDHCP: true,
					Address:   "::/0",
					Gateway:   "::",
				},
				DNS: []string{
					"192.168.0.1",
					"192.168.0.2",
				},
				URI: "http://www.google.com/",
			}}[0],
		},
		{
			caseName: "valid efivars with static addresses for IPv4 and IPv6",
			efivars:  "fixtures/valid_static",
			expectedResult: &[]Config{{
				MAC:  "3c:ec:ef:c4:45:80",
				VLAN: 1,
				IPv4: IPv4{
					Static:  true,
					Address: "192.168.0.11/24",
					Gateway: "192.168.0.254",
				},
				IPv6: IPv6{
					Static:    true,
					ForceDHCP: false,
					Address:   "bdf9:564f:9126:fcfa:a1c8:818:d700:44f2/64",
					Gateway:   "::",
				},
				DNS: []string{
					"192.168.0.1",
					"192.168.0.2",
				},
				URI: "http://www.google.com/",
			}}[0],
		},
		{
			caseName:    "invalid efivars with invalid load option",
			efivars:     "fixtures/invalid_lo",
			expectedErr: common.ErrDataSize,
		},
		{
			caseName:    "invalid efivars with invalid DNS",
			efivars:     "fixtures/invalid_dns",
			expectedErr: common.ErrDataRepresentation,
		},
		{
			caseName:    "invalid efivars with invalid IPv4",
			efivars:     "fixtures/invalid_ipv4",
			expectedErr: common.ErrDataRepresentation,
		},
		{
			caseName:    "invalid efivars with invalid IPv6",
			efivars:     "fixtures/invalid_ipv6",
			expectedErr: common.ErrDataRepresentation,
		},
		{
			caseName:    "invalid efivars with invalid MAC",
			efivars:     "fixtures/invalid_mac",
			expectedErr: common.ErrDataRepresentation,
		},
		{
			caseName:    "invalid efivars with invalid URI",
			efivars:     "fixtures/invalid_uri",
			expectedErr: common.ErrDataRepresentation,
		},
		{
			caseName:    "invalid efivars with invalid VLAN",
			efivars:     "fixtures/invalid_vlan",
			expectedErr: common.ErrDataRepresentation,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			cfg, err := Load(testCase.efivars)

			if testCase.expectedErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, testCase.expectedErr)
				assert.Nil(t, cfg)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, cfg)
				assert.Equal(t, testCase.expectedResult, cfg)
			}
		})
	}
}