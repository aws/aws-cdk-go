package previewawscodepipelineevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codepipeline@CodePipelineActionExecutionStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codePipelineActionExecutionStateChangeProps := &CodePipelineActionExecutionStateChangeProps{
//   	Action: []*string{
//   		jsii.String("action"),
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
//   	ExecutionId: []*string{
//   		jsii.String("executionId"),
//   	},
//   	ExecutionResult: &ExecutionResult{
//   		ExternalExecutionId: []*string{
//   			jsii.String("externalExecutionId"),
//   		},
//   		ExternalExecutionSummary: []*string{
//   			jsii.String("externalExecutionSummary"),
//   		},
//   		ExternalExecutionUrl: []*string{
//   			jsii.String("externalExecutionUrl"),
//   		},
//   	},
//   	InputArtifacts: &InputArtifacts{
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		S3Location: &S3Location{
//   			Bucket: []*string{
//   				jsii.String("bucket"),
//   			},
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   		},
//   	},
//   	OutputArtifacts: &OutputArtifacts{
//   		Name: []*string{
//   			jsii.String("name"),
//   		},
//   		S3Location: &S3Location{
//   			Bucket: []*string{
//   				jsii.String("bucket"),
//   			},
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   		},
//   	},
//   	Pipeline: []*string{
//   		jsii.String("pipeline"),
//   	},
//   	Region: []*string{
//   		jsii.String("region"),
//   	},
//   	Stage: []*string{
//   		jsii.String("stage"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	Type: &Type{
//   		Category: []*string{
//   			jsii.String("category"),
//   		},
//   		Owner: []*string{
//   			jsii.String("owner"),
//   		},
//   		Provider: []*string{
//   			jsii.String("provider"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type CodePipelineActionExecutionStateChange_CodePipelineActionExecutionStateChangeProps struct {
	// action property.
	//
	// Specify an array of string values to match this event if the actual value of action is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Action *[]*string `field:"optional" json:"action" yaml:"action"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// execution-id property.
	//
	// Specify an array of string values to match this event if the actual value of execution-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionId *[]*string `field:"optional" json:"executionId" yaml:"executionId"`
	// execution-result property.
	//
	// Specify an array of string values to match this event if the actual value of execution-result is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionResult *CodePipelineActionExecutionStateChange_ExecutionResult `field:"optional" json:"executionResult" yaml:"executionResult"`
	// input-artifacts property.
	//
	// Specify an array of string values to match this event if the actual value of input-artifacts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputArtifacts *CodePipelineActionExecutionStateChange_InputArtifacts `field:"optional" json:"inputArtifacts" yaml:"inputArtifacts"`
	// output-artifacts property.
	//
	// Specify an array of string values to match this event if the actual value of output-artifacts is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OutputArtifacts *CodePipelineActionExecutionStateChange_OutputArtifacts `field:"optional" json:"outputArtifacts" yaml:"outputArtifacts"`
	// pipeline property.
	//
	// Specify an array of string values to match this event if the actual value of pipeline is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Pipeline *[]*string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// region property.
	//
	// Specify an array of string values to match this event if the actual value of region is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Region *[]*string `field:"optional" json:"region" yaml:"region"`
	// stage property.
	//
	// Specify an array of string values to match this event if the actual value of stage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Stage *[]*string `field:"optional" json:"stage" yaml:"stage"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// type property.
	//
	// Specify an array of string values to match this event if the actual value of type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Type *CodePipelineActionExecutionStateChange_Type `field:"optional" json:"type" yaml:"type"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

