package awsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CapacityReservation.
// Experimental.
type ICapacityReservationRef interface {
	constructs.IConstruct
	// A reference to a CapacityReservation resource.
	// Experimental.
	CapacityReservationRef() *CapacityReservationReference
}

// The jsii proxy for ICapacityReservationRef
type jsiiProxy_ICapacityReservationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICapacityReservationRef) CapacityReservationRef() *CapacityReservationReference {
	var returns *CapacityReservationReference
	_jsii_.Get(
		j,
		"capacityReservationRef",
		&returns,
	)
	return returns
}

