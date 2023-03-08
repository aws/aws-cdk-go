package awsssmcontacts


// Properties for defining a `CfnContactChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContactChannelProps := &cfnContactChannelProps{
//   	channelAddress: jsii.String("channelAddress"),
//   	channelName: jsii.String("channelName"),
//   	channelType: jsii.String("channelType"),
//   	contactId: jsii.String("contactId"),
//
//   	// the properties below are optional
//   	deferActivation: jsii.Boolean(false),
//   }
//
type CfnContactChannelProps struct {
	// The details that Incident Manager uses when trying to engage the contact channel.
	ChannelAddress *string `field:"required" json:"channelAddress" yaml:"channelAddress"`
	// The name of the contact channel.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The type of the contact channel. Incident Manager supports three contact methods:.
	//
	// - SMS
	// - VOICE
	// - EMAIL.
	ChannelType *string `field:"required" json:"channelType" yaml:"channelType"`
	// The Amazon Resource Name (ARN) of the contact you are adding the contact channel to.
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
	// If you want to activate the channel at a later time, you can choose to defer activation.
	//
	// Incident Manager can't engage your contact channel until it has been activated.
	DeferActivation interface{} `field:"optional" json:"deferActivation" yaml:"deferActivation"`
}

