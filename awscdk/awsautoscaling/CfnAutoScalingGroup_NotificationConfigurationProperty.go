package awsautoscaling


// A structure that specifies an Amazon SNS notification configuration for the `NotificationConfigurations` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.
//
// For an example template snippet, see [Configure Amazon EC2 Auto Scaling resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-auto-scaling.html) .
//
// For more information, see [Get Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &NotificationConfigurationProperty{
//   	TopicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	NotificationTypes: []*string{
//   		jsii.String("notificationTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-notificationconfiguration.html
//
type CfnAutoScalingGroup_NotificationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-notificationconfiguration.html#cfn-autoscaling-autoscalinggroup-notificationconfiguration-topicarn
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-notificationconfiguration.html#cfn-autoscaling-autoscalinggroup-notificationconfiguration-notificationtypes
	//
	NotificationTypes *[]*string `field:"optional" json:"notificationTypes" yaml:"notificationTypes"`
}

