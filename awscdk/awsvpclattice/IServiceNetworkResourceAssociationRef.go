package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkResourceAssociation.
// Experimental.
type IServiceNetworkResourceAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceNetworkResourceAssociationRef
type jsiiProxy_IServiceNetworkResourceAssociationRef struct {
	internal.Type__constructsIConstruct
}

