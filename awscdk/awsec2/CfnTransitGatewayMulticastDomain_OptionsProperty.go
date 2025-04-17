package awsec2


// The options for the transit gateway multicast domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionsProperty := &OptionsProperty{
//   	AutoAcceptSharedAssociations: jsii.String("autoAcceptSharedAssociations"),
//   	Igmpv2Support: jsii.String("igmpv2Support"),
//   	StaticSourcesSupport: jsii.String("staticSourcesSupport"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaymulticastdomain-options.html
//
type CfnTransitGatewayMulticastDomain_OptionsProperty struct {
	// Indicates whether to automatically accept cross-account subnet associations that are associated with the transit gateway multicast domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaymulticastdomain-options.html#cfn-ec2-transitgatewaymulticastdomain-options-autoacceptsharedassociations
	//
	AutoAcceptSharedAssociations *string `field:"optional" json:"autoAcceptSharedAssociations" yaml:"autoAcceptSharedAssociations"`
	// Specify whether to enable Internet Group Management Protocol (IGMP) version 2 for the transit gateway multicast domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaymulticastdomain-options.html#cfn-ec2-transitgatewaymulticastdomain-options-igmpv2support
	//
	Igmpv2Support *string `field:"optional" json:"igmpv2Support" yaml:"igmpv2Support"`
	// Specify whether to enable support for statically configuring multicast group sources for a domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-transitgatewaymulticastdomain-options.html#cfn-ec2-transitgatewaymulticastdomain-options-staticsourcessupport
	//
	StaticSourcesSupport *string `field:"optional" json:"staticSourcesSupport" yaml:"staticSourcesSupport"`
}

