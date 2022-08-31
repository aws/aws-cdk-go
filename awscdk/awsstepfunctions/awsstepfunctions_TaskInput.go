package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Type union for task classes that accept multiple types of payload.
//
// Example:
//   var fn function
//
//   tasks.NewLambdaInvoke(this, jsii.String("Invoke with callback"), &lambdaInvokeProps{
//   	lambdaFunction: fn,
//   	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
//   	payload: sfn.taskInput.fromObject(map[string]interface{}{
//   		"token": sfn.JsonPath.taskToken,
//   		"input": sfn.JsonPath.stringAt(jsii.String("$.someField")),
//   	}),
//   })
//
// Experimental.
type TaskInput interface {
	// type of task input.
	// Experimental.
	Type() InputType
	// payload for the corresponding input type.
	//
	// It can be a JSON-encoded object, context, data, etc.
	// Experimental.
	Value() interface{}
}

// The jsii proxy struct for TaskInput
type jsiiProxy_TaskInput struct {
	_ byte // padding
}

func (j *jsiiProxy_TaskInput) Type() InputType {
	var returns InputType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskInput) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a part of the task context as task input.
//
// Use this when you want to use a subobject or string from
// the current task context as complete payload
// to a task.
// Deprecated: Use `fromJsonPathAt`.
func TaskInput_FromContextAt(path *string) TaskInput {
	_init_.Initialize()

	var returns TaskInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.TaskInput",
		"fromContextAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Use a part of the execution data as task input.
//
// Use this when you want to use a subobject or string from
// the current state machine execution as complete payload
// to a task.
// Deprecated: Use `fromJsonPathAt`.
func TaskInput_FromDataAt(path *string) TaskInput {
	_init_.Initialize()

	var returns TaskInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.TaskInput",
		"fromDataAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Use a part of the execution data or task context as task input.
//
// Use this when you want to use a subobject or string from
// the current state machine execution or the current task context
// as complete payload to a task.
// Experimental.
func TaskInput_FromJsonPathAt(path *string) TaskInput {
	_init_.Initialize()

	var returns TaskInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.TaskInput",
		"fromJsonPathAt",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Use an object as task input.
//
// This object may contain JSON path fields as object values, if desired.
// Experimental.
func TaskInput_FromObject(obj *map[string]interface{}) TaskInput {
	_init_.Initialize()

	var returns TaskInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.TaskInput",
		"fromObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Use a literal string as task input.
//
// This might be a JSON-encoded object, or just a text.
// Experimental.
func TaskInput_FromText(text *string) TaskInput {
	_init_.Initialize()

	var returns TaskInput

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.TaskInput",
		"fromText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

