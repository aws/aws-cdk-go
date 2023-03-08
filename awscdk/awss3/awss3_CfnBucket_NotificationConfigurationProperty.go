package awss3


// Describes the notification configuration for an Amazon S3 bucket.
//
// > If you create the target resource and related permissions in the same template, you might have a circular dependency.
// >
// > For example, you might use the `AWS::Lambda::Permission` resource to grant the bucket permission to invoke an AWS Lambda function. However, AWS CloudFormation can't create the bucket until the bucket has permission to invoke the function ( AWS CloudFormation checks whether the bucket can invoke the function). If you're using Refs to pass the bucket name, this leads to a circular dependency.
// >
// > To avoid this dependency, you can create all resources without specifying the notification configuration. Then, update the stack with a notification configuration.
// >
// > For more information on permissions, see [AWS::Lambda::Permission](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html) and [Granting Permissions to Publish Event Notification Messages to a Destination](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#grant-destinations-permissions-to-s3) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &notificationConfigurationProperty{
//   	eventBridgeConfiguration: &eventBridgeConfigurationProperty{
//   		eventBridgeEnabled: jsii.Boolean(false),
//   	},
//   	lambdaConfigurations: []interface{}{
//   		&lambdaConfigurationProperty{
//   			event: jsii.String("event"),
//   			function: jsii.String("function"),
//
//   			// the properties below are optional
//   			filter: &notificationFilterProperty{
//   				s3Key: &s3KeyFilterProperty{
//   					rules: []interface{}{
//   						&filterRuleProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	queueConfigurations: []interface{}{
//   		&queueConfigurationProperty{
//   			event: jsii.String("event"),
//   			queue: jsii.String("queue"),
//
//   			// the properties below are optional
//   			filter: &notificationFilterProperty{
//   				s3Key: &s3KeyFilterProperty{
//   					rules: []interface{}{
//   						&filterRuleProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	topicConfigurations: []interface{}{
//   		&topicConfigurationProperty{
//   			event: jsii.String("event"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			filter: &notificationFilterProperty{
//   				s3Key: &s3KeyFilterProperty{
//   					rules: []interface{}{
//   						&filterRuleProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_NotificationConfigurationProperty struct {
	// Enables delivery of events to Amazon EventBridge.
	EventBridgeConfiguration interface{} `field:"optional" json:"eventBridgeConfiguration" yaml:"eventBridgeConfiguration"`
	// Describes the AWS Lambda functions to invoke and the events for which to invoke them.
	LambdaConfigurations interface{} `field:"optional" json:"lambdaConfigurations" yaml:"lambdaConfigurations"`
	// The Amazon Simple Queue Service queues to publish messages to and the events for which to publish messages.
	QueueConfigurations interface{} `field:"optional" json:"queueConfigurations" yaml:"queueConfigurations"`
	// The topic to which notifications are sent and the events for which notifications are generated.
	TopicConfigurations interface{} `field:"optional" json:"topicConfigurations" yaml:"topicConfigurations"`
}

