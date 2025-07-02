package awsmediatailor


// Properties for defining a `CfnChannelPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnChannelPolicyProps := &CfnChannelPolicyProps{
//   	ChannelName: jsii.String("channelName"),
//   	Policy: policy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channelpolicy.html
//
type CfnChannelPolicyProps struct {
	// The name of the channel associated with this Channel Policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channelpolicy.html#cfn-mediatailor-channelpolicy-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The IAM policy for the channel.
	//
	// IAM policies are used to control access to your channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channelpolicy.html#cfn-mediatailor-channelpolicy-policy
	//
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

