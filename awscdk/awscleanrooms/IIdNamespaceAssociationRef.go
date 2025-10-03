package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdNamespaceAssociation.
// Experimental.
type IIdNamespaceAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdNamespaceAssociationRef
type jsiiProxy_IIdNamespaceAssociationRef struct {
	internal.Type__constructsIConstruct
}

