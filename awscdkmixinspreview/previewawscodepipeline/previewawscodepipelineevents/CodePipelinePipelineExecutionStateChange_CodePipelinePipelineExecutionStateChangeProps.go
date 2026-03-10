package previewawscodepipelineevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.codepipeline@CodePipelinePipelineExecutionStateChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codePipelinePipelineExecutionStateChangeProps := &CodePipelinePipelineExecutionStateChangeProps{
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
//   	ExecutionTrigger: &ExecutionTrigger{
//   		TriggerDetail: []*string{
//   			jsii.String("triggerDetail"),
//   		},
//   		TriggerType: []*string{
//   			jsii.String("triggerType"),
//   		},
//   	},
//   	Pipeline: []*string{
//   		jsii.String("pipeline"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   	Version: []*string{
//   		jsii.String("version"),
//   	},
//   }
//
// Experimental.
type CodePipelinePipelineExecutionStateChange_CodePipelinePipelineExecutionStateChangeProps struct {
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
	// execution-trigger property.
	//
	// Specify an array of string values to match this event if the actual value of execution-trigger is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionTrigger *CodePipelinePipelineExecutionStateChange_ExecutionTrigger `field:"optional" json:"executionTrigger" yaml:"executionTrigger"`
	// pipeline property.
	//
	// Specify an array of string values to match this event if the actual value of pipeline is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Pipeline *[]*string `field:"optional" json:"pipeline" yaml:"pipeline"`
	// state property.
	//
	// Specify an array of string values to match this event if the actual value of state is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
	// version property.
	//
	// Specify an array of string values to match this event if the actual value of version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Version *[]*string `field:"optional" json:"version" yaml:"version"`
}

