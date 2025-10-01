package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Delivery.
// Experimental.
type IDeliveryRef interface {
	constructs.IConstruct
	// A reference to a Delivery resource.
	// Experimental.
	DeliveryRef() *DeliveryReference
}

// The jsii proxy for IDeliveryRef
type jsiiProxy_IDeliveryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeliveryRef) DeliveryRef() *DeliveryReference {
	var returns *DeliveryReference
	_jsii_.Get(
		j,
		"deliveryRef",
		&returns,
	)
	return returns
}

