package awspinpoint


// Properties for defining a `CfnGCMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGCMChannelProps := &CfnGCMChannelProps{
//   	ApplicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	ApiKey: jsii.String("apiKey"),
//   	DefaultAuthenticationMethod: jsii.String("defaultAuthenticationMethod"),
//   	Enabled: jsii.Boolean(false),
//   	ServiceJson: jsii.String("serviceJson"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html
//
type CfnGCMChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the GCM channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Web API key, also called the *server key* , that you received from Google to communicate with Google services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-apikey
	//
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-defaultauthenticationmethod
	//
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Specifies whether to enable the GCM channel for the Amazon Pinpoint application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-gcmchannel.html#cfn-pinpoint-gcmchannel-servicejson
	//
	ServiceJson *string `field:"optional" json:"serviceJson" yaml:"serviceJson"`
}

