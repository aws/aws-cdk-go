package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolUserToGroupAttachment.
// Experimental.
type IUserPoolUserToGroupAttachmentRef interface {
	constructs.IConstruct
	// A reference to a UserPoolUserToGroupAttachment resource.
	// Experimental.
	UserPoolUserToGroupAttachmentRef() *UserPoolUserToGroupAttachmentReference
}

// The jsii proxy for IUserPoolUserToGroupAttachmentRef
type jsiiProxy_IUserPoolUserToGroupAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserPoolUserToGroupAttachmentRef) UserPoolUserToGroupAttachmentRef() *UserPoolUserToGroupAttachmentReference {
	var returns *UserPoolUserToGroupAttachmentReference
	_jsii_.Get(
		j,
		"userPoolUserToGroupAttachmentRef",
		&returns,
	)
	return returns
}

