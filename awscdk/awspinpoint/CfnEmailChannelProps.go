package awspinpoint


// Properties for defining a `CfnEmailChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailChannelProps := &CfnEmailChannelProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	FromAddress: jsii.String("fromAddress"),
//   	Identity: jsii.String("identity"),
//
//   	// the properties below are optional
//   	ConfigurationSet: jsii.String("configurationSet"),
//   	Enabled: jsii.Boolean(false),
//   	OrchestrationSendingRoleArn: jsii.String("orchestrationSendingRoleArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html
//
type CfnEmailChannelProps struct {
	// The unique identifier for the Amazon Pinpoint application that you're specifying the email channel for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The verified email address that you want to send email from when you send email through the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-fromaddress
	//
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
	// The Amazon Resource Name (ARN) of the identity, verified with Amazon Simple Email Service (Amazon SES), that you want to use when you send email through the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-identity
	//
	Identity *string `field:"required" json:"identity" yaml:"identity"`
	// The [Amazon SES configuration set](https://docs.aws.amazon.com/ses/latest/APIReference/API_ConfigurationSet.html) that you want to apply to messages that you send through the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-configurationset
	//
	ConfigurationSet *string `field:"optional" json:"configurationSet" yaml:"configurationSet"`
	// Specifies whether to enable the email channel for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of an IAM role for Amazon Pinpoint to use to send email from your campaigns or journeys through Amazon SES .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-orchestrationsendingrolearn
	//
	OrchestrationSendingRoleArn *string `field:"optional" json:"orchestrationSendingRoleArn" yaml:"orchestrationSendingRoleArn"`
	// The ARN of the AWS Identity and Access Management (IAM) role that you want Amazon Pinpoint to use when it submits email-related event data for the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-emailchannel.html#cfn-pinpoint-emailchannel-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

