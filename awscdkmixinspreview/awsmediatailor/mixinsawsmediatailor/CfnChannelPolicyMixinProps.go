package mixinsawsmediatailor


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
//   	ChannelName: jsii.String("channelName"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channelpolicy.html
//
type CfnChannelPolicyMixinProps struct {
	// The name of the channel associated with this Channel Policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channelpolicy.html#cfn-mediatailor-channelpolicy-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The IAM policy for the channel.
	//
	// IAM policies are used to control access to your channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channelpolicy.html#cfn-mediatailor-channelpolicy-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

