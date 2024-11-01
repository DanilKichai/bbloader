package generator

import (
	"bootstrap/config"
	"bootstrap/unyaml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	var testCases = []struct {
		caseName          string
		file              string
		expectedResult    *unyaml.UnYAML
		expectedErr       error
		expecterErrSubstr *string
	}{
		{
			caseName: "valid template",
			file:     "fixtures/bootstrap_valid.tmpl",
			expectedResult: &[]unyaml.UnYAML{
				{
					{
						Path: "/run/systemd/network/eth0.network",
						Type: "file",
						Data: "[Match]\nName=eth0\n\n[Network]\nAddress=\"192.168.0.101/24\"\nGateway=\"192.168.0.1\"\n",
					},
				},
			}[0],
		},
		{
			caseName:    "not exists",
			file:        "fixtures/bootstrap_not_exists.tmpl",
			expectedErr: os.ErrNotExist,
		},
		{
			caseName:          "invalid template with errors at execution",
			file:              "fixtures/bootstrap_invalid(execution).tmpl",
			expecterErrSubstr: &[]string{"execute template: "}[0],
		},
		{
			caseName:          "invalid template with errors at parsing",
			file:              "fixtures/bootstrap_invalid(parsing).tmpl",
			expecterErrSubstr: &[]string{"parse template: "}[0],
		},
		{
			caseName:          "invalid template with errors at unmarshalling",
			file:              "fixtures/bootstrap_invalid(unmarshalling).tmpl",
			expecterErrSubstr: &[]string{"unmarshal YAML: "}[0],
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			var tc = struct {
				Config *config.Config
			}{
				Config: &[]config.Config{
					{
						IPv4: config.IPv4{
							Address: "192.168.0.101/24",
							Gateway: "192.168.0.1",
						},
						DNS: []string{
							"192.168.0.1",
						},
					},
				}[0],
			}

			gen, err := Load(testCase.file, tc)

			if testCase.expectedErr != nil || testCase.expecterErrSubstr != nil {
				require.Error(t, err)
				if testCase.expectedErr != nil {
					assert.ErrorIs(t, err, testCase.expectedErr)
				}

				if testCase.expecterErrSubstr != nil {
					assert.ErrorContains(t, err, *testCase.expecterErrSubstr)
				}
				assert.Nil(t, gen)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, gen)
				assert.Equal(t, testCase.expectedResult, gen)
			}
		})
	}
}
