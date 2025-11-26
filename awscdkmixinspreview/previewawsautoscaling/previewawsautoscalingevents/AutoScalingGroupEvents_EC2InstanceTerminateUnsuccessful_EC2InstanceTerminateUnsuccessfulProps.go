package previewawsautoscalingevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for AutoScalingGroup aws.autoscaling@EC2InstanceTerminateUnsuccessful event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2InstanceTerminateUnsuccessfulProps := &EC2InstanceTerminateUnsuccessfulProps{
//   	ActivityId: []*string{
//   		jsii.String("activityId"),
//   	},
//   	AutoScalingGroupName: []*string{
//   		jsii.String("autoScalingGroupName"),
//   	},
//   	Cause: []*string{
//   		jsii.String("cause"),
//   	},
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	Destination: []*string{
//   		jsii.String("destination"),
//   	},
//   	Details: &Details{
//   		AvailabilityZone: []*string{
//   			jsii.String("availabilityZone"),
//   		},
//   		SubnetId: []*string{
//   			jsii.String("subnetId"),
//   		},
//   	},
//   	Ec2InstanceId: []*string{
//   		jsii.String("ec2InstanceId"),
//   	},
//   	EndTime: []*string{
//   		jsii.String("endTime"),
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
//   	Origin: []*string{
//   		jsii.String("origin"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	StatusCode: []*string{
//   		jsii.String("statusCode"),
//   	},
//   	StatusMessage: []*string{
//   		jsii.String("statusMessage"),
//   	},
//   }
//
// Experimental.
type AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps struct {
	// ActivityId property.
	//
	// Specify an array of string values to match this event if the actual value of ActivityId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ActivityId *[]*string `field:"optional" json:"activityId" yaml:"activityId"`
	// AutoScalingGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of AutoScalingGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the AutoScalingGroup reference.
	//
	// Experimental.
	AutoScalingGroupName *[]*string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Cause property.
	//
	// Specify an array of string values to match this event if the actual value of Cause is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cause *[]*string `field:"optional" json:"cause" yaml:"cause"`
	// Description property.
	//
	// Specify an array of string values to match this event if the actual value of Description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// Destination property.
	//
	// Specify an array of string values to match this event if the actual value of Destination is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Destination *[]*string `field:"optional" json:"destination" yaml:"destination"`
	// Details property.
	//
	// Specify an array of string values to match this event if the actual value of Details is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Details *AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_Details `field:"optional" json:"details" yaml:"details"`
	// EC2InstanceId property.
	//
	// Specify an array of string values to match this event if the actual value of EC2InstanceId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ec2InstanceId *[]*string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// EndTime property.
	//
	// Specify an array of string values to match this event if the actual value of EndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTime *[]*string `field:"optional" json:"endTime" yaml:"endTime"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// Origin property.
	//
	// Specify an array of string values to match this event if the actual value of Origin is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Origin *[]*string `field:"optional" json:"origin" yaml:"origin"`
	// RequestId property.
	//
	// Specify an array of string values to match this event if the actual value of RequestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// StartTime property.
	//
	// Specify an array of string values to match this event if the actual value of StartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// StatusCode property.
	//
	// Specify an array of string values to match this event if the actual value of StatusCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusCode *[]*string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// StatusMessage property.
	//
	// Specify an array of string values to match this event if the actual value of StatusMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StatusMessage *[]*string `field:"optional" json:"statusMessage" yaml:"statusMessage"`
}

