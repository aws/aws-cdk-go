package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPoolRoleAttachment.
// Experimental.
type IIdentityPoolRoleAttachmentRef interface {
	constructs.IConstruct
	// A reference to a IdentityPoolRoleAttachment resource.
	// Experimental.
	IdentityPoolRoleAttachmentRef() *IdentityPoolRoleAttachmentReference
}

// The jsii proxy for IIdentityPoolRoleAttachmentRef
type jsiiProxy_IIdentityPoolRoleAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIdentityPoolRoleAttachmentRef) IdentityPoolRoleAttachmentRef() *IdentityPoolRoleAttachmentReference {
	var returns *IdentityPoolRoleAttachmentReference
	_jsii_.Get(
		j,
		"identityPoolRoleAttachmentRef",
		&returns,
	)
	return returns
}

