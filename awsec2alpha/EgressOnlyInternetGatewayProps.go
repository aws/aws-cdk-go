package awsec2alpha


// Properties to define an egress-only internet gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   var vpcV2 vpcV2
//
//   egressOnlyInternetGatewayProps := &EgressOnlyInternetGatewayProps{
//   	Vpc: vpcV2,
//
//   	// the properties below are optional
//   	EgressOnlyInternetGatewayName: jsii.String("egressOnlyInternetGatewayName"),
//   }
//
// Experimental.
type EgressOnlyInternetGatewayProps struct {
	// The ID of the VPC for which to create the egress-only internet gateway.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// The resource name of the egress-only internet gateway.
	// Default: none.
	//
	// Experimental.
	EgressOnlyInternetGatewayName *string `field:"optional" json:"egressOnlyInternetGatewayName" yaml:"egressOnlyInternetGatewayName"`
}

