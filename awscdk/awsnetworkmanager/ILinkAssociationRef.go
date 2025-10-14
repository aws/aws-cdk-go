package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LinkAssociation.
// Experimental.
type ILinkAssociationRef interface {
	constructs.IConstruct
	// A reference to a LinkAssociation resource.
	// Experimental.
	LinkAssociationRef() *LinkAssociationReference
}

// The jsii proxy for ILinkAssociationRef
type jsiiProxy_ILinkAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILinkAssociationRef) LinkAssociationRef() *LinkAssociationReference {
	var returns *LinkAssociationReference
	_jsii_.Get(
		j,
		"linkAssociationRef",
		&returns,
	)
	return returns
}

