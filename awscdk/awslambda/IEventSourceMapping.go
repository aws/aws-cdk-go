package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
)

// Represents an event source mapping for a lambda function.
// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html
//
type IEventSourceMapping interface {
	awscdk.IResource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
}

// The jsii proxy for IEventSourceMapping
type jsiiProxy_IEventSourceMapping struct {
	internal.Type__awscdkIResource
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

