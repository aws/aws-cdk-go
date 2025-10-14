package awsroute53profiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53profiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfileAssociation.
// Experimental.
type IProfileAssociationRef interface {
	constructs.IConstruct
	// A reference to a ProfileAssociation resource.
	// Experimental.
	ProfileAssociationRef() *ProfileAssociationReference
}

// The jsii proxy for IProfileAssociationRef
type jsiiProxy_IProfileAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProfileAssociationRef) ProfileAssociationRef() *ProfileAssociationReference {
	var returns *ProfileAssociationReference
	_jsii_.Get(
		j,
		"profileAssociationRef",
		&returns,
	)
	return returns
}

