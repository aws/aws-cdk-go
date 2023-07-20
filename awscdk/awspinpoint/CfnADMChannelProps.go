package awspinpoint


// Properties for defining a `CfnADMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnADMChannelProps := &CfnADMChannelProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ClientId: jsii.String("clientId"),
//   	ClientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html
//
type CfnADMChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the ADM channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Client ID that you received from Amazon to send messages by using ADM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Client Secret that you received from Amazon to send messages by using ADM.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-clientsecret
	//
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Specifies whether to enable the ADM channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-admchannel.html#cfn-pinpoint-admchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

