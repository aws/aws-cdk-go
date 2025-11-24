package mixinsawssmsvoice


// The phone number's two-way SMS configuration object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   twoWayProperty := &TwoWayProperty{
//   	ChannelArn: jsii.String("channelArn"),
//   	ChannelRole: jsii.String("channelRole"),
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html
//
type CfnPhoneNumberPropsMixin_TwoWayProperty struct {
	// The Amazon Resource Name (ARN) of the two way channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html#cfn-smsvoice-phonenumber-twoway-channelarn
	//
	ChannelArn *string `field:"optional" json:"channelArn" yaml:"channelArn"`
	// An optional IAM Role Arn for a service to assume, to be able to post inbound SMS messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html#cfn-smsvoice-phonenumber-twoway-channelrole
	//
	ChannelRole *string `field:"optional" json:"channelRole" yaml:"channelRole"`
	// By default this is set to false.
	//
	// When set to true you can receive incoming text messages from your end recipients using the TwoWayChannelArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-smsvoice-phonenumber-twoway.html#cfn-smsvoice-phonenumber-twoway-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

