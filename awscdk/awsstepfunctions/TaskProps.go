package awsstepfunctions

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Props that are common to all tasks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var parameters interface{}
//   var stepFunctionsTask iStepFunctionsTask
//
//   taskProps := &TaskProps{
//   	Task: stepFunctionsTask,
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   	InputPath: jsii.String("inputPath"),
//   	OutputPath: jsii.String("outputPath"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	ResultPath: jsii.String("resultPath"),
//   	Timeout: duration,
//   }
//
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
type TaskProps struct {
	// Actual task to be invoked in this workflow.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Task IStepFunctionsTask `field:"required" json:"task" yaml:"task"`
	// An optional description for this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Parameters to invoke the task with.
	//
	// It is not recommended to use this field. The object that is passed in
	// the `task` property will take care of returning the right values for the
	// `Parameters` field in the Step Functions definition.
	//
	// The various classes that implement `IStepFunctionsTask` will take a
	// properties which make sense for the task type. For example, for
	// `InvokeFunction` the field that populates the `parameters` field will be
	// called `payload`, and for the `PublishToTopic` the `parameters` field
	// will be populated via a combination of the referenced topic, subject and
	// message.
	//
	// If passed anyway, the keys in this map will override the parameters
	// returned by the task object.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-parameters
	//
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// Maximum run time of this state.
	//
	// If the state takes longer than this amount of time to complete, a 'Timeout' error is raised.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

