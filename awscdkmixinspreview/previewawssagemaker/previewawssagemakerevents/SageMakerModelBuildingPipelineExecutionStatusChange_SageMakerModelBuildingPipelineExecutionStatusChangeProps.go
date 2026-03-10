package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerModelBuildingPipelineExecutionStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelBuildingPipelineExecutionStatusChangeProps := &SageMakerModelBuildingPipelineExecutionStatusChangeProps{
//   	CurrentPipelineExecutionStatus: []*string{
//   		jsii.String("currentPipelineExecutionStatus"),
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
//   	ExecutionEndTime: []*string{
//   		jsii.String("executionEndTime"),
//   	},
//   	ExecutionStartTime: []*string{
//   		jsii.String("executionStartTime"),
//   	},
//   	PipelineArn: []*string{
//   		jsii.String("pipelineArn"),
//   	},
//   	PipelineExecutionArn: []*string{
//   		jsii.String("pipelineExecutionArn"),
//   	},
//   	PipelineExecutionDescription: []*string{
//   		jsii.String("pipelineExecutionDescription"),
//   	},
//   	PipelineExecutionDisplayName: []*string{
//   		jsii.String("pipelineExecutionDisplayName"),
//   	},
//   	PreviousPipelineExecutionStatus: []*string{
//   		jsii.String("previousPipelineExecutionStatus"),
//   	},
//   }
//
// Experimental.
type SageMakerModelBuildingPipelineExecutionStatusChange_SageMakerModelBuildingPipelineExecutionStatusChangeProps struct {
	// currentPipelineExecutionStatus property.
	//
	// Specify an array of string values to match this event if the actual value of currentPipelineExecutionStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentPipelineExecutionStatus *[]*string `field:"optional" json:"currentPipelineExecutionStatus" yaml:"currentPipelineExecutionStatus"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// executionEndTime property.
	//
	// Specify an array of string values to match this event if the actual value of executionEndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionEndTime *[]*string `field:"optional" json:"executionEndTime" yaml:"executionEndTime"`
	// executionStartTime property.
	//
	// Specify an array of string values to match this event if the actual value of executionStartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionStartTime *[]*string `field:"optional" json:"executionStartTime" yaml:"executionStartTime"`
	// pipelineArn property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineArn *[]*string `field:"optional" json:"pipelineArn" yaml:"pipelineArn"`
	// pipelineExecutionArn property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineExecutionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineExecutionArn *[]*string `field:"optional" json:"pipelineExecutionArn" yaml:"pipelineExecutionArn"`
	// pipelineExecutionDescription property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineExecutionDescription is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineExecutionDescription *[]*string `field:"optional" json:"pipelineExecutionDescription" yaml:"pipelineExecutionDescription"`
	// pipelineExecutionDisplayName property.
	//
	// Specify an array of string values to match this event if the actual value of pipelineExecutionDisplayName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PipelineExecutionDisplayName *[]*string `field:"optional" json:"pipelineExecutionDisplayName" yaml:"pipelineExecutionDisplayName"`
	// previousPipelineExecutionStatus property.
	//
	// Specify an array of string values to match this event if the actual value of previousPipelineExecutionStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousPipelineExecutionStatus *[]*string `field:"optional" json:"previousPipelineExecutionStatus" yaml:"previousPipelineExecutionStatus"`
}

