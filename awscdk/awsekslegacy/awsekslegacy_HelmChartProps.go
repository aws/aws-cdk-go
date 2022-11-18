package awsekslegacy


// Helm Chart properties.
//
// Example:
//   var cluster cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("NginxIngress"), &helmChartProps{
//   	cluster: cluster,
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
//   // or, option2: use `addChart`
//   cluster.addChart(jsii.String("NginxIngress"), &helmChartOptions{
//   	chart: jsii.String("nginx-ingress"),
//   	repository: jsii.String("https://helm.nginx.com/stable"),
//   	namespace: jsii.String("kube-system"),
//   })
//
// Experimental.
type HelmChartProps struct {
	// The name of the chart.
	// Experimental.
	Chart *string `field:"required" json:"chart" yaml:"chart"`
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
	// The values to be used by the chart.
	// Experimental.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

