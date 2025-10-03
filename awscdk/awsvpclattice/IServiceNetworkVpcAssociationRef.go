package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceNetworkVpcAssociation.
// Experimental.
type IServiceNetworkVpcAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IServiceNetworkVpcAssociationRef
type jsiiProxy_IServiceNetworkVpcAssociationRef struct {
	internal.Type__constructsIConstruct
}

