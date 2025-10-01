package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackUserAssociation.
// Experimental.
type IStackUserAssociationRef interface {
	constructs.IConstruct
	// A reference to a StackUserAssociation resource.
	// Experimental.
	StackUserAssociationRef() *StackUserAssociationReference
}

// The jsii proxy for IStackUserAssociationRef
type jsiiProxy_IStackUserAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStackUserAssociationRef) StackUserAssociationRef() *StackUserAssociationReference {
	var returns *StackUserAssociationReference
	_jsii_.Get(
		j,
		"stackUserAssociationRef",
		&returns,
	)
	return returns
}

