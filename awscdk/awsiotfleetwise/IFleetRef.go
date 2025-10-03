package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Fleet.
// Experimental.
type IFleetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFleetRef
type jsiiProxy_IFleetRef struct {
	internal.Type__constructsIConstruct
}

