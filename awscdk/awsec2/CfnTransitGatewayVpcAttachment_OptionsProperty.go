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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayvpcattachment-options.html
//
type CfnTransitGatewayVpcAttachment_OptionsProperty struct {
	// Enable or disable appliance mode support.
	//
	// The default is `disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayvpcattachment-options.html#cfn-ec2-transitgatewayvpcattachment-options-appliancemodesupport
	//
	ApplianceModeSupport *string `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Enable or disable DNS support.
	//
	// The default is `disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayvpcattachment-options.html#cfn-ec2-transitgatewayvpcattachment-options-dnssupport
	//
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Enable or disable IPv6 support.
	//
	// The default is `disable` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewayvpcattachment-options.html#cfn-ec2-transitgatewayvpcattachment-options-ipv6support
	//
	Ipv6Support *string `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

