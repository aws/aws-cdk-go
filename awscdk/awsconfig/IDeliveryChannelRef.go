package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DeliveryChannel.
// Experimental.
type IDeliveryChannelRef interface {
	constructs.IConstruct
	// A reference to a DeliveryChannel resource.
	// Experimental.
	DeliveryChannelRef() *DeliveryChannelReference
}

// The jsii proxy for IDeliveryChannelRef
type jsiiProxy_IDeliveryChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDeliveryChannelRef) DeliveryChannelRef() *DeliveryChannelReference {
	var returns *DeliveryChannelReference
	_jsii_.Get(
		j,
		"deliveryChannelRef",
		&returns,
	)
	return returns
}

