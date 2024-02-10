package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for DynamoPutItem Task.
//
// Example:
//   var myTable table
//
//   tasks.NewDynamoPutItem(this, jsii.String("PutItem"), &DynamoPutItemProps{
//   	Item: map[string]dynamoAttributeValue{
//   		"MessageId": tasks.*dynamoAttributeValue_fromString(jsii.String("message-007")),
//   		"Text": tasks.*dynamoAttributeValue_fromString(sfn.JsonPath_stringAt(jsii.String("$.bar"))),
//   		"TotalCount": tasks.*dynamoAttributeValue_fromNumber(jsii.Number(10)),
//   	},
//   	Table: myTable,
//   })
//
type DynamoPutItemProps struct {
	// An optional description for this state.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
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
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: - The entire task input (JSON path '$').
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
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
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: - The entire JSON node determined by the state input, the task result,
	// and resultPath is passed to the next state (JSON path '$').
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: - Replaces the entire input with the result (JSON path '$').
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
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
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
	// A map of attribute name/value pairs, one for each attribute.
	//
	// Only the primary key attributes are required;
	// you can optionally provide other attribute name-value pairs for the item.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-Item
	//
	Item *map[string]DynamoAttributeValue `field:"required" json:"item" yaml:"item"`
	// The name of the table where the item should be written .
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// A condition that must be satisfied in order for a conditional PutItem operation to succeed.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ConditionExpression
	//
	// Default: - No condition expression.
	//
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ExpressionAttributeNames
	//
	// Default: - No expression attribute names.
	//
	ExpressionAttributeNames *map[string]*string `field:"optional" json:"expressionAttributeNames" yaml:"expressionAttributeNames"`
	// One or more values that can be substituted in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ExpressionAttributeValues
	//
	// Default: - No expression attribute values.
	//
	ExpressionAttributeValues *map[string]DynamoAttributeValue `field:"optional" json:"expressionAttributeValues" yaml:"expressionAttributeValues"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ReturnConsumedCapacity
	//
	// Default: DynamoConsumedCapacity.NONE
	//
	ReturnConsumedCapacity DynamoConsumedCapacity `field:"optional" json:"returnConsumedCapacity" yaml:"returnConsumedCapacity"`
	// The item collection metrics to returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LSI.html#LSI.ItemCollections
	//
	// Default: DynamoItemCollectionMetrics.NONE
	//
	ReturnItemCollectionMetrics DynamoItemCollectionMetrics `field:"optional" json:"returnItemCollectionMetrics" yaml:"returnItemCollectionMetrics"`
	// Use ReturnValues if you want to get the item attributes as they appeared before they were updated with the PutItem request.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ReturnValues
	//
	// Default: DynamoReturnValues.NONE
	//
	ReturnValues DynamoReturnValues `field:"optional" json:"returnValues" yaml:"returnValues"`
}

