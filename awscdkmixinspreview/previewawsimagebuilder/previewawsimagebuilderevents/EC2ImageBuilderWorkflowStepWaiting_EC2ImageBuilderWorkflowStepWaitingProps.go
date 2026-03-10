package previewawsimagebuilderevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.imagebuilder@EC2ImageBuilderWorkflowStepWaiting event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eC2ImageBuilderWorkflowStepWaitingProps := &EC2ImageBuilderWorkflowStepWaitingProps{
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
//   	WorkflowExecutionId: []*string{
//   		jsii.String("workflowExecutionId"),
//   	},
//   	WorkflowStepExecutionId: []*string{
//   		jsii.String("workflowStepExecutionId"),
//   	},
//   	WorkflowStepName: []*string{
//   		jsii.String("workflowStepName"),
//   	},
//   }
//
// Experimental.
type EC2ImageBuilderWorkflowStepWaiting_EC2ImageBuilderWorkflowStepWaitingProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// workflow-execution-id property.
	//
	// Specify an array of string values to match this event if the actual value of workflow-execution-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WorkflowExecutionId *[]*string `field:"optional" json:"workflowExecutionId" yaml:"workflowExecutionId"`
	// workflow-step-execution-id property.
	//
	// Specify an array of string values to match this event if the actual value of workflow-step-execution-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WorkflowStepExecutionId *[]*string `field:"optional" json:"workflowStepExecutionId" yaml:"workflowStepExecutionId"`
	// workflow-step-name property.
	//
	// Specify an array of string values to match this event if the actual value of workflow-step-name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	WorkflowStepName *[]*string `field:"optional" json:"workflowStepName" yaml:"workflowStepName"`
}

