package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CapacityReservationFleet.
// Experimental.
type ICapacityReservationFleetRef interface {
	constructs.IConstruct
	// A reference to a CapacityReservationFleet resource.
	// Experimental.
	CapacityReservationFleetRef() *CapacityReservationFleetReference
}

// The jsii proxy for ICapacityReservationFleetRef
type jsiiProxy_ICapacityReservationFleetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICapacityReservationFleetRef) CapacityReservationFleetRef() *CapacityReservationFleetReference {
	var returns *CapacityReservationFleetReference
	_jsii_.Get(
		j,
		"capacityReservationFleetRef",
		&returns,
	)
	return returns
}

