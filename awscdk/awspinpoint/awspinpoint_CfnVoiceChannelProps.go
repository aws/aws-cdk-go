package awspinpoint


// Properties for defining a `CfnVoiceChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVoiceChannelProps := &cfnVoiceChannelProps{
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnVoiceChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the voice channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the voice channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

