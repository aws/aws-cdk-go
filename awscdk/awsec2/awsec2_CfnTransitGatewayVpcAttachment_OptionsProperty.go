package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &optionsProperty{
//   	applianceModeSupport: jsii.String("applianceModeSupport"),
//   	dnsSupport: jsii.String("dnsSupport"),
//   	ipv6Support: jsii.String("ipv6Support"),
//   }
//
type CfnTransitGatewayVpcAttachment_OptionsProperty struct {
	// `CfnTransitGatewayVpcAttachment.OptionsProperty.ApplianceModeSupport`.
	ApplianceModeSupport *string `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// `CfnTransitGatewayVpcAttachment.OptionsProperty.DnsSupport`.
	DnsSupport *string `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// `CfnTransitGatewayVpcAttachment.OptionsProperty.Ipv6Support`.
	Ipv6Support *string `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

