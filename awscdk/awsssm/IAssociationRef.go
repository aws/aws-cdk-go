package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Association.
// Experimental.
type IAssociationRef interface {
	constructs.IConstruct
	// A reference to a Association resource.
	// Experimental.
	AssociationRef() *AssociationReference
}

// The jsii proxy for IAssociationRef
type jsiiProxy_IAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssociationRef) AssociationRef() *AssociationReference {
	var returns *AssociationReference
	_jsii_.Get(
		j,
		"associationRef",
		&returns,
	)
	return returns
}

