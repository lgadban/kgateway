package settings_test

import (
	"os"
	"testing"

	"github.com/onsi/gomega"

	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/extensions2/settings"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/wellknown"
)

func TestSettings(t *testing.T) {
	testCases := []struct {
		// name of the test case
		name string

		// env vars that are set at the beginning of test (and removed after test)
		envVars map[string]string

		// if set, then these are the expected populated settings
		expectedSettings *settings.Settings

		// if set, then an error parsing the settings is expected to occur
		expectedErrorStr string
	}{
		{
			// TODO: this test case does not fail when a new field is added to Settings
			// but not updated here. should it?
			name:    "defaults to empty or default values",
			envVars: map[string]string{},
			expectedSettings: &settings.Settings{
				DnsLookupFamily:        "V4_PREFERRED",
				EnableIstioIntegration: false,
				EnableAutoMtls:         false,
				StsClusterName:         "",
				StsUri:                 "",
				XdsServiceName:         wellknown.DefaultXdsService,
				XdsServicePort:         wellknown.DefaultXdsPort,
				UseRustFormations:      false,
			},
		},
		{
			name: "all values set",
			envVars: map[string]string{
				"KGW_DNS_LOOKUP_FAMILY":        "V4_ONLY",
				"KGW_ENABLE_ISTIO_INTEGRATION": "true",
				"KGW_ENABLE_AUTO_MTLS":         "true",
				"KGW_STS_CLUSTER_NAME":         "my-cluster",
				"KGW_STS_URI":                  "my.sts.uri",
				"KGW_XDS_SERVICE_NAME":         "custom-svc",
				"KGW_XDS_SERVICE_PORT":         "1234",
				"KGW_USE_RUST_FORMATIONS":      "true",
			},
			expectedSettings: &settings.Settings{
				DnsLookupFamily:        "V4_ONLY",
				EnableIstioIntegration: true,
				EnableAutoMtls:         true,
				StsClusterName:         "my-cluster",
				StsUri:                 "my.sts.uri",
				XdsServiceName:         "custom-svc",
				XdsServicePort:         1234,
				UseRustFormations:      true,
			},
		},
		{
			name: "errors on invalid bool",
			envVars: map[string]string{
				"KGW_ENABLE_ISTIO_INTEGRATION": "true123",
			},
			expectedErrorStr: "invalid syntax",
		},
		{
			name: "errors on invalid port",
			envVars: map[string]string{
				"KGW_XDS_SERVICE_PORT": "a123",
			},
			expectedErrorStr: "invalid syntax",
		},
		{
			name: "ignores other env vars",
			envVars: map[string]string{
				"KGW_DOES_NOT_EXIST":   "true",
				"ANOTHER_VAR":          "abc",
				"KGW_ENABLE_AUTO_MTLS": "true",
			},
			expectedSettings: &settings.Settings{
				DnsLookupFamily: "V4_PREFERRED",
				EnableAutoMtls:  true,
				XdsServiceName:  wellknown.DefaultXdsService,
				XdsServicePort:  wellknown.DefaultXdsPort,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := gomega.NewWithT(t)

			t.Cleanup(func() {
				for k := range tc.envVars {
					err := os.Unsetenv(k)
					g.Expect(err).NotTo(gomega.HaveOccurred())
				}
			})

			for k, v := range tc.envVars {
				err := os.Setenv(k, v)
				g.Expect(err).NotTo(gomega.HaveOccurred())
			}
			s, err := settings.BuildSettings()
			if tc.expectedErrorStr != "" {
				g.Expect(err).To(gomega.HaveOccurred())
				g.Expect(err.Error()).To(gomega.ContainSubstring(tc.expectedErrorStr))
			} else {
				g.Expect(err).NotTo(gomega.HaveOccurred())
				g.Expect(s).To(gomega.Equal(tc.expectedSettings))
			}
		})
	}
}
