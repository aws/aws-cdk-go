package awspinpoint


// Properties for defining a `CfnADMChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnADMChannelProps := &cfnADMChannelProps{
//   	applicationId: jsii.String("applicationId"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnADMChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that the ADM channel applies to.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The Client ID that you received from Amazon to send messages by using ADM.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Client Secret that you received from Amazon to send messages by using ADM.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Specifies whether to enable the ADM channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

