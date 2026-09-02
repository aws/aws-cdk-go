package awsmsk


// Includes information about the channel state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   channelStateInfoProperty := &ChannelStateInfoProperty{
//   	Code: jsii.String("code"),
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channelstateinfo.html
//
type CfnChannel_ChannelStateInfoProperty struct {
	// Code for channel state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channelstateinfo.html#cfn-msk-channel-channelstateinfo-code
	//
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Message for channel state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-channelstateinfo.html#cfn-msk-channel-channelstateinfo-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

