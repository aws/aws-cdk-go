package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyAssociation.
// Experimental.
type IPolicyAssociationRef interface {
	constructs.IConstruct
	// A reference to a PolicyAssociation resource.
	// Experimental.
	PolicyAssociationRef() *PolicyAssociationReference
}

// The jsii proxy for IPolicyAssociationRef
type jsiiProxy_IPolicyAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPolicyAssociationRef) PolicyAssociationRef() *PolicyAssociationReference {
	var returns *PolicyAssociationReference
	_jsii_.Get(
		j,
		"policyAssociationRef",
		&returns,
	)
	return returns
}

