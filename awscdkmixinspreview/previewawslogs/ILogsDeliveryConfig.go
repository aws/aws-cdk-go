package previewawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// The individual elements of a logs delivery integration.
// Experimental.
type ILogsDeliveryConfig interface {
	// The logs delivery.
	// Experimental.
	Delivery() interfacesawslogs.IDeliveryRef
	// The logs delivery destination.
	// Experimental.
	DeliveryDestination() interfacesawslogs.IDeliveryDestinationRef
	// The logs delivery source.
	// Experimental.
	DeliverySource() interfacesawslogs.IDeliverySourceRef
}

// The jsii proxy for ILogsDeliveryConfig
type jsiiProxy_ILogsDeliveryConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_ILogsDeliveryConfig) Delivery() interfacesawslogs.IDeliveryRef {
	var returns interfacesawslogs.IDeliveryRef
	_jsii_.Get(
		j,
		"delivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogsDeliveryConfig) DeliveryDestination() interfacesawslogs.IDeliveryDestinationRef {
	var returns interfacesawslogs.IDeliveryDestinationRef
	_jsii_.Get(
		j,
		"deliveryDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogsDeliveryConfig) DeliverySource() interfacesawslogs.IDeliverySourceRef {
	var returns interfacesawslogs.IDeliverySourceRef
	_jsii_.Get(
		j,
		"deliverySource",
		&returns,
	)
	return returns
}

