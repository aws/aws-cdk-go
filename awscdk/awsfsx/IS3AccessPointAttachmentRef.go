package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a S3AccessPointAttachment.
// Experimental.
type IS3AccessPointAttachmentRef interface {
	constructs.IConstruct
	// A reference to a S3AccessPointAttachment resource.
	// Experimental.
	S3AccessPointAttachmentRef() *S3AccessPointAttachmentReference
}

// The jsii proxy for IS3AccessPointAttachmentRef
type jsiiProxy_IS3AccessPointAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IS3AccessPointAttachmentRef) S3AccessPointAttachmentRef() *S3AccessPointAttachmentReference {
	var returns *S3AccessPointAttachmentReference
	_jsii_.Get(
		j,
		"s3AccessPointAttachmentRef",
		&returns,
	)
	return returns
}

