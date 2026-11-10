package awsecs


// Properties for defining a `CfnGatewayClusterAssociations`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayClusterAssociationsProps := &CfnGatewayClusterAssociationsProps{
//   	Cluster: jsii.String("cluster"),
//   	Gateway: jsii.String("gateway"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gatewayclusterassociations.html
//
type CfnGatewayClusterAssociationsProps struct {
	// The name or full Amazon Resource Name (ARN) of the cluster to associate the gateway with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gatewayclusterassociations.html#cfn-ecs-gatewayclusterassociations-cluster
	//
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The name or full Amazon Resource Name (ARN) of the gateway to associate with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-gatewayclusterassociations.html#cfn-ecs-gatewayclusterassociations-gateway
	//
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
}

