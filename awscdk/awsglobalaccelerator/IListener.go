package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
)

// Interface of the Listener.
type IListener interface {
	awscdk.IResource
	// The ARN of the listener.
	ListenerArn() *string
}

// The jsii proxy for IListener
type jsiiProxy_IListener struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

