package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceActionAssociation.
// Experimental.
type IServiceActionAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceActionAssociationRef
type jsiiProxy_IServiceActionAssociationRef struct {
	internal.Type__constructsIConstruct
}

