package awscodestarnotifications


// The level of detail to include in the notifications for this resource.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import chatbot "github.com/aws/aws-cdk-go/awscdk"
//
//
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic1"))
//
//   slack := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &SlackChannelConfigurationProps{
//   	SlackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
//   	SlackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
//   	SlackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
//   })
//
//   rule := notifications.NewNotificationRule(this, jsii.String("NotificationRule"), &NotificationRuleProps{
//   	Source: project,
//   	Events: []*string{
//   		jsii.String("codebuild-project-build-state-succeeded"),
//   		jsii.String("codebuild-project-build-state-failed"),
//   	},
//   	Targets: []iNotificationRuleTarget{
//   		topic,
//   	},
//   	NotificationRuleName: jsii.String("MyNotificationRuleName"),
//   	Enabled: jsii.Boolean(true),
//   	 // The default is true
//   	DetailType: notifications.DetailType_FULL,
//   	 // The default is FULL
//   	CreatedBy: jsii.String("Jone Doe"),
//   })
//   rule.AddTarget(slack)
//
type DetailType string

const (
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	DetailType_BASIC DetailType = "BASIC"
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType_FULL DetailType = "FULL"
)

