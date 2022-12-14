package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &optionsProperty{
//   	autoAcceptSharedAssociations: jsii.String("autoAcceptSharedAssociations"),
//   	igmpv2Support: jsii.String("igmpv2Support"),
//   	staticSourcesSupport: jsii.String("staticSourcesSupport"),
//   }
//
type CfnTransitGatewayMulticastDomain_OptionsProperty struct {
	// `CfnTransitGatewayMulticastDomain.OptionsProperty.AutoAcceptSharedAssociations`.
	AutoAcceptSharedAssociations *string `field:"optional" json:"autoAcceptSharedAssociations" yaml:"autoAcceptSharedAssociations"`
	// `CfnTransitGatewayMulticastDomain.OptionsProperty.Igmpv2Support`.
	Igmpv2Support *string `field:"optional" json:"igmpv2Support" yaml:"igmpv2Support"`
	// `CfnTransitGatewayMulticastDomain.OptionsProperty.StaticSourcesSupport`.
	StaticSourcesSupport *string `field:"optional" json:"staticSourcesSupport" yaml:"staticSourcesSupport"`
}

