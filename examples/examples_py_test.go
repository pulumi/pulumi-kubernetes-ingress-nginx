//go:build python || all
// +build python all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	basePython := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePython
}

// func TestPyExamples(t *testing.T) {
// 	tests := map[string]struct {
// 		directoryName    string
// 		additionalConfig map[string]string
// 	}{
// 		"TestSimpleIngressNginxPy": {directoryName: "simple-nginx-py"},
// 	}
// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			p := pulumitest.NewPulumiTest(t, test.directoryName,
// 				opttest.LocalProviderPath("pulumi-kubernetes-ingress-nginx", filepath.Join(getCwd(t), "..", "bin")),
// 			)
// 			if test.additionalConfig != nil {
// 				for key, value := range test.additionalConfig {
// 					p.SetConfig(t, key, value)
// 				}
// 			}
// 			p.Install(t)
// 			p.Up(t)
// 			p.Preview(t, optpreview.ExpectNoChanges())
// 			p.Refresh(t, optrefresh.ExpectNoChanges())
// 		})
// 	}
// }
