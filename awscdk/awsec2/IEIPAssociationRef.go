package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EIPAssociation.
// Experimental.
type IEIPAssociationRef interface {
	constructs.IConstruct
	// A reference to a EIPAssociation resource.
	// Experimental.
	EipAssociationRef() *EIPAssociationReference
}

// The jsii proxy for IEIPAssociationRef
type jsiiProxy_IEIPAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEIPAssociationRef) EipAssociationRef() *EIPAssociationReference {
	var returns *EIPAssociationReference
	_jsii_.Get(
		j,
		"eipAssociationRef",
		&returns,
	)
	return returns
}

