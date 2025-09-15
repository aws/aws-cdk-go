package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PodIdentityAssociation.
// Experimental.
type IPodIdentityAssociationRef interface {
	constructs.IConstruct
	// A reference to a PodIdentityAssociation resource.
	// Experimental.
	PodIdentityAssociationRef() *PodIdentityAssociationReference
}

// The jsii proxy for IPodIdentityAssociationRef
type jsiiProxy_IPodIdentityAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPodIdentityAssociationRef) PodIdentityAssociationRef() *PodIdentityAssociationReference {
	var returns *PodIdentityAssociationReference
	_jsii_.Get(
		j,
		"podIdentityAssociationRef",
		&returns,
	)
	return returns
}

