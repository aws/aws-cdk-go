package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CrossAccountAttachment.
// Experimental.
type ICrossAccountAttachmentRef interface {
	constructs.IConstruct
	// A reference to a CrossAccountAttachment resource.
	// Experimental.
	CrossAccountAttachmentRef() *CrossAccountAttachmentReference
}

// The jsii proxy for ICrossAccountAttachmentRef
type jsiiProxy_ICrossAccountAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICrossAccountAttachmentRef) CrossAccountAttachmentRef() *CrossAccountAttachmentReference {
	var returns *CrossAccountAttachmentReference
	_jsii_.Get(
		j,
		"crossAccountAttachmentRef",
		&returns,
	)
	return returns
}

