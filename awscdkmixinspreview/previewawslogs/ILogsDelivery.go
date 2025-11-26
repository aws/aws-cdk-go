package previewawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the delivery of vended logs to a destination.
// Experimental.
type ILogsDelivery interface {
	// Binds the destination to a delivery source and creates a delivery connection between the source and destination.
	//
	// Returns: The delivery reference.
	// Experimental.
	Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy for ILogsDelivery
type jsiiProxy_ILogsDelivery struct {
	_ byte // padding
}

func (i *jsiiProxy_ILogsDelivery) Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := i.validateBindParameters(scope, deliverySource, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, deliverySource, sourceResourceArn},
		&returns,
	)

	return returns
}

