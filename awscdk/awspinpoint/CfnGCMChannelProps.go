package awspinpoint


// Properties for defining a `CfnGCMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGCMChannelProps := &CfnGCMChannelProps{
//   	ApiKey: jsii.String("apiKey"),
//   	ApplicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html
//
type CfnGCMChannelProps struct {
	// The Web API key, also called the *server key* , that you received from Google to communicate with Google services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-apikey
	//
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The unique identifier for the Amazon Pinpoint application that the GCM channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the GCM channel for the Amazon Pinpoint application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

