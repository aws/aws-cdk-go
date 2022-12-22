package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties that define what kind of task should be created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var metricDimensions interface{}
//   var parameters interface{}
//   var policyStatement policyStatement
//
//   stepFunctionsTaskConfig := &stepFunctionsTaskConfig{
//   	resourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	heartbeat: duration,
//   	metricDimensions: map[string]interface{}{
//   		"metricDimensionsKey": metricDimensions,
//   	},
//   	metricPrefixPlural: jsii.String("metricPrefixPlural"),
//   	metricPrefixSingular: jsii.String("metricPrefixSingular"),
//   	parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	policyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   }
//
// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
type StepFunctionsTaskConfig struct {
	// The resource that represents the work to be executed.
	//
	// Either the ARN of a Lambda Function or Activity, or a special
	// ARN.
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Maximum time between heart beats.
	//
	// If the time between heart beats takes longer than this, a 'Timeout' error is raised.
	//
	// This is only relevant when using an Activity type as resource.
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// The dimensions to attach to metrics.
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	MetricDimensions *map[string]interface{} `field:"optional" json:"metricDimensions" yaml:"metricDimensions"`
	// Prefix for plural metric names of activity actions.
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	MetricPrefixPlural *string `field:"optional" json:"metricPrefixPlural" yaml:"metricPrefixPlural"`
	// Prefix for singular metric names of activity actions.
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	MetricPrefixSingular *string `field:"optional" json:"metricPrefixSingular" yaml:"metricPrefixSingular"`
	// Parameters pass a collection of key-value pairs, either static values or JSONPath expressions that select from the input.
	//
	// The meaning of these parameters is task-dependent.
	//
	// Its values will be merged with the `parameters` property which is configured directly
	// on the Task state.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters
	//
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Additional policy statements to add to the execution role.
	// Deprecated: used by `IStepFunctionsTask`. `IStepFunctionsTask` is deprecated and replaced by `TaskStateBase`.
	PolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"policyStatements" yaml:"policyStatements"`
}

