package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationHDFS.
// Experimental.
type ILocationHDFSRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationHDFSRef
type jsiiProxy_ILocationHDFSRef struct {
	internal.Type__constructsIConstruct
}

