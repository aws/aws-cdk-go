package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Fleet.
// Experimental.
type IFleetRef interface {
	constructs.IConstruct
	// A reference to a Fleet resource.
	// Experimental.
	FleetRef() *FleetReference
}

// The jsii proxy for IFleetRef
type jsiiProxy_IFleetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFleetRef) FleetRef() *FleetReference {
	var returns *FleetReference
	_jsii_.Get(
		j,
		"fleetRef",
		&returns,
	)
	return returns
}

