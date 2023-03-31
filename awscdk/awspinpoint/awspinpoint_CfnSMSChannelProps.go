package awspinpoint


// Properties for defining a `CfnSMSChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSMSChannelProps := &cfnSMSChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	senderId: jsii.String("senderId"),
//   	shortCode: jsii.String("shortCode"),
//   }
//
type CfnSMSChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the SMS channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the SMS channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The identity that you want to display on recipients' devices when they receive messages from the SMS channel.
	//
	// > SenderIDs are only supported in certain countries and regions. For more information, see [Supported Countries and Regions](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-countries.html) in the *Amazon Pinpoint User Guide* .
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// The registered short code that you want to use when you send messages through the SMS channel.
	//
	// > For information about obtaining a dedicated short code for sending SMS messages, see [Requesting Dedicated Short Codes for SMS Messaging with Amazon Pinpoint](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-short-code.html) in the *Amazon Pinpoint User Guide* .
	ShortCode *string `field:"optional" json:"shortCode" yaml:"shortCode"`
}

