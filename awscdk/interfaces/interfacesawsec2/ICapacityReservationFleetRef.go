package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CapacityReservationFleet.
// Experimental.
type ICapacityReservationFleetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CapacityReservationFleet resource.
	// Experimental.
	CapacityReservationFleetRef() *CapacityReservationFleetReference
}

// The jsii proxy for ICapacityReservationFleetRef
type jsiiProxy_ICapacityReservationFleetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ICapacityReservationFleetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityReservationFleetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

