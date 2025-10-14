package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SourceApiAssociation.
// Experimental.
type ISourceApiAssociationRef interface {
	constructs.IConstruct
	// A reference to a SourceApiAssociation resource.
	// Experimental.
	SourceApiAssociationRef() *SourceApiAssociationReference
}

// The jsii proxy for ISourceApiAssociationRef
type jsiiProxy_ISourceApiAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISourceApiAssociationRef) SourceApiAssociationRef() *SourceApiAssociationReference {
	var returns *SourceApiAssociationReference
	_jsii_.Get(
		j,
		"sourceApiAssociationRef",
		&returns,
	)
	return returns
}

