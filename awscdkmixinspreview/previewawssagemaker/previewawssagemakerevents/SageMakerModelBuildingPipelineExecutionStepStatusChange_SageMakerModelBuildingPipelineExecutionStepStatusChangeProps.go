package previewawssagemakerevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.sagemaker@SageMakerModelBuildingPipelineExecutionStepStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sageMakerModelBuildingPipelineExecutionStepStatusChangeProps := &SageMakerModelBuildingPipelineExecutionStepStatusChangeProps{
//   	CacheHitResult: &CacheHitResult{
//   		SourcePipelineExecutionArn: []*string{
//   			jsii.String("sourcePipelineExecutionArn"),
//   		},
//   	},
//   	CurrentStepStatus: []*string{
//   		jsii.String("currentStepStatus"),
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
//   	FailureReason: []*string{
//   		jsii.String("failureReason"),
//   	},
//   	Metadata: &Metadata{
//   		ProcessingJob: &ProcessingJob{
//   			Arn: []*string{
//   				jsii.String("arn"),
//   			},
//   		},
//   	},
//   	PipelineArn: []*string{
//   		jsii.String("pipelineArn"),
//   	},
//   	PipelineExecutionArn: []*string{
//   		jsii.String("pipelineExecutionArn"),
//   	},
//   	PreviousStepStatus: []*string{
//   		jsii.String("previousStepStatus"),
//   	},
//   	StepEndTime: []*string{
//   		jsii.String("stepEndTime"),
//   	},
//   	StepName: []*string{
//   		jsii.String("stepName"),
//   	},
//   	StepStartTime: []*string{
//   		jsii.String("stepStartTime"),
//   	},
//   	StepType: []*string{
//   		jsii.String("stepType"),
//   	},
//   }
//
// Experimental.
type SageMakerModelBuildingPipelineExecutionStepStatusChange_SageMakerModelBuildingPipelineExecutionStepStatusChangeProps struct {
	// cacheHitResult property.
	//
	// Specify an array of string values to match this event if the actual value of cacheHitResult is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CacheHitResult *SageMakerModelBuildingPipelineExecutionStepStatusChange_CacheHitResult `field:"optional" json:"cacheHitResult" yaml:"cacheHitResult"`
	// currentStepStatus property.
	//
	// Specify an array of string values to match this event if the actual value of currentStepStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CurrentStepStatus *[]*string `field:"optional" json:"currentStepStatus" yaml:"currentStepStatus"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// failureReason property.
	//
	// Specify an array of string values to match this event if the actual value of failureReason is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureReason *[]*string `field:"optional" json:"failureReason" yaml:"failureReason"`
	// metadata property.
	//
	// Specify an array of string values to match this event if the actual value of metadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metadata *SageMakerModelBuildingPipelineExecutionStepStatusChange_Metadata `field:"optional" json:"metadata" yaml:"metadata"`
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
	// previousStepStatus property.
	//
	// Specify an array of string values to match this event if the actual value of previousStepStatus is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PreviousStepStatus *[]*string `field:"optional" json:"previousStepStatus" yaml:"previousStepStatus"`
	// stepEndTime property.
	//
	// Specify an array of string values to match this event if the actual value of stepEndTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StepEndTime *[]*string `field:"optional" json:"stepEndTime" yaml:"stepEndTime"`
	// stepName property.
	//
	// Specify an array of string values to match this event if the actual value of stepName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StepName *[]*string `field:"optional" json:"stepName" yaml:"stepName"`
	// stepStartTime property.
	//
	// Specify an array of string values to match this event if the actual value of stepStartTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StepStartTime *[]*string `field:"optional" json:"stepStartTime" yaml:"stepStartTime"`
	// stepType property.
	//
	// Specify an array of string values to match this event if the actual value of stepType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StepType *[]*string `field:"optional" json:"stepType" yaml:"stepType"`
}

