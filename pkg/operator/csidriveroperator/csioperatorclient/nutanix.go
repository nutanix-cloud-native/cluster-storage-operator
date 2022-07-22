package csioperatorclient

import (
	"os"
	"strings"

	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	NutanixCSIDriverName          = "csi.nutanix.com"
	envNutanixDriverOperatorImage = "NUTANIX_DRIVER_OPERATOR_IMAGE"
	envNutanixDriverImage         = "NUTANIX_DRIVER_IMAGE"
)

func GetNutanixCSIOperatorConfig() CSIOperatorConfig {
	pairs := []string{
		"${OPERATOR_IMAGE}", os.Getenv(envNutanixDriverOperatorImage),
		"${DRIVER_IMAGE}", os.Getenv(envNutanixDriverImage),
	}

	olmOpts := &OLMOptions{
		OLMPackageName:            "nutanixcsioperator",
		OLMOperatorDeploymentName: "nutanixcsioperator",
		CRResource: schema.GroupVersionResource{
			Group:    "crd.nutanix.com",
			Version:  "v1alpha1",
			Resource: "nutanixcsistorage",
		},
	}
	return CSIOperatorConfig{
		CSIDriverName:   NutanixCSIDriverName,
		ConditionPrefix: "NUTANIX",
		Platform:        configv1.NutanixPlatformType,
		StaticAssets: []string{
			"csidriveroperators/nutanix/02_sa.yaml",
			"csidriveroperators/nutanix/03_role.yaml",
			"csidriveroperators/nutanix/04_rolebinding.yaml",
			"csidriveroperators/nutanix/05_clusterrole.yaml",
			"csidriveroperators/nutanix/06_clusterrolebinding.yaml",
		},
		DeploymentAsset: "csidriveroperators/nutanix/07_deployment.yaml",
		CRAsset:         "csidriveroperators/nutanix/08_cr.yaml",
		ImageReplacer:   strings.NewReplacer(pairs...),
		AllowDisabled:   false,
		OLMOptions:      olmOpts,
	}
}
