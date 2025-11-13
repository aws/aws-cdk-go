package interfacesawskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliveryStream.
// Experimental.
type IDeliveryStreamRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DeliveryStream resource.
	// Experimental.
	DeliveryStreamRef() *DeliveryStreamReference
}

// The jsii proxy for IDeliveryStreamRef
type jsiiProxy_IDeliveryStreamRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDeliveryStreamRef) DeliveryStreamRef() *DeliveryStreamReference {
	var returns *DeliveryStreamReference
	_jsii_.Get(
		j,
		"deliveryStreamRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStreamRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStreamRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

