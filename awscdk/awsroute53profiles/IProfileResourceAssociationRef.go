package awsroute53profiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53profiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfileResourceAssociation.
// Experimental.
type IProfileResourceAssociationRef interface {
	constructs.IConstruct
	// A reference to a ProfileResourceAssociation resource.
	// Experimental.
	ProfileResourceAssociationRef() *ProfileResourceAssociationReference
}

// The jsii proxy for IProfileResourceAssociationRef
type jsiiProxy_IProfileResourceAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProfileResourceAssociationRef) ProfileResourceAssociationRef() *ProfileResourceAssociationReference {
	var returns *ProfileResourceAssociationReference
	_jsii_.Get(
		j,
		"profileResourceAssociationRef",
		&returns,
	)
	return returns
}

