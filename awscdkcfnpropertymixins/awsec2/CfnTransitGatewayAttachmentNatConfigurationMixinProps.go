package awsec2


// Properties for CfnTransitGatewayAttachmentNatConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTransitGatewayAttachmentNatConfigurationMixinProps := &CfnTransitGatewayAttachmentNatConfigurationMixinProps{
//   	NatGatewayRules: []interface{}{
//   		&NatRuleProperty{
//   			MatchSourceCidr: jsii.String("matchSourceCidr"),
//   			NatCidr: jsii.String("natCidr"),
//   			RuleNumber: jsii.Number(123),
//   		},
//   	},
//   	NatGatewaySupport: jsii.String("natGatewaySupport"),
//   	NatSystemRouteSupport: jsii.String("natSystemRouteSupport"),
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachmentnatconfiguration.html
//
type CfnTransitGatewayAttachmentNatConfigurationMixinProps struct {
	// The complete set of NAT rules for the attachment.
	//
	// Updating this property replaces the entire rule set: rules omitted from the new value are removed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachmentnatconfiguration.html#cfn-ec2-transitgatewayattachmentnatconfiguration-natgatewayrules
	//
	NatGatewayRules interface{} `field:"optional" json:"natGatewayRules" yaml:"natGatewayRules"`
	// Whether NAT gateway support is enabled on the attachment.
	//
	// Set to disable to stage NAT rules without activating them; the rules are retained and become active when this is later set to enable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachmentnatconfiguration.html#cfn-ec2-transitgatewayattachmentnatconfiguration-natgatewaysupport
	//
	NatGatewaySupport *string `field:"optional" json:"natGatewaySupport" yaml:"natGatewaySupport"`
	// Whether routes for the NAT CIDRs are installed into the transit gateway route tables associated with the attachment.
	//
	// Whether this is configurable depends on the attachment; where the value is forced, supplying the opposite value is rejected. Omitting this property leaves the setting to the service: on a new configuration the service applies the attachment's default, but on an attachment where the value is configurable it keeps the value already in effect rather than reverting to that default. Removing this property from a template, or deleting this resource, therefore leaves the setting at its last value rather than restoring the default, and a resource later created on the same attachment without this property inherits that value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachmentnatconfiguration.html#cfn-ec2-transitgatewayattachmentnatconfiguration-natsystemroutesupport
	//
	NatSystemRouteSupport *string `field:"optional" json:"natSystemRouteSupport" yaml:"natSystemRouteSupport"`
	// The ID of the transit gateway attachment to configure NAT on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachmentnatconfiguration.html#cfn-ec2-transitgatewayattachmentnatconfiguration-transitgatewayattachmentid
	//
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}

