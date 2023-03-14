package awsec2


// The options for the transit gateway multicast domain.
//
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
	// Indicates whether to automatically accept cross-account subnet associations that are associated with the transit gateway multicast domain.
	AutoAcceptSharedAssociations *string `field:"optional" json:"autoAcceptSharedAssociations" yaml:"autoAcceptSharedAssociations"`
	// Specify whether to enable Internet Group Management Protocol (IGMP) version 2 for the transit gateway multicast domain.
	Igmpv2Support *string `field:"optional" json:"igmpv2Support" yaml:"igmpv2Support"`
	// Specify whether to enable support for statically configuring multicast group sources for a domain.
	StaticSourcesSupport *string `field:"optional" json:"staticSourcesSupport" yaml:"staticSourcesSupport"`
}

