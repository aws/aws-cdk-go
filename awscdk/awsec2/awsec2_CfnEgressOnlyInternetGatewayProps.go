package awsec2


// Properties for defining a `CfnEgressOnlyInternetGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEgressOnlyInternetGatewayProps := &cfnEgressOnlyInternetGatewayProps{
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnEgressOnlyInternetGatewayProps struct {
	// The ID of the VPC for which to create the egress-only internet gateway.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

