package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationEFS.
// Experimental.
type ILocationEFSRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationEFSRef
type jsiiProxy_ILocationEFSRef struct {
	internal.Type__constructsIConstruct
}

