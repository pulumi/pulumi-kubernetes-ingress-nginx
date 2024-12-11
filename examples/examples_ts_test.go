//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optpreview"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optrefresh"
)

func TestTsExamples(t *testing.T) {
	tests := map[string]struct {
		directoryName    string
		additionalConfig map[string]string
	}{
		"TestSimpleCertManagerTs": {directoryName: "simple-nginx-ts"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := pulumitest.NewPulumiTest(t, test.directoryName,
				opttest.LocalProviderPath("pulumi-kubernetes-ingress-nginx", filepath.Join(getCwd(t), "..", "bin")),
				opttest.YarnLink("@pulumi/kubernetes-ingress-nginx"),
			)
			if test.additionalConfig != nil {
				for key, value := range test.additionalConfig {
					p.SetConfig(t, key, value)
				}
			}
			p.Up(t)
			p.Preview(t, optpreview.ExpectNoChanges())
			p.Refresh(t, optrefresh.ExpectNoChanges())
		})
	}
}
