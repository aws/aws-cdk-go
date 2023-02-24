package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Helm Chart options.
//
// Example:
//   import s3Assets "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   chartAsset := s3Assets.NewAsset(this, jsii.String("ChartAsset"), &AssetProps{
//   	Path: jsii.String("/path/to/asset"),
//   })
//
//   cluster.addHelmChart(jsii.String("test-chart"), &HelmChartOptions{
//   	ChartAsset: chartAsset,
//   })
//
type HelmChartOptions struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://charts.helm.sh/stable/
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	//
	// For nested values use a nested dictionary. For example:
	// values: {
	//   installationCRDs: true,
	//   webhook: { port: 9443 }
	// }.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
}

