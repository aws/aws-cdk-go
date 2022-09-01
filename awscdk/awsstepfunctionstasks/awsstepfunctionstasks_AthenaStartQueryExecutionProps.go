package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for starting a Query Execution.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Start Athena Query"), &athenaStartQueryExecutionProps{
//   	queryString: sfn.jsonPath.stringAt(jsii.String("$.queryString")),
//   	queryExecutionContext: &queryExecutionContext{
//   		databaseName: jsii.String("mydatabase"),
//   	},
//   	resultConfiguration: &resultConfiguration{
//   		encryptionConfiguration: &encryptionConfiguration{
//   			encryptionOption: tasks.encryptionOption_S3_MANAGED,
//   		},
//   		outputLocation: &location{
//   			bucketName: jsii.String("query-results-bucket"),
//   			objectKey: jsii.String("folder"),
//   		},
//   	},
//   })
//
type AthenaStartQueryExecutionProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
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
	// Timeout for the state machine.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Query that will be started.
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Unique string string to ensure idempotence.
	ClientRequestToken *string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
	// Database within which query executes.
	QueryExecutionContext *QueryExecutionContext `field:"optional" json:"queryExecutionContext" yaml:"queryExecutionContext"`
	// Configuration on how and where to save query.
	ResultConfiguration *ResultConfiguration `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
	// Configuration on how and where to save query.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

