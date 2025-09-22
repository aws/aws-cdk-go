package awsec2


// Describes the VPC attachment options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &OptionsProperty{
//   	ApplianceModeSupport: jsii.String("applianceModeSupport"),
//   	DnsSupport: jsii.String("dnsSupport"),
//   	Ipv6Support: jsii.String("ipv6Support"),
//   	SecurityGroupReferencingSupport: jsii.String("securityGroupReferencingSupport"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachment-options.html
//
type CfnTransitGatewayAttachment_OptionsProperty struct {
	// Enable or disable appliance mode support.
	//
	// The default is `disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachment-options.html#cfn-ec2-transitgatewayattachment-options-appliancemodesupport
	//
	ApplianceModeSupport *string `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Enable or disable DNS support.
	//
	// The default is `disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachment-options.html#cfn-ec2-transitgatewayattachment-options-dnssupport
	//
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Enable or disable IPv6 support.
	//
	// The default is `disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachment-options.html#cfn-ec2-transitgatewayattachment-options-ipv6support
	//
	Ipv6Support *string `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
	// Enables you to reference a security group across VPCs attached to a transit gateway (TGW).
	//
	// Use this option to simplify security group management and control of instance-to-instance traffic across VPCs that are connected by transit gateway. You can also use this option to migrate from VPC peering (which was the only option that supported security group referencing) to transit gateways (which now also support security group referencing). This option is disabled by default and there are no additional costs to use this feature.
	//
	// For important information about this feature, see [Create a transit gateway](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-transit-gateways.html#create-tgw) in the *AWS Transit Gateway Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayattachment-options.html#cfn-ec2-transitgatewayattachment-options-securitygroupreferencingsupport
	//
	SecurityGroupReferencingSupport *string `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
}

