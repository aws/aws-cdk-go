package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for EmrAddStep.
//
// Example:
//   tasks.NewEmrAddStep(this, jsii.String("Task"), &EmrAddStepProps{
//   	ClusterId: jsii.String("ClusterId"),
//   	Name: jsii.String("StepName"),
//   	Jar: jsii.String("Jar"),
//   	ActionOnFailure: tasks.ActionOnFailure_CONTINUE,
//   })
//
type EmrAddStepProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The ClusterId to add the Step to.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// A path to a JAR file run during the step.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	Jar *string `field:"required" json:"jar" yaml:"jar"`
	// The name of the Step.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_StepConfig.html
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to take when the cluster step fails.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_StepConfig.html
	//
	ActionOnFailure ActionOnFailure `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
	// A list of command line arguments passed to the JAR file's main function when executed.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key value pairs to your main function.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

