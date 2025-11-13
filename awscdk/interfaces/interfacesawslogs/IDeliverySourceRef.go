package interfacesawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliverySource.
// Experimental.
type IDeliverySourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DeliverySource resource.
	// Experimental.
	DeliverySourceRef() *DeliverySourceReference
}

// The jsii proxy for IDeliverySourceRef
type jsiiProxy_IDeliverySourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDeliverySourceRef) DeliverySourceRef() *DeliverySourceReference {
	var returns *DeliverySourceReference
	_jsii_.Get(
		j,
		"deliverySourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliverySourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliverySourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

