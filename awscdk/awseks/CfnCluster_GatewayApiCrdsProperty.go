package awseks


// The Gateway API Custom Resource Definition (CRD) configuration for the cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayApiCrdsProperty := &GatewayApiCrdsProperty{
//   	Managed: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-gatewayapicrds.html
//
type CfnCluster_GatewayApiCrdsProperty struct {
	// Indicates whether EKS manages the lifecycle of the Gateway API CRDs on the cluster.
	//
	// If true, EKS installs and upgrades the Gateway API CRDs. Requires ElasticLoadBalancing to be enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-gatewayapicrds.html#cfn-eks-cluster-gatewayapicrds-managed
	//
	Managed interface{} `field:"required" json:"managed" yaml:"managed"`
}

