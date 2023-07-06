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
type CfnTransitGatewayVpcAttachment_OptionsProperty struct {
	// Enable or disable appliance mode support.
	//
	// The default is `disable` .
	ApplianceModeSupport *string `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Enable or disable DNS support.
	//
	// The default is `disable` .
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Enable or disable IPv6 support.
	//
	// The default is `disable` .
	Ipv6Support *string `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

