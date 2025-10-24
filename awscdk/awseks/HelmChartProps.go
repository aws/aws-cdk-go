package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Helm Chart properties.
//
// Example:
//   var cluster Cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("MyOCIChart"), &HelmChartProps{
//   	Cluster: Cluster,
//   	Chart: jsii.String("some-chart"),
//   	Repository: jsii.String("oci://${ACCOUNT_ID}.dkr.ecr.${ACCOUNT_REGION}.amazonaws.com/${REPO_NAME}"),
//   	Namespace: jsii.String("oci"),
//   	Version: jsii.String("0.0.1"),
//   })
//
type HelmChartProps struct {
	// Whether or not Helm should treat this operation as atomic;
	//
	// if set, upgrade process rolls back changes
	// made in case of failed upgrade. The --wait flag will be set automatically if --atomic is used.
	// Default: false.
	//
	Atomic *bool `field:"optional" json:"atomic" yaml:"atomic"`
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Default: - No chart name. Implies `chartAsset` is used.
	//
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Default: - No chart asset. Implies `chart` is used.
	//
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Default: true.
	//
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Default: default.
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Default: - If no release name is given, it will use the last 53 characters of the node's unique id.
	//
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://charts.helm.sh/stable/
	// Default: - No repository will be used, which means that the chart needs to be an absolute URL.
	//
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// if set, no CRDs will be installed.
	// Default: - CRDs are installed if not already present.
	//
	SkipCrds *bool `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Default: Duration.minutes(5)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	//
	// For nested values use a nested dictionary. For example:
	// values: {
	//  installationCRDs: true,
	//  webhook: { port: 9443 }
	// }.
	// Default: - No values are provided to the chart.
	//
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	// Default: - If this is not specified, the latest version is installed.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Default: - Helm will not wait before marking release as successful.
	//
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

