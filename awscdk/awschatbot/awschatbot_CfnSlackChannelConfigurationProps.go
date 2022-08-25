package awschatbot


// Properties for defining a `CfnSlackChannelConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSlackChannelConfigurationProps := &cfnSlackChannelConfigurationProps{
//   	configurationName: jsii.String("configurationName"),
//   	iamRoleArn: jsii.String("iamRoleArn"),
//   	slackChannelId: jsii.String("slackChannelId"),
//   	slackWorkspaceId: jsii.String("slackWorkspaceId"),
//
//   	// the properties below are optional
//   	guardrailPolicies: []*string{
//   		jsii.String("guardrailPolicies"),
//   	},
//   	loggingLevel: jsii.String("loggingLevel"),
//   	snsTopicArns: []*string{
//   		jsii.String("snsTopicArns"),
//   	},
//   	userRoleRequired: jsii.Boolean(false),
//   }
//
type CfnSlackChannelConfigurationProps struct {
	// The name of the configuration.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot .
	//
	// This is a user-definworked role that AWS Chatbot will assume. This is not the service-linked role. For more information, see [IAM Policies for AWS Chatbot](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html) .
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The ID of the Slack channel.
	//
	// To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link. The channel ID is the 9-character string at the end of the URL. For example, `ABCBBLZZZ` .
	SlackChannelId *string `field:"required" json:"slackChannelId" yaml:"slackChannelId"`
	// The ID of the Slack workspace authorized with AWS Chatbot .
	//
	// To get the workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console. Then you can copy and paste the workspace ID from the console. For more details, see steps 1-4 in [Setting Up AWS Chatbot with Slack](https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro) in the *AWS Chatbot User Guide* .
	SlackWorkspaceId *string `field:"required" json:"slackWorkspaceId" yaml:"slackWorkspaceId"`
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies *[]*string `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Logging levels include `ERROR` , `INFO` , or `NONE` .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot .
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Enables use of a user role requirement in your chat configuration.
	UserRoleRequired interface{} `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

