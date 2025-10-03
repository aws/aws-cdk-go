package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolUserToGroupAttachment.
// Experimental.
type IUserPoolUserToGroupAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserPoolUserToGroupAttachmentRef
type jsiiProxy_IUserPoolUserToGroupAttachmentRef struct {
	internal.Type__constructsIConstruct
}

