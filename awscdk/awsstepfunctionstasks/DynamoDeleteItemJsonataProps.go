package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdynamodb"
)

// Properties for DynamoDeleteItem Task using JSONata.
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
//   var outputs interface{}
//   var tableRef ITableRef
//   var taskRole TaskRole
//   var timeout Timeout
//
//   dynamoDeleteItemJsonataProps := &DynamoDeleteItemJsonataProps{
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
//   	ConditionExpression: jsii.String("conditionExpression"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	ExpressionAttributeNames: map[string]*string{
//   		"expressionAttributeNamesKey": jsii.String("expressionAttributeNames"),
//   	},
//   	ExpressionAttributeValues: map[string]DynamoAttributeValue{
//   		"expressionAttributeValuesKey": dynamoAttributeValue,
//   	},
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	Outputs: outputs,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	ReturnConsumedCapacity: awscdk.Aws_stepfunctions_tasks.DynamoConsumedCapacity_INDEXES,
//   	ReturnItemCollectionMetrics: awscdk.*Aws_stepfunctions_tasks.DynamoItemCollectionMetrics_SIZE,
//   	ReturnValues: awscdk.*Aws_stepfunctions_tasks.DynamoReturnValues_NONE,
//   	StateName: jsii.String("stateName"),
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type DynamoDeleteItemJsonataProps struct {
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
	// A condition that must be satisfied in order for a conditional DeleteItem to succeed.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ConditionExpression
	//
	// Default: - No condition expression.
	//
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ExpressionAttributeNames
	//
	// Default: - No expression attribute names.
	//
	ExpressionAttributeNames *map[string]*string `field:"optional" json:"expressionAttributeNames" yaml:"expressionAttributeNames"`
	// One or more values that can be substituted in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ExpressionAttributeValues
	//
	// Default: - No expression attribute values.
	//
	ExpressionAttributeValues *map[string]DynamoAttributeValue `field:"optional" json:"expressionAttributeValues" yaml:"expressionAttributeValues"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ReturnConsumedCapacity
	//
	// Default: DynamoConsumedCapacity.NONE
	//
	ReturnConsumedCapacity DynamoConsumedCapacity `field:"optional" json:"returnConsumedCapacity" yaml:"returnConsumedCapacity"`
	// Determines whether item collection metrics are returned.
	//
	// If set to SIZE, the response includes statistics about item collections, if any,
	// that were modified during the operation are returned in the response.
	// If set to NONE (the default), no statistics are returned.
	// Default: DynamoItemCollectionMetrics.NONE
	//
	ReturnItemCollectionMetrics DynamoItemCollectionMetrics `field:"optional" json:"returnItemCollectionMetrics" yaml:"returnItemCollectionMetrics"`
	// Use ReturnValues if you want to get the item attributes as they appeared before they were deleted.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ReturnValues
	//
	// Default: DynamoReturnValues.NONE
	//
	ReturnValues DynamoReturnValues `field:"optional" json:"returnValues" yaml:"returnValues"`
}

