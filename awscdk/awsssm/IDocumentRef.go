package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Document.
// Experimental.
type IDocumentRef interface {
	constructs.IConstruct
	// A reference to a Document resource.
	// Experimental.
	DocumentRef() *DocumentReference
}

// The jsii proxy for IDocumentRef
type jsiiProxy_IDocumentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDocumentRef) DocumentRef() *DocumentReference {
	var returns *DocumentReference
	_jsii_.Get(
		j,
		"documentRef",
		&returns,
	)
	return returns
}

