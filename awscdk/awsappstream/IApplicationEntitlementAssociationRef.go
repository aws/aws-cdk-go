package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationEntitlementAssociation.
// Experimental.
type IApplicationEntitlementAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationEntitlementAssociationRef
type jsiiProxy_IApplicationEntitlementAssociationRef struct {
	internal.Type__constructsIConstruct
}

