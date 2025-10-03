package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PodIdentityAssociation.
// Experimental.
type IPodIdentityAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPodIdentityAssociationRef
type jsiiProxy_IPodIdentityAssociationRef struct {
	internal.Type__constructsIConstruct
}

