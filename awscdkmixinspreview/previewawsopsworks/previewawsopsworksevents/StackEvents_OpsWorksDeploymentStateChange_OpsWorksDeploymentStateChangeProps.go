package previewawsopsworksevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for Stack aws.opsworks@OpsWorksDeploymentStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   opsWorksDeploymentStateChangeProps := &OpsWorksDeploymentStateChangeProps{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	DeploymentId: []*string{
//   		jsii.String("deploymentId"),
//   	},
//   	Duration: []*string{
//   		jsii.String("duration"),
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
//   	InstanceIds: []*string{
//   		jsii.String("instanceIds"),
//   	},
//   	StackId: []*string{
//   		jsii.String("stackId"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   }
//
// Experimental.
type StackEvents_OpsWorksDeploymentStateChange_OpsWorksDeploymentStateChangeProps struct {
	// command property.
	//
	// Specify an array of string values to match this event if the actual value of command is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// deployment-id property.
	//
	// Specify an array of string values to match this event if the actual value of deployment-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DeploymentId *[]*string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// duration property.
	//
	// Specify an array of string values to match this event if the actual value of duration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Duration *[]*string `field:"optional" json:"duration" yaml:"duration"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// instance-ids property.
	//
	// Specify an array of string values to match this event if the actual value of instance-ids is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// stack-id property.
	//
	// Specify an array of string values to match this event if the actual value of stack-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Stack reference.
	//
	// Experimental.
	StackId *[]*string `field:"optional" json:"stackId" yaml:"stackId"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
}

