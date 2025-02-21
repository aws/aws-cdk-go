package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Interface which all EventBus based classes MUST implement.
type IEventBus interface {
	awscdk.IResource
	// Create an EventBridge archive to send events to.
	//
	// When you create an archive, incoming events might not immediately start being sent to the archive.
	// Allow a short period of time for changes to take effect.
	Archive(id *string, props *BaseArchiveProps) Archive
	// Grants an IAM Principal to send custom events to the eventBus so that they can be matched to rules.
	GrantPutEventsTo(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of this event bus resource.
	EventBusArn() *string
	// The physical ID of this event bus resource.
	EventBusName() *string
	// The JSON policy of this event bus resource.
	EventBusPolicy() *string
	// The partner event source to associate with this event bus resource.
	EventSourceName() *string
}

// The jsii proxy for IEventBus
type jsiiProxy_IEventBus struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEventBus) Archive(id *string, props *BaseArchiveProps) Archive {
	if err := i.validateArchiveParameters(id, props); err != nil {
		panic(err)
	}
	var returns Archive

	_jsii_.Invoke(
		i,
		"archive",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventBus) GrantPutEventsTo(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantPutEventsToParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPutEventsTo",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEventBus) EventBusArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBus) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBus) EventBusPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBus) EventSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceName",
		&returns,
	)
	return returns
}

