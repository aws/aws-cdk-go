package awspinpoint


// Properties for defining a `CfnGCMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGCMChannelProps := &cfnGCMChannelProps{
//   	apiKey: jsii.String("apiKey"),
//   	applicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnGCMChannelProps struct {
	// The Web API key, also called the *server key* , that you received from Google to communicate with Google services.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The unique identifier for the Amazon Pinpoint application that the GCM channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Specifies whether to enable the GCM channel for the Amazon Pinpoint application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

