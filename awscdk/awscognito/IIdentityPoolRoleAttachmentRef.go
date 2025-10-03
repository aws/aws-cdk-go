package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPoolRoleAttachment.
// Experimental.
type IIdentityPoolRoleAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentityPoolRoleAttachmentRef
type jsiiProxy_IIdentityPoolRoleAttachmentRef struct {
	internal.Type__constructsIConstruct
}

