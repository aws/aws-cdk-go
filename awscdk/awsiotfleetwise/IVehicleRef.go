package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Vehicle.
// Experimental.
type IVehicleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVehicleRef
type jsiiProxy_IVehicleRef struct {
	internal.Type__constructsIConstruct
}

