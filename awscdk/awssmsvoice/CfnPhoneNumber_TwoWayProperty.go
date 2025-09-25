package awssmsvoice


// The phone number's two-way SMS configuration object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   twoWayProperty := &TwoWayProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	ChannelArn: jsii.String("channelArn"),
//   	ChannelRole: jsii.String("channelRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html
//
type CfnPhoneNumber_TwoWayProperty struct {
	// By default this is set to false.
	//
	// When set to true you can receive incoming text messages from your end recipients using the TwoWayChannelArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html#cfn-smsvoice-phonenumber-twoway-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Amazon Resource Name (ARN) of the two way channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html#cfn-smsvoice-phonenumber-twoway-channelarn
	//
	ChannelArn *string `field:"optional" json:"channelArn" yaml:"channelArn"`
	// An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html#cfn-smsvoice-phonenumber-twoway-channelrole
	//
	ChannelRole *string `field:"optional" json:"channelRole" yaml:"channelRole"`
}

