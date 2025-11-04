package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliveryDestination.
// Experimental.
type IDeliveryDestinationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a DeliveryDestination resource.
	// Experimental.
	DeliveryDestinationRef() *DeliveryDestinationReference
}

// The jsii proxy for IDeliveryDestinationRef
type jsiiProxy_IDeliveryDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IDeliveryDestinationRef) DeliveryDestinationRef() *DeliveryDestinationReference {
	var returns *DeliveryDestinationReference
	_jsii_.Get(
		j,
		"deliveryDestinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryDestinationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

