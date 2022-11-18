package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents/internal"
)

// Represents a field in the event pattern.
// Experimental.
type EventField interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	// Experimental.
	CreationStack() *[]*string
	// Human readable display hint about the event pattern.
	// Experimental.
	DisplayHint() *string
	// the path to a field in the event pattern.
	// Experimental.
	Path() *string
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_ctx awscdk.IResolveContext) interface{}
	// Convert the path to the field in the event pattern to JSON.
	// Experimental.
	ToJSON() *string
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
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
// Experimental.
func EventField_FromPath(path *string) *string {
	_init_.Initialize()

	if err := validateEventField_FromPathParameters(path); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_events.EventField",
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
		"monocdk.aws_events.EventField",
		"account",
		&returns,
	)
	return returns
}

func EventField_DetailType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_events.EventField",
		"detailType",
		&returns,
	)
	return returns
}

func EventField_EventId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_events.EventField",
		"eventId",
		&returns,
	)
	return returns
}

func EventField_Region() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_events.EventField",
		"region",
		&returns,
	)
	return returns
}

func EventField_Source() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_events.EventField",
		"source",
		&returns,
	)
	return returns
}

func EventField_Time() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_events.EventField",
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

