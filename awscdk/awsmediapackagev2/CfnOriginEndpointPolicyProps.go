package awsmediapackagev2


// Properties for defining a `CfnOriginEndpointPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnOriginEndpointPolicyProps := &CfnOriginEndpointPolicyProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	OriginEndpointName: jsii.String("originEndpointName"),
//   	Policy: policy,
//
//   	// the properties below are optional
//   	CdnAuthConfiguration: &CdnAuthConfigurationProperty{
//   		CdnIdentifierSecretArns: []*string{
//   			jsii.String("cdnIdentifierSecretArns"),
//   		},
//   		SecretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html
//
type CfnOriginEndpointPolicyProps struct {
	// The name of the channel group associated with the origin endpoint policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-channelgroupname
	//
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The channel name associated with the origin endpoint policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The name of the origin endpoint associated with the origin endpoint policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-originendpointname
	//
	OriginEndpointName *string `field:"required" json:"originEndpointName" yaml:"originEndpointName"`
	// The policy associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// The settings to enable CDN authorization headers in MediaPackage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpointpolicy.html#cfn-mediapackagev2-originendpointpolicy-cdnauthconfiguration
	//
	CdnAuthConfiguration interface{} `field:"optional" json:"cdnAuthConfiguration" yaml:"cdnAuthConfiguration"`
}

