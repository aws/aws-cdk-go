package awseks


// Helm chart options that can be set for AlbControllerChart To add any new supported values refer https://github.com/kubernetes-sigs/aws-load-balancer-controller/blob/main/helm/aws-load-balancer-controller/values.yaml.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv34"
//
//
//   eks.NewCluster(this, jsii.String("HelloEKS"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_34(),
//   	AlbController: &AlbControllerOptions{
//   		Version: eks.AlbControllerVersion_V2_8_2(),
//   		AdditionalHelmChartValues: &AlbControllerHelmChartOptions{
//   			EnableWafv2: jsii.Boolean(false),
//   		},
//   	},
//   	KubectlLayer: kubectlv34.NewKubectlV34Layer(this, jsii.String("kubectl")),
//   })
//
type AlbControllerHelmChartOptions struct {
	// Enable or disable AWS WAF on the ALB ingress controller.
	// Default: - no value defined for this helm chart option, so it will not be set in the helm chart values.
	//
	EnableWaf *bool `field:"optional" json:"enableWaf" yaml:"enableWaf"`
	// Enable or disable AWS WAFv2 on the ALB ingress controller.
	// Default: - no value defined for this helm chart option, so it will not be set in the helm chart values.
	//
	EnableWafv2 *bool `field:"optional" json:"enableWafv2" yaml:"enableWafv2"`
}

