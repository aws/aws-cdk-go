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
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage awsstepfunctions.QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	// Default: - None (Task is executed using the State Machine's execution role).
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Default: - None.
	//
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Default: - None.
	//
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: $.
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to select part of the state to be the output to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: $.
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// Used to specify and transform output from the state.
	//
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: $.
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
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
	// Default: ActionOnFailure.CONTINUE
	//
	ActionOnFailure ActionOnFailure `field:"optional" json:"actionOnFailure" yaml:"actionOnFailure"`
	// A list of command line arguments passed to the JAR file's main function when executed.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Default: - No args.
	//
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The Amazon Resource Name (ARN) of the runtime role for a step on the cluster.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_AddJobFlowSteps.html#API_AddJobFlowSteps_RequestSyntax
	//
	// Default: - Uses EC2 instance profile role.
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Default: - No mainClass.
	//
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key value pairs to your main function.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Default: - No properties.
	//
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

