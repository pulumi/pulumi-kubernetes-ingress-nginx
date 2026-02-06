//go:build nodejs || all
// +build nodejs all

package examples

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
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
		"TestSimpleIngressNginxTs": {directoryName: "simple-nginx-ts"},
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

func TestTsIngressNginxMapProperties(t *testing.T) {
	t.Run("TestSimpleIngressNginxTsConfig", func(t *testing.T) {
		p := pulumitest.NewPulumiTest(t, "simple-nginx-ts",
			opttest.LocalProviderPath("pulumi-kubernetes-ingress-nginx", filepath.Join(getCwd(t), "..", "bin")),
			opttest.YarnLink("@pulumi/kubernetes-ingress-nginx"),
		)

		namespace := p.Up(t).Outputs["namespace"].Value.(string)

		labelSelector := "app.kubernetes.io/component=controller"
		getConfigMapCmd := exec.Command("kubectl", "get", "configmap", "-n", namespace,
			"-l", labelSelector, "-o", "jsonpath={.items[0].metadata.name}")
		var configMapOut bytes.Buffer
		getConfigMapCmd.Stdout = &configMapOut
		getConfigMapCmd.Stderr = &configMapOut

		if err := getConfigMapCmd.Run(); err != nil {
			t.Fatalf("Failed to get ConfigMap name: %v\nOutput: %s", err, configMapOut.String())
		}
		configMapName := strings.TrimSpace(configMapOut.String())
		t.Logf("Retrieved ConfigMap name: %s", configMapName)

		// Step 3: Get the hsts-max-age value from the dynamically retrieved ConfigMap
		kubectlCmd := exec.Command(
			"kubectl",
			"get",
			"configmap",
			configMapName,
			"-n",
			namespace,
			"-o",
			"jsonpath={.data.hsts-max-age}",
		)
		var kubectlOut bytes.Buffer
		kubectlCmd.Stdout = &kubectlOut
		kubectlCmd.Stderr = &kubectlOut

		if err := kubectlCmd.Run(); err != nil {
			t.Fatalf("Kubectl command failed: %v\nOutput: %s", err, kubectlOut.String())
		}
		hstsMaxAge := kubectlOut.String()
		t.Logf("hsts-max-age from ConfigMap: %s", hstsMaxAge)

		// Validate the value of hsts-max-age
		expectedValue := "31536000"
		if hstsMaxAge != expectedValue {
			t.Errorf("hsts-max-age mismatch: expected %s, got %s", expectedValue, hstsMaxAge)
		}
	})
}
