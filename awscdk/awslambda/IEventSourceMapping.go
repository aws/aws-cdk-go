package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an event source mapping for a lambda function.
// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html
//
type IEventSourceMapping interface {
	IEventSourceMappingRef
	awscdk.IResource
	// The ARN of the event source mapping (i.e. arn:aws:lambda:region:account-id:event-source-mapping/event-source-mapping-id).
	EventSourceMappingArn() *string
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
}

// The jsii proxy for IEventSourceMapping
type jsiiProxy_IEventSourceMapping struct {
	jsiiProxy_IEventSourceMappingRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEventSourceMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IEventSourceMapping) EventSourceMappingArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSourceMapping) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSourceMapping) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventSourceMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

