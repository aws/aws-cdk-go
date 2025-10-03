package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxLustre.
// Experimental.
type ILocationFSxLustreRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationFSxLustreRef
type jsiiProxy_ILocationFSxLustreRef struct {
	internal.Type__constructsIConstruct
}

