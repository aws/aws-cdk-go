package awsec2


// Properties for defining a `CfnEgressOnlyInternetGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEgressOnlyInternetGatewayProps := &CfnEgressOnlyInternetGatewayProps{
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html
//
type CfnEgressOnlyInternetGatewayProps struct {
	// The ID of the VPC for which to create the egress-only internet gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html#cfn-ec2-egressonlyinternetgateway-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

