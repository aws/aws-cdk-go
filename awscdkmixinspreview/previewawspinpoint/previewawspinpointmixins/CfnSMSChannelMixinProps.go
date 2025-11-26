package previewawspinpointmixins


// Properties for CfnSMSChannelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSMSChannelMixinProps := &CfnSMSChannelMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	Enabled: jsii.Boolean(false),
//   	SenderId: jsii.String("senderId"),
//   	ShortCode: jsii.String("shortCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html
//
type CfnSMSChannelMixinProps struct {
	// The unique identifier for the Amazon Pinpoint application that the SMS channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html#cfn-pinpoint-smschannel-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the SMS channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html#cfn-pinpoint-smschannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The identity that you want to display on recipients' devices when they receive messages from the SMS channel.
	//
	// > SenderIDs are only supported in certain countries and regions. For more information, see [Supported Countries and Regions](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) in the *Amazon Pinpoint User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html#cfn-pinpoint-smschannel-senderid
	//
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// The registered short code that you want to use when you send messages through the SMS channel.
	//
	// > For information about obtaining a dedicated short code for sending SMS messages, see [Requesting Dedicated Short Codes for SMS Messaging with Amazon Pinpoint](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-short-code.html) in the *Amazon Pinpoint User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-smschannel.html#cfn-pinpoint-smschannel-shortcode
	//
	ShortCode *string `field:"optional" json:"shortCode" yaml:"shortCode"`
}

