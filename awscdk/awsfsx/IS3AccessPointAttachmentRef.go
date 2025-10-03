package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfsx/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a S3AccessPointAttachment.
// Experimental.
type IS3AccessPointAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IS3AccessPointAttachmentRef
type jsiiProxy_IS3AccessPointAttachmentRef struct {
	internal.Type__constructsIConstruct
}

