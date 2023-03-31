package awspinpoint


// Properties for defining a `CfnBaiduChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBaiduChannelProps := &cfnBaiduChannelProps{
//   	apiKey: jsii.String("apiKey"),
//   	applicationId: jsii.String("applicationId"),
//   	secretKey: jsii.String("secretKey"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBaiduChannelProps struct {
	// The API key that you received from the Baidu Cloud Push service to communicate with the service.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The unique identifier for the Amazon Pinpoint application that you're configuring the Baidu channel for.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The secret key that you received from the Baidu Cloud Push service to communicate with the service.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Specifies whether to enable the Baidu channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

