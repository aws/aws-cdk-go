package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataRepositoryAssociation.
// Experimental.
type IDataRepositoryAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDataRepositoryAssociationRef
type jsiiProxy_IDataRepositoryAssociationRef struct {
	internal.Type__constructsIConstruct
}

