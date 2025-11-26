package previewawsautoscalingevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for AutoScalingGroup aws.autoscaling@EC2InstanceLaunchLifecycleAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceLaunchLifecycleActionProps := &EC2InstanceLaunchLifecycleActionProps{
//   	AutoScalingGroupName: []*string{
//   		jsii.String("autoScalingGroupName"),
//   	},
//   	Destination: []*string{
//   		jsii.String("destination"),
//   	},
//   	Ec2InstanceId: []*string{
//   		jsii.String("ec2InstanceId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	LifecycleActionToken: []*string{
//   		jsii.String("lifecycleActionToken"),
//   	},
//   	LifecycleHookName: []*string{
//   		jsii.String("lifecycleHookName"),
//   	},
//   	LifecycleTransition: []*string{
//   		jsii.String("lifecycleTransition"),
//   	},
//   	NotificationMetadata: []*string{
//   		jsii.String("notificationMetadata"),
//   	},
//   	Origin: []*string{
//   		jsii.String("origin"),
//   	},
//   }
//
// Experimental.
type AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps struct {
	// AutoScalingGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of AutoScalingGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the AutoScalingGroup reference.
	//
	// Experimental.
	AutoScalingGroupName *[]*string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Destination property.
	//
	// Specify an array of string values to match this event if the actual value of Destination is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Destination *[]*string `field:"optional" json:"destination" yaml:"destination"`
	// EC2InstanceId property.
	//
	// Specify an array of string values to match this event if the actual value of EC2InstanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ec2InstanceId *[]*string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// LifecycleActionToken property.
	//
	// Specify an array of string values to match this event if the actual value of LifecycleActionToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LifecycleActionToken *[]*string `field:"optional" json:"lifecycleActionToken" yaml:"lifecycleActionToken"`
	// LifecycleHookName property.
	//
	// Specify an array of string values to match this event if the actual value of LifecycleHookName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LifecycleHookName *[]*string `field:"optional" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// LifecycleTransition property.
	//
	// Specify an array of string values to match this event if the actual value of LifecycleTransition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LifecycleTransition *[]*string `field:"optional" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// NotificationMetadata property.
	//
	// Specify an array of string values to match this event if the actual value of NotificationMetadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationMetadata *[]*string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// Origin property.
	//
	// Specify an array of string values to match this event if the actual value of Origin is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Origin *[]*string `field:"optional" json:"origin" yaml:"origin"`
}

