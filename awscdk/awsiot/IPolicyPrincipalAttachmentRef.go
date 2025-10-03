package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyPrincipalAttachment.
// Experimental.
type IPolicyPrincipalAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPolicyPrincipalAttachmentRef
type jsiiProxy_IPolicyPrincipalAttachmentRef struct {
	internal.Type__constructsIConstruct
}

