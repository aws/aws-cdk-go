package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaceIndex.
// Experimental.
type IPlaceIndexRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPlaceIndexRef
type jsiiProxy_IPlaceIndexRef struct {
	internal.Type__constructsIConstruct
}

