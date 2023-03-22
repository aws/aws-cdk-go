package awspinpoint


// Properties for defining a `CfnEmailChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailChannelProps := &cfnEmailChannelProps{
//   	applicationId: jsii.String("applicationId"),
//   	fromAddress: jsii.String("fromAddress"),
//   	identity: jsii.String("identity"),
//
//   	// the properties below are optional
//   	configurationSet: jsii.String("configurationSet"),
//   	enabled: jsii.Boolean(false),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnEmailChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that you're specifying the email channel for.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The verified email address that you want to send email from when you send email through the channel.
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
	// The Amazon Resource Name (ARN) of the identity, verified with Amazon Simple Email Service (Amazon SES), that you want to use when you send email through the channel.
	Identity *string `field:"required" json:"identity" yaml:"identity"`
	// The [Amazon SES configuration set](https://docs.aws.amazon.com/ses/latest/APIReference/API_ConfigurationSet.html) that you want to apply to messages that you send through the channel.
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Specifies whether to enable the email channel for the application.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the AWS Identity and Access Management (IAM) role that you want Amazon Pinpoint to use when it submits email-related event data for the channel.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

