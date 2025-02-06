package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Timeout for a task or heartbeat.
//
// Example:
//   tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
//   	GlueJobName: jsii.String("my-glue-job"),
//   	Arguments: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"key": jsii.String("value"),
//   	}),
//   	TaskTimeout: sfn.Timeout_Duration(awscdk.Duration_Minutes(jsii.Number(30))),
//   	NotifyDelayAfter: awscdk.Duration_*Minutes(jsii.Number(5)),
//   })
//
type Timeout interface {
	// JSONata expression for this timeout.
	JsonataExpression() *string
	// Path for this timeout.
	Path() *string
	// Seconds for this timeout.
	Seconds() *float64
}

// The jsii proxy struct for Timeout
type jsiiProxy_Timeout struct {
	_ byte // padding
}

func (j *jsiiProxy_Timeout) JsonataExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonataExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Timeout) Seconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"seconds",
		&returns,
	)
	return returns
}


func NewTimeout_Override(t Timeout) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Timeout",
		nil, // no parameters
		t,
	)
}

// Use a dynamic timeout specified by a path in the state input.
//
// The path must select a field whose value is a positive integer.
func Timeout_At(path *string) Timeout {
	_init_.Initialize()

	if err := validateTimeout_AtParameters(path); err != nil {
		panic(err)
	}
	var returns Timeout

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Timeout",
		"at",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Use a duration as timeout.
func Timeout_Duration(duration awscdk.Duration) Timeout {
	_init_.Initialize()

	if err := validateTimeout_DurationParameters(duration); err != nil {
		panic(err)
	}
	var returns Timeout

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Timeout",
		"duration",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Use a dynamic timeout specified by a JSONata expression.
//
// The JSONata expression value must be a positive integer.
func Timeout_Jsonata(jsonataExpression *string) Timeout {
	_init_.Initialize()

	if err := validateTimeout_JsonataParameters(jsonataExpression); err != nil {
		panic(err)
	}
	var returns Timeout

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Timeout",
		"jsonata",
		[]interface{}{jsonataExpression},
		&returns,
	)

	return returns
}

