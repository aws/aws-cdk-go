package awsec2


// The state of VPC Block Public Access (BPA).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   blockPublicAccessStatesProperty := &BlockPublicAccessStatesProperty{
//   	InternetGatewayBlockMode: jsii.String("internetGatewayBlockMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-subnet-blockpublicaccessstates.html
//
type CfnSubnet_BlockPublicAccessStatesProperty struct {
	// The mode of VPC BPA.
	//
	// - `off` : VPC BPA is not enabled and traffic is allowed to and from internet gateways and egress-only internet gateways in this Region.
	// - `block-bidirectional` : Block all traffic to and from internet gateways and egress-only internet gateways in this Region (except for excluded VPCs and subnets).
	// - `block-ingress` : Block all internet traffic to the VPCs in this Region (except for VPCs or subnets which are excluded). Only traffic to and from NAT gateways and egress-only internet gateways is allowed because these gateways only allow outbound connections to be established.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-subnet-blockpublicaccessstates.html#cfn-ec2-subnet-blockpublicaccessstates-internetgatewayblockmode
	//
	InternetGatewayBlockMode *string `field:"optional" json:"internetGatewayBlockMode" yaml:"internetGatewayBlockMode"`
}

