package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkServiceAssociation.
// Experimental.
type IServiceNetworkServiceAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceNetworkServiceAssociationRef
type jsiiProxy_IServiceNetworkServiceAssociationRef struct {
	internal.Type__constructsIConstruct
}

