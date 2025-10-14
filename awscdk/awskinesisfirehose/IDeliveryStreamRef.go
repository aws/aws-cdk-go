package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliveryStream.
// Experimental.
type IDeliveryStreamRef interface {
	constructs.IConstruct
	// A reference to a DeliveryStream resource.
	// Experimental.
	DeliveryStreamRef() *DeliveryStreamReference
}

// The jsii proxy for IDeliveryStreamRef
type jsiiProxy_IDeliveryStreamRef struct {
	internal.Type__constructsIConstruct
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

