package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationEntitlementAssociation.
// Experimental.
type IApplicationEntitlementAssociationRef interface {
	constructs.IConstruct
	// A reference to a ApplicationEntitlementAssociation resource.
	// Experimental.
	ApplicationEntitlementAssociationRef() *ApplicationEntitlementAssociationReference
}

// The jsii proxy for IApplicationEntitlementAssociationRef
type jsiiProxy_IApplicationEntitlementAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationEntitlementAssociationRef) ApplicationEntitlementAssociationRef() *ApplicationEntitlementAssociationReference {
	var returns *ApplicationEntitlementAssociationReference
	_jsii_.Get(
		j,
		"applicationEntitlementAssociationRef",
		&returns,
	)
	return returns
}

