package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TagOptionAssociation.
// Experimental.
type ITagOptionAssociationRef interface {
	constructs.IConstruct
	// A reference to a TagOptionAssociation resource.
	// Experimental.
	TagOptionAssociationRef() *TagOptionAssociationReference
}

// The jsii proxy for ITagOptionAssociationRef
type jsiiProxy_ITagOptionAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITagOptionAssociationRef) TagOptionAssociationRef() *TagOptionAssociationReference {
	var returns *TagOptionAssociationReference
	_jsii_.Get(
		j,
		"tagOptionAssociationRef",
		&returns,
	)
	return returns
}

