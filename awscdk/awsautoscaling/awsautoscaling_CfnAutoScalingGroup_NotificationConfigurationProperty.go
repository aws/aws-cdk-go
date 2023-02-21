package awsautoscaling


// A structure that specifies an Amazon SNS notification configuration for the `NotificationConfigurations` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource.
//
// For an example template snippet, see [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html) .
//
// For more information, see [Get Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &notificationConfigurationProperty{
//   	topicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	notificationTypes: []*string{
//   		jsii.String("notificationTypes"),
//   	},
//   }
//
type CfnAutoScalingGroup_NotificationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// A list of event types that send a notification. Event types can include any of the following types.
	//
	// *Allowed values* :
	//
	// - `autoscaling:EC2_INSTANCE_LAUNCH`
	// - `autoscaling:EC2_INSTANCE_LAUNCH_ERROR`
	// - `autoscaling:EC2_INSTANCE_TERMINATE`
	// - `autoscaling:EC2_INSTANCE_TERMINATE_ERROR`
	// - `autoscaling:TEST_NOTIFICATION`.
	NotificationTypes *[]*string `field:"optional" json:"notificationTypes" yaml:"notificationTypes"`
}

