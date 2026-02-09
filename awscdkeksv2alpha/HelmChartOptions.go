package awscdkeksv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Helm Chart options.
//
// Example:
//   import s3Assets "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster Cluster
//
//   chartAsset := s3Assets.NewAsset(this, jsii.String("ChartAsset"), &AssetProps{
//   	Path: jsii.String("/path/to/asset"),
//   })
//
//   cluster.addHelmChart(jsii.String("test-chart"), &HelmChartOptions{
//   	ChartAsset: chartAsset,
//   })
//
// Experimental.
type HelmChartOptions struct {
	// Whether or not Helm should treat this operation as atomic;
	//
	// if set, upgrade process rolls back changes
	// made in case of failed upgrade. The --wait flag will be set automatically if --atomic is used.
	// Default: false.
	//
	// Experimental.
	Atomic *bool `field:"optional" json:"atomic" yaml:"atomic"`
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Default: - No chart name. Implies `chartAsset` is used.
	//
	// Experimental.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Default: - No chart asset. Implies `chart` is used.
	//
	// Experimental.
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Default: true.
	//
	// Experimental.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Default: default.
	//
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Default: - If no release name is given, it will use the last 53 characters of the node's unique id.
	//
	// Experimental.
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The removal policy applied to the custom resource that manages the Helm chart.
	//
	// The removal policy controls what happens to the resource if it stops being managed by CloudFormation.
	// This can happen in one of three situations:
	//
	// - The resource is removed from the template, so CloudFormation stops managing it
	// - A change to the resource is made that requires it to be replaced, so CloudFormation stops managing it
	// - The stack is deleted, so CloudFormation stops managing all resources in it.
	// Default: RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The repository which contains the chart.
	//
	// For example: https://charts.helm.sh/stable/
	// Default: - No repository will be used, which means that the chart needs to be an absolute URL.
	//
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// if set, no CRDs will be installed.
	// Default: - CRDs are installed if not already present.
	//
	// Experimental.
	SkipCrds *bool `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Default: Duration.minutes(5)
	//
	// Experimental.
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
	// Experimental.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	// Default: - If this is not specified, the latest version is installed.
	//
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Default: - Helm will not wait before marking release as successful.
	//
	// Experimental.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
}

