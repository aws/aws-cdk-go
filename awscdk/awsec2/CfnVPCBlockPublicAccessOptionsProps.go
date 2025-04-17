package awsec2


// Properties for defining a `CfnVPCBlockPublicAccessOptions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCBlockPublicAccessOptionsProps := &CfnVPCBlockPublicAccessOptionsProps{
//   	InternetGatewayBlockMode: jsii.String("internetGatewayBlockMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessoptions.html
//
type CfnVPCBlockPublicAccessOptionsProps struct {
	// The desired VPC Block Public Access mode for internet gateways in your account.
	//
	// We do not allow you to create this resource type in an "off" mode since off is the default value.
	//
	// - `block-bidirectional` : Block all traffic to and from internet gateways and egress-only internet gateways in this Region (except for excluded VPCs and subnets).
	// - `block-ingress` : Block all internet traffic to the VPCs in this Region (except for VPCs or subnets which are excluded). Only traffic to and from NAT gateways and egress-only internet gateways is allowed because these gateways only allow outbound connections to be established.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessoptions.html#cfn-ec2-vpcblockpublicaccessoptions-internetgatewayblockmode
	//
	InternetGatewayBlockMode *string `field:"required" json:"internetGatewayBlockMode" yaml:"internetGatewayBlockMode"`
}

