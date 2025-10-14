package awsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Vehicle.
// Experimental.
type IVehicleRef interface {
	constructs.IConstruct
	// A reference to a Vehicle resource.
	// Experimental.
	VehicleRef() *VehicleReference
}

// The jsii proxy for IVehicleRef
type jsiiProxy_IVehicleRef struct {
	internal.Type__constructsIConstruct
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

