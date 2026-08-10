package previewawsstatesevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.states@StepFunctionsExecutionStatusChange event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stepFunctionsExecutionStatusChangeProps := &StepFunctionsExecutionStatusChangeProps{
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
//   	ExecutionArn: []*string{
//   		jsii.String("executionArn"),
//   	},
//   	Input: []*string{
//   		jsii.String("input"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	Output: []*string{
//   		jsii.String("output"),
//   	},
//   	StartDate: []*string{
//   		jsii.String("startDate"),
//   	},
//   	StateMachineArn: []*string{
//   		jsii.String("stateMachineArn"),
//   	},
//   	Status: []*string{
//   		jsii.String("status"),
//   	},
//   	StopDate: []*string{
//   		jsii.String("stopDate"),
//   	},
//   }
//
// Experimental.
type StepFunctionsExecutionStatusChange_StepFunctionsExecutionStatusChangeProps struct {
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// executionArn property.
	//
	// Specify an array of string values to match this event if the actual value of executionArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Filter with the Execution reference.
	//
	// Experimental.
	ExecutionArn *[]*string `field:"optional" json:"executionArn" yaml:"executionArn"`
	// input property.
	//
	// Specify an array of string values to match this event if the actual value of input is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Input *[]*string `field:"optional" json:"input" yaml:"input"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// output property.
	//
	// Specify an array of string values to match this event if the actual value of output is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Output *[]*string `field:"optional" json:"output" yaml:"output"`
	// startDate property.
	//
	// Specify an array of string values to match this event if the actual value of startDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartDate *[]*string `field:"optional" json:"startDate" yaml:"startDate"`
	// stateMachineArn property.
	//
	// Specify an array of string values to match this event if the actual value of stateMachineArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StateMachineArn *[]*string `field:"optional" json:"stateMachineArn" yaml:"stateMachineArn"`
	// status property.
	//
	// Specify an array of string values to match this event if the actual value of status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Status *[]*string `field:"optional" json:"status" yaml:"status"`
	// stopDate property.
	//
	// Specify an array of string values to match this event if the actual value of stopDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StopDate *[]*string `field:"optional" json:"stopDate" yaml:"stopDate"`
}

