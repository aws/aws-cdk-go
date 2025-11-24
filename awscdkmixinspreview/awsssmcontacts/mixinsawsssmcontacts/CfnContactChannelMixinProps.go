package mixinsawsssmcontacts


// Properties for CfnContactChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactChannelMixinProps := &CfnContactChannelMixinProps{
//   	ChannelAddress: jsii.String("channelAddress"),
//   	ChannelName: jsii.String("channelName"),
//   	ChannelType: jsii.String("channelType"),
//   	ContactId: jsii.String("contactId"),
//   	DeferActivation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contactchannel.html
//
type CfnContactChannelMixinProps struct {
	// The details that Incident Manager uses when trying to engage the contact channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contactchannel.html#cfn-ssmcontacts-contactchannel-channeladdress
	//
	ChannelAddress *string `field:"optional" json:"channelAddress" yaml:"channelAddress"`
	// The name of the contact channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contactchannel.html#cfn-ssmcontacts-contactchannel-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The type of the contact channel. Incident Manager supports three contact methods:.
	//
	// - SMS
	// - VOICE
	// - EMAIL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contactchannel.html#cfn-ssmcontacts-contactchannel-channeltype
	//
	ChannelType *string `field:"optional" json:"channelType" yaml:"channelType"`
	// The Amazon Resource Name (ARN) of the contact you are adding the contact channel to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contactchannel.html#cfn-ssmcontacts-contactchannel-contactid
	//
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
	// If you want to activate the channel at a later time, you can choose to defer activation.
	//
	// Incident Manager can't engage your contact channel until it has been activated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contactchannel.html#cfn-ssmcontacts-contactchannel-deferactivation
	//
	DeferActivation interface{} `field:"optional" json:"deferActivation" yaml:"deferActivation"`
}

