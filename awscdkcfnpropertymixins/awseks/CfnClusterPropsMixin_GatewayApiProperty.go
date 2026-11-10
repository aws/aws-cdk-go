package awseks


// The Gateway API configuration for the load balancing capability of EKS Auto Mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayApiProperty := &GatewayApiProperty{
//   	Crds: &GatewayApiCrdsProperty{
//   		Managed: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-gatewayapi.html
//
type CfnClusterPropsMixin_GatewayApiProperty struct {
	// The Gateway API Custom Resource Definition (CRD) configuration for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-gatewayapi.html#cfn-eks-cluster-gatewayapi-crds
	//
	Crds interface{} `field:"optional" json:"crds" yaml:"crds"`
}

