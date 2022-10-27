package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
)

// Properties to reference an existing listener.
type INetworkListener interface {
	awscdk.IResource
	// ARN of the listener.
	ListenerArn() *string
}

// The jsii proxy for INetworkListener
type jsiiProxy_INetworkListener struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INetworkListener) ListenerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerArn",
		&returns,
	)
	return returns
}

