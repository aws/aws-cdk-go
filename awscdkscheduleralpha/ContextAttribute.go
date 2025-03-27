package awscdkscheduleralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A set of convenient static methods representing the Scheduler Context Attributes.
//
// These Context Attributes keywords can be used inside a ScheduleTargetInput.
// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule-context-attributes.html
//
// Deprecated.
type ContextAttribute interface {
	// Deprecated.
	Name() *string
	// Convert the path to the field in the event pattern to JSON.
	// Deprecated.
	ToString() *string
}

// The jsii proxy struct for ContextAttribute
type jsiiProxy_ContextAttribute struct {
	_ byte // padding
}

func (j *jsiiProxy_ContextAttribute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Escape hatch for other Context Attributes that may be added in the future.
// Deprecated.
func ContextAttribute_FromName(name *string) *string {
	_init_.Initialize()

	if err := validateContextAttribute_FromNameParameters(name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-scheduler-alpha.ContextAttribute",
		"fromName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func ContextAttribute_AttemptNumber() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-scheduler-alpha.ContextAttribute",
		"attemptNumber",
		&returns,
	)
	return returns
}

func ContextAttribute_ExecutionId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-scheduler-alpha.ContextAttribute",
		"executionId",
		&returns,
	)
	return returns
}

func ContextAttribute_ScheduleArn() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-scheduler-alpha.ContextAttribute",
		"scheduleArn",
		&returns,
	)
	return returns
}

func ContextAttribute_ScheduledTime() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-scheduler-alpha.ContextAttribute",
		"scheduledTime",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContextAttribute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

