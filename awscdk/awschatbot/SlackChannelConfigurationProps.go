package awschatbot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a new Slack channel configuration.
//
// Example:
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//   var project Project
//
//
//   target := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &SlackChannelConfigurationProps{
//   	SlackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	SlackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	SlackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := project.notifyOnBuildSucceeded(jsii.String("NotifyOnBuildSucceeded"), target)
//
type SlackChannelConfigurationProps struct {
	// The name of Slack channel configuration.
	SlackChannelConfigurationName *string `field:"required" json:"slackChannelConfigurationName" yaml:"slackChannelConfigurationName"`
	// The ID of the Slack channel.
	//
	// To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link.
	// The channel ID is the 9-character string at the end of the URL. For example, ABCBBLZZZ.
	SlackChannelId *string `field:"required" json:"slackChannelId" yaml:"slackChannelId"`
	// The ID of the Slack workspace authorized with AWS Chatbot.
	//
	// To get the workspace ID, you must perform the initial authorization flow with Slack in the AWS Chatbot console.
	// Then you can copy and paste the workspace ID from the console.
	// For more details, see steps 1-4 in Setting Up AWS Chatbot with Slack in the AWS Chatbot User Guide.
	// See: https://docs.aws.amazon.com/chatbot/latest/adminguide/setting-up.html#Setup_intro
	//
	SlackWorkspaceId *string `field:"required" json:"slackWorkspaceId" yaml:"slackWorkspaceId"`
	// A list of IAM managed policies that are applied as channel guardrails.
	// Default: - The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	//
	GuardrailPolicies *[]awsiam.IManagedPolicy `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// Specifies the logging level for this configuration.
	//
	// This property affects the log entries pushed to Amazon CloudWatch Logs.
	// Default: LoggingLevel.NONE
	//
	LoggingLevel LoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Default: logs.RetentionDays.INFINITE
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	// Default: - Default AWS SDK retry options.
	//
	LogRetentionRetryOptions *awslogs.LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Default: - A new role is created.
	//
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The SNS topics that deliver notifications to AWS Chatbot.
	// Default: None.
	//
	NotificationTopics *[]awssns.ITopic `field:"optional" json:"notificationTopics" yaml:"notificationTopics"`
	// The permission role of Slack channel configuration.
	// Default: - A role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Enables use of a user role requirement in your chat configuration.
	// Default: false.
	//
	UserRoleRequired *bool `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

