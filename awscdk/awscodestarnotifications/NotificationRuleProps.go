package awscodestarnotifications


// Properties for a new notification rule.
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
type NotificationRuleProps struct {
	// The name or email alias of the person who created the notification rule.
	//
	// If not specified, it means that the creator's alias is not provided.
	// Default: - No alias provided.
	//
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	// Default: DetailType.FULL
	//
	DetailType DetailType `field:"optional" json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	// Default: - generated from the `id`.
	//
	NotificationRuleName *string `field:"optional" json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// The Amazon Resource Name (ARN) of the resource to associate with the notification rule.
	//
	// Currently, Supported sources include pipelines in AWS CodePipeline, build projects in AWS CodeBuild, and repositories in AWS CodeCommit in this L2 constructor.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html#cfn-codestarnotifications-notificationrule-resource
	//
	Source INotificationRuleSource `field:"required" json:"source" yaml:"source"`
	// The targets to register for the notification destination.
	// Default: - No targets are added to the rule. Use `addTarget()` to add a target.
	//
	Targets *[]INotificationRuleTarget `field:"optional" json:"targets" yaml:"targets"`
}

