# AWS CodeStarNotifications Construct Library

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## NotificationRule

The `NotificationRule` construct defines an AWS CodeStarNotifications rule.
The rule specifies the events you want notifications about and the targets
(such as Amazon SNS topics or AWS Chatbot clients configured for Slack)
where you want to receive them.
Notification targets are objects that implement the `INotificationRuleTarget`
interface and notification source is object that implement the `INotificationRuleSource` interface.

## Notification Targets

This module includes classes that implement the `INotificationRuleTarget` interface for SNS and slack in AWS Chatbot.

The following targets are supported:

* `SNS`: specify event and notify to SNS topic.
* `AWS Chatbot`: specify event and notify to slack channel and only support `SlackChannelConfiguration`.

## Examples

```go
import notifications "github.com/aws/aws-cdk-go/awscdk"
import codebuild "github.com/aws/aws-cdk-go/awscdk"
import sns "github.com/aws/aws-cdk-go/awscdk"
import chatbot "github.com/aws/aws-cdk-go/awscdk"


project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))

topic := sns.NewTopic(this, jsii.String("MyTopic1"))

slack := chatbot.NewSlackChannelConfiguration(this, jsii.String("MySlackChannel"), &slackChannelConfigurationProps{
	slackChannelConfigurationName: jsii.String("YOUR_CHANNEL_NAME"),
	slackWorkspaceId: jsii.String("YOUR_SLACK_WORKSPACE_ID"),
	slackChannelId: jsii.String("YOUR_SLACK_CHANNEL_ID"),
})

rule := notifications.NewNotificationRule(this, jsii.String("NotificationRule"), &notificationRuleProps{
	source: project,
	events: []*string{
		jsii.String("codebuild-project-build-state-succeeded"),
		jsii.String("codebuild-project-build-state-failed"),
	},
	targets: []iNotificationRuleTarget{
		topic,
	},
})
rule.addTarget(slack)
```

## Notification Source

This module includes classes that implement the `INotificationRuleSource` interface for AWS CodeBuild,
AWS CodePipeline and will support AWS CodeCommit, AWS CodeDeploy in future.

The following sources are supported:

* `AWS CodeBuild`: support codebuild project to trigger notification when event specified.
* `AWS CodePipeline`: support codepipeline to trigger notification when event specified.

## Events

For the complete list of supported event types for CodeBuild and CodePipeline, see:

* [Events for notification rules on build projects](https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-buildproject).
* [Events for notification rules on pipelines](https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-pipeline).
