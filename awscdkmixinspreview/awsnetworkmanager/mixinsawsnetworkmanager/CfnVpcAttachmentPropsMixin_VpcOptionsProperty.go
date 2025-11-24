package mixinsawsnetworkmanager


// Describes the VPC options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcOptionsProperty := &VpcOptionsProperty{
//   	ApplianceModeSupport: jsii.Boolean(false),
//   	DnsSupport: jsii.Boolean(false),
//   	Ipv6Support: jsii.Boolean(false),
//   	SecurityGroupReferencingSupport: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-vpcoptions.html
//
type CfnVpcAttachmentPropsMixin_VpcOptionsProperty struct {
	// Indicates whether appliance mode is supported.
	//
	// If enabled, traffic flow between a source and destination use the same Availability Zone for the VPC attachment for the lifetime of that flow. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-vpcoptions.html#cfn-networkmanager-vpcattachment-vpcoptions-appliancemodesupport
	//
	// Default: - false.
	//
	ApplianceModeSupport interface{} `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Indicates whether DNS is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-vpcoptions.html#cfn-networkmanager-vpcattachment-vpcoptions-dnssupport
	//
	// Default: - true.
	//
	DnsSupport interface{} `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Indicates whether IPv6 is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-vpcoptions.html#cfn-networkmanager-vpcattachment-vpcoptions-ipv6support
	//
	// Default: - false.
	//
	Ipv6Support interface{} `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
	// Indicates whether security group referencing is enabled for this VPC attachment.
	//
	// The default is `true` . However, at the core network policy-level the default is set to `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-vpcoptions.html#cfn-networkmanager-vpcattachment-vpcoptions-securitygroupreferencingsupport
	//
	// Default: - true.
	//
	SecurityGroupReferencingSupport interface{} `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
}

