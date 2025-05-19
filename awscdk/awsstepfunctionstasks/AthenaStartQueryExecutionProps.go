package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for starting a Query Execution.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &AthenaStartQueryExecutionProps{
//   	QueryString: sfn.JsonPath_StringAt(jsii.String("$.queryString")),
//   	QueryExecutionContext: &QueryExecutionContext{
//   		DatabaseName: jsii.String("mydatabase"),
//   	},
//   	ResultConfiguration: &ResultConfiguration{
//   		EncryptionConfiguration: &EncryptionConfiguration{
//   			EncryptionOption: tasks.EncryptionOption_S3_MANAGED,
//   		},
//   		OutputLocation: &Location{
//   			BucketName: jsii.String("amzn-s3-demo-bucket"),
//   			ObjectKey: jsii.String("folder"),
//   		},
//   	},
//   	ExecutionParameters: []*string{
//   		jsii.String("param1"),
//   		jsii.String("param2"),
//   	},
//   })
//
type AthenaStartQueryExecutionProps struct {
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
	// Query that will be started.
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Unique string string to ensure idempotence.
	// Default: - No client request token.
	//
	ClientRequestToken *string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
	// A list of values for the parameters in a query.
	//
	// The values are applied sequentially to the parameters in the query in the order
	// in which the parameters occur.
	// Default: - No parameters.
	//
	ExecutionParameters *[]*string `field:"optional" json:"executionParameters" yaml:"executionParameters"`
	// Database within which query executes.
	// Default: - No query execution context.
	//
	QueryExecutionContext *QueryExecutionContext `field:"optional" json:"queryExecutionContext" yaml:"queryExecutionContext"`
	// Configuration on how and where to save query.
	// Default: - No result configuration.
	//
	ResultConfiguration *ResultConfiguration `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
	// Specifies, in minutes, the maximum age of a previous query result that Athena should consider for reuse.
	// Default: - Query results are not reused.
	//
	ResultReuseConfigurationMaxAge awscdk.Duration `field:"optional" json:"resultReuseConfigurationMaxAge" yaml:"resultReuseConfigurationMaxAge"`
	// Configuration on how and where to save query.
	// Default: - No work group.
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

