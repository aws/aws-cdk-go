package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
)

// Properties for DynamoGetItem Task using JSONPath.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var dynamoAttributeValue DynamoAttributeValue
//   var dynamoProjectionExpression DynamoProjectionExpression
//   var resultSelector interface{}
//   var tableRef ITableRef
//   var taskRole TaskRole
//   var timeout Timeout
//
//   dynamoGetItemJsonPathProps := &DynamoGetItemJsonPathProps{
//   	Key: map[string]DynamoAttributeValue{
//   		"keyKey": dynamoAttributeValue,
//   	},
//   	Table: tableRef,
//
//   	// the properties below are optional
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	ConsistentRead: jsii.Boolean(false),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	ExpressionAttributeNames: map[string]*string{
//   		"expressionAttributeNamesKey": jsii.String("expressionAttributeNames"),
//   	},
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	InputPath: jsii.String("inputPath"),
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	OutputPath: jsii.String("outputPath"),
//   	ProjectionExpression: []DynamoProjectionExpression{
//   		dynamoProjectionExpression,
//   	},
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ResultPath: jsii.String("resultPath"),
//   	ResultSelector: map[string]interface{}{
//   		"resultSelectorKey": resultSelector,
//   	},
//   	ReturnConsumedCapacity: awscdk.Aws_stepfunctions_tasks.DynamoConsumedCapacity_INDEXES,
//   	StateName: jsii.String("stateName"),
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type DynamoGetItemJsonPathProps struct {
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
	// Primary key of the item to retrieve.
	//
	// For the primary key, you must provide all of the attributes.
	// For example, with a simple primary key, you only need to provide a value for the partition key.
	// For a composite primary key, you must provide values for both the partition key and the sort key.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-Key
	//
	Key *map[string]DynamoAttributeValue `field:"required" json:"key" yaml:"key"`
	// The name of the table containing the requested item.
	Table interfacesawsdynamodb.ITableRef `field:"required" json:"table" yaml:"table"`
	// Determines the read consistency model: If set to true, then the operation uses strongly consistent reads;
	//
	// otherwise, the operation uses eventually consistent reads.
	// Default: false.
	//
	ConsistentRead *bool `field:"optional" json:"consistentRead" yaml:"consistentRead"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ExpressionAttributeNames
	//
	// Default: - No expression attributes.
	//
	ExpressionAttributeNames *map[string]*string `field:"optional" json:"expressionAttributeNames" yaml:"expressionAttributeNames"`
	// An array of DynamoProjectionExpression that identifies one or more attributes to retrieve from the table.
	//
	// These attributes can include scalars, sets, or elements of a JSON document.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ProjectionExpression
	//
	// Default: - No projection expression.
	//
	ProjectionExpression *[]DynamoProjectionExpression `field:"optional" json:"projectionExpression" yaml:"projectionExpression"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ReturnConsumedCapacity
	//
	// Default: DynamoConsumedCapacity.NONE
	//
	ReturnConsumedCapacity DynamoConsumedCapacity `field:"optional" json:"returnConsumedCapacity" yaml:"returnConsumedCapacity"`
}

