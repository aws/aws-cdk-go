package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliverySource.
// Experimental.
type IDeliverySourceRef interface {
	constructs.IConstruct
	// A reference to a DeliverySource resource.
	// Experimental.
	DeliverySourceRef() *DeliverySourceReference
}

// The jsii proxy for IDeliverySourceRef
type jsiiProxy_IDeliverySourceRef struct {
	internal.Type__constructsIConstruct
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

