package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CrossAccountAttachment.
// Experimental.
type ICrossAccountAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICrossAccountAttachmentRef
type jsiiProxy_ICrossAccountAttachmentRef struct {
	internal.Type__constructsIConstruct
}

