package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackFleetAssociation.
// Experimental.
type IStackFleetAssociationRef interface {
	constructs.IConstruct
	// A reference to a StackFleetAssociation resource.
	// Experimental.
	StackFleetAssociationRef() *StackFleetAssociationReference
}

// The jsii proxy for IStackFleetAssociationRef
type jsiiProxy_IStackFleetAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStackFleetAssociationRef) StackFleetAssociationRef() *StackFleetAssociationReference {
	var returns *StackFleetAssociationReference
	_jsii_.Get(
		j,
		"stackFleetAssociationRef",
		&returns,
	)
	return returns
}

