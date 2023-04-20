package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Role to be assumed by the State Machine's execution role for invoking a task's resource.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var submitLambda function
//   var iamRole role
//
//
//   // use a fixed role for all task invocations
//   role := sfn.TaskRole_FromRole(iamRole)
//   // or use a json expression to resolve the role at runtime based on task inputs
//   //const role = sfn.TaskRole.fromRoleArnJsonPath('$.RoleArn');
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &LambdaInvokeProps{
//   	LambdaFunction: submitLambda,
//   	OutputPath: jsii.String("$.Payload"),
//   	// use credentials
//   	Credentials: &Credentials{
//   		Role: *Role,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-task-state.html#task-state-fields
//
type TaskRole interface {
	// Retrieves the resource for use in IAM Policies for this TaskRole.
	Resource() *string
	// Retrieves the roleArn for this TaskRole.
	RoleArn() *string
}

// The jsii proxy struct for TaskRole
type jsiiProxy_TaskRole struct {
	_ byte // padding
}

func (j *jsiiProxy_TaskRole) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}


func NewTaskRole_Override(t TaskRole) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.TaskRole",
		nil, // no parameters
		t,
	)
}

// Construct a task role based on the provided IAM Role.
func TaskRole_FromRole(role awsiam.IRole) TaskRole {
	_init_.Initialize()

	if err := validateTaskRole_FromRoleParameters(role); err != nil {
		panic(err)
	}
	var returns TaskRole

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.TaskRole",
		"fromRole",
		[]interface{}{role},
		&returns,
	)

	return returns
}

// Construct a task role retrieved from task inputs using a json expression.
//
// Example:
//   sfn.TaskRole_FromRoleArnJsonPath(jsii.String("$.RoleArn"))
//
func TaskRole_FromRoleArnJsonPath(expression *string) TaskRole {
	_init_.Initialize()

	if err := validateTaskRole_FromRoleArnJsonPathParameters(expression); err != nil {
		panic(err)
	}
	var returns TaskRole

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.TaskRole",
		"fromRoleArnJsonPath",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

