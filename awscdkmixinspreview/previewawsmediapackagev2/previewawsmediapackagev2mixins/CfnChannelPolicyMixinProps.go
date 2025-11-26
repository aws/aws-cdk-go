package previewawsmediapackagev2mixins


// Properties for CfnChannelPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policy interface{}
//
//   cfnChannelPolicyMixinProps := &CfnChannelPolicyMixinProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelpolicy.html
//
type CfnChannelPolicyMixinProps struct {
	// The name of the channel group associated with the channel policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelpolicy.html#cfn-mediapackagev2-channelpolicy-channelgroupname
	//
	ChannelGroupName *string `field:"optional" json:"channelGroupName" yaml:"channelGroupName"`
	// The name of the channel associated with the channel policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelpolicy.html#cfn-mediapackagev2-channelpolicy-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The policy associated with the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelpolicy.html#cfn-mediapackagev2-channelpolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

