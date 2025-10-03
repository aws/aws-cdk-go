package awscodegurureviewer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodegurureviewer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryAssociation.
// Experimental.
type IRepositoryAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRepositoryAssociationRef
type jsiiProxy_IRepositoryAssociationRef struct {
	internal.Type__constructsIConstruct
}

