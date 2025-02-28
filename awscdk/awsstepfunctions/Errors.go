package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Predefined error strings Error names in Amazon States Language - https://states-language.net/spec.html#appendix-a Error handling in Step Functions - https://docs.aws.amazon.com/step-functions/latest/dg/concepts-error-handling.html.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a table
//   table := dynamodb.NewTable(this, jsii.String("montable"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   })
//
//   finalStatus := sfn.NewPass(this, jsii.String("final step"))
//
//   // States language JSON to put an item into DynamoDB
//   // snippet generated from https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-code-snippet.html#tutorial-code-snippet-1
//   stateJson := map[string]interface{}{
//   	"Type": jsii.String("Task"),
//   	"Resource": jsii.String("arn:aws:states:::dynamodb:putItem"),
//   	"Parameters": map[string]interface{}{
//   		"TableName": table.tableName,
//   		"Item": map[string]map[string]*string{
//   			"id": map[string]*string{
//   				"S": jsii.String("MyEntry"),
//   			},
//   		},
//   	},
//   	"ResultPath": nil,
//   }
//
//   // custom state which represents a task to insert data into DynamoDB
//   custom := sfn.NewCustomState(this, jsii.String("my custom task"), &CustomStateProps{
//   	StateJson: StateJson,
//   })
//
//   // catch errors with addCatch
//   errorHandler := sfn.NewPass(this, jsii.String("handle failure"))
//   custom.AddCatch(errorHandler)
//
//   // retry the task if something goes wrong
//   custom.AddRetry(&RetryProps{
//   	Errors: []*string{
//   		sfn.Errors_ALL(),
//   	},
//   	Interval: awscdk.Duration_Seconds(jsii.Number(10)),
//   	MaxAttempts: jsii.Number(5),
//   })
//
//   chain := sfn.Chain_Start(custom).Next(finalStatus)
//
//   sm := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(chain),
//   	Timeout: awscdk.Duration_*Seconds(jsii.Number(30)),
//   	Comment: jsii.String("a super cool state machine"),
//   })
//
//   // don't forget permissions. You need to assign them
//   table.GrantWriteData(sm)
//
type Errors interface {
}

// The jsii proxy struct for Errors
type jsiiProxy_Errors struct {
	_ byte // padding
}

func NewErrors() Errors {
	_init_.Initialize()

	j := jsiiProxy_Errors{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewErrors_Override(e Errors) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		nil, // no parameters
		e,
	)
}

func Errors_ALL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"ALL",
		&returns,
	)
	return returns
}

func Errors_BRANCH_FAILED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"BRANCH_FAILED",
		&returns,
	)
	return returns
}

func Errors_HEARTBEAT_TIMEOUT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"HEARTBEAT_TIMEOUT",
		&returns,
	)
	return returns
}

func Errors_NO_CHOICE_MATCHED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"NO_CHOICE_MATCHED",
		&returns,
	)
	return returns
}

func Errors_PARAMETER_PATH_FAILURE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"PARAMETER_PATH_FAILURE",
		&returns,
	)
	return returns
}

func Errors_PERMISSIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"PERMISSIONS",
		&returns,
	)
	return returns
}

func Errors_RESULT_PATH_MATCH_FAILURE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"RESULT_PATH_MATCH_FAILURE",
		&returns,
	)
	return returns
}

func Errors_TASKS_FAILED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"TASKS_FAILED",
		&returns,
	)
	return returns
}

func Errors_TIMEOUT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"TIMEOUT",
		&returns,
	)
	return returns
}

