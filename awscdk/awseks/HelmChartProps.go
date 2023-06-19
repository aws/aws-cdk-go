package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
)

// Helm Chart properties.
//
// Example:
//   var cluster cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("NginxIngress"), &HelmChartProps{
//   	Cluster: Cluster,
//   	Chart: jsii.String("nginx-ingress"),
//   	Repository: jsii.String("https://helm.nginx.com/stable"),
//   	Namespace: jsii.String("kube-system"),
//   })
//
//   // or, option2: use `addHelmChart`
//   cluster.addHelmChart(jsii.String("NginxIngress"), &HelmChartOptions{
//   	Chart: jsii.String("nginx-ingress"),
//   	Repository: jsii.String("https://helm.nginx.com/stable"),
//   	Namespace: jsii.String("kube-system"),
//   })
//
// Experimental.
type HelmChartProps struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	// Experimental.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	// Experimental.
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	// Experimental.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	// Experimental.
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://kubernetes-charts.storage.googleapis.com/
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	// Experimental.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	// Experimental.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

