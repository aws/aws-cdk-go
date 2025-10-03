package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationNFS.
// Experimental.
type ILocationNFSRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILocationNFSRef
type jsiiProxy_ILocationNFSRef struct {
	internal.Type__constructsIConstruct
}

