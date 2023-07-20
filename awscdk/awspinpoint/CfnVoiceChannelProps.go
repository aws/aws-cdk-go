package awspinpoint


// Properties for defining a `CfnVoiceChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVoiceChannelProps := &CfnVoiceChannelProps{
//   	ApplicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-voicechannel.html
//
type CfnVoiceChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the voice channel applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-voicechannel.html#cfn-pinpoint-voicechannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the voice channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-voicechannel.html#cfn-pinpoint-voicechannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

