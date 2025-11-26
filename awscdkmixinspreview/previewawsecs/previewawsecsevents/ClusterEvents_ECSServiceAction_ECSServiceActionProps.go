package previewawsecsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Cluster aws.ecs@ECSServiceAction event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSServiceActionProps := &ECSServiceActionProps{
//   	CapacityProviderArns: []*string{
//   		jsii.String("capacityProviderArns"),
//   	},
//   	ClusterArn: []*string{
//   		jsii.String("clusterArn"),
//   	},
//   	ContainerInstanceArns: []*string{
//   		jsii.String("containerInstanceArns"),
//   	},
//   	ContainerPort: []*string{
//   		jsii.String("containerPort"),
//   	},
//   	CreatedAt: []*string{
//   		jsii.String("createdAt"),
//   	},
//   	DesiredCount: []*string{
//   		jsii.String("desiredCount"),
//   	},
//   	Ec2InstanceIds: []*string{
//   		jsii.String("ec2InstanceIds"),
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
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	Reason: []*string{
//   		jsii.String("reason"),
//   	},
//   	ServiceRegistryArns: []*string{
//   		jsii.String("serviceRegistryArns"),
//   	},
//   	TargetGroupArns: []*string{
//   		jsii.String("targetGroupArns"),
//   	},
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//   	TaskArns: []*string{
//   		jsii.String("taskArns"),
//   	},
//   	TaskSetArns: []*string{
//   		jsii.String("taskSetArns"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSServiceAction_ECSServiceActionProps struct {
	// capacityProviderArns property.
	//
	// Specify an array of string values to match this event if the actual value of capacityProviderArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CapacityProviderArns *[]*string `field:"optional" json:"capacityProviderArns" yaml:"capacityProviderArns"`
	// clusterArn property.
	//
	// Specify an array of string values to match this event if the actual value of clusterArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Cluster reference.
	//
	// Experimental.
	ClusterArn *[]*string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// containerInstanceArns property.
	//
	// Specify an array of string values to match this event if the actual value of containerInstanceArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerInstanceArns *[]*string `field:"optional" json:"containerInstanceArns" yaml:"containerInstanceArns"`
	// containerPort property.
	//
	// Specify an array of string values to match this event if the actual value of containerPort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerPort *[]*string `field:"optional" json:"containerPort" yaml:"containerPort"`
	// createdAt property.
	//
	// Specify an array of string values to match this event if the actual value of createdAt is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreatedAt *[]*string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// desiredCount property.
	//
	// Specify an array of string values to match this event if the actual value of desiredCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DesiredCount *[]*string `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// ec2InstanceIds property.
	//
	// Specify an array of string values to match this event if the actual value of ec2InstanceIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Ec2InstanceIds *[]*string `field:"optional" json:"ec2InstanceIds" yaml:"ec2InstanceIds"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// reason property.
	//
	// Specify an array of string values to match this event if the actual value of reason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
	// serviceRegistryArns property.
	//
	// Specify an array of string values to match this event if the actual value of serviceRegistryArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceRegistryArns *[]*string `field:"optional" json:"serviceRegistryArns" yaml:"serviceRegistryArns"`
	// targetGroupArns property.
	//
	// Specify an array of string values to match this event if the actual value of targetGroupArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// targets property.
	//
	// Specify an array of string values to match this event if the actual value of targets is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Targets *[]*string `field:"optional" json:"targets" yaml:"targets"`
	// taskArns property.
	//
	// Specify an array of string values to match this event if the actual value of taskArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskArns *[]*string `field:"optional" json:"taskArns" yaml:"taskArns"`
	// taskSetArns property.
	//
	// Specify an array of string values to match this event if the actual value of taskSetArns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TaskSetArns *[]*string `field:"optional" json:"taskSetArns" yaml:"taskSetArns"`
}

