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
type CfnTransitGatewayAttachment_OptionsProperty struct {
	// Indicates whether appliance mode support is enabled.
	ApplianceModeSupport *string `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Indicates whether DNS support is enabled.
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Indicates whether IPv6 support is disabled.
	Ipv6Support *string `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

