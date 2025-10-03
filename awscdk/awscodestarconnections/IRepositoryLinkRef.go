package awscodestarconnections

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryLink.
// Experimental.
type IRepositoryLinkRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRepositoryLinkRef
type jsiiProxy_IRepositoryLinkRef struct {
	internal.Type__constructsIConstruct
}

