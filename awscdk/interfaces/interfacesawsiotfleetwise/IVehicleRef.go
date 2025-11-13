package interfacesawsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Vehicle.
// Experimental.
type IVehicleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Vehicle resource.
	// Experimental.
	VehicleRef() *VehicleReference
}

// The jsii proxy for IVehicleRef
type jsiiProxy_IVehicleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVehicleRef) VehicleRef() *VehicleReference {
	var returns *VehicleReference
	_jsii_.Get(
		j,
		"vehicleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVehicleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVehicleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

