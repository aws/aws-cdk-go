package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
)

// Represents a field in the event pattern.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
//   	LogEvent: targets.LogGroupTargetInput_FromObject(&LogGroupTargetInputOptions{
//   		Timestamp: events.EventField_FromPath(jsii.String("$.time")),
//   		Message: events.EventField_*FromPath(jsii.String("$.detail-type")),
//   	}),
//   }))
//
type EventField interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// Human readable display hint about the event pattern.
	DisplayHint() *string
	// the path to a field in the event pattern.
	Path() *string
	// Produce the Token's value at resolution time.
	Resolve(_ctx awscdk.IResolveContext) interface{}
	// Convert the path to the field in the event pattern to JSON.
	ToJSON() *string
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	ToString() *string
}

// The jsii proxy struct for EventField
type jsiiProxy_EventField struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_EventField) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventField) DisplayHint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayHint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventField) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Extract a custom JSON path from the event.
func EventField_FromPath(path *string) *string {
	_init_.Initialize()

	if err := validateEventField_FromPathParameters(path); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventField",
		"fromPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func EventField_Account() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"account",
		&returns,
	)
	return returns
}

func EventField_DetailType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"detailType",
		&returns,
	)
	return returns
}

func EventField_EventId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"eventId",
		&returns,
	)
	return returns
}

func EventField_Region() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"region",
		&returns,
	)
	return returns
}

func EventField_Source() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"source",
		&returns,
	)
	return returns
}

func EventField_Time() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"time",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventField) Resolve(_ctx awscdk.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_ctx); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_ctx},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventField) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventField) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

