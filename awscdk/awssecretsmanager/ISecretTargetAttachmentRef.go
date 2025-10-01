package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecretTargetAttachment.
// Experimental.
type ISecretTargetAttachmentRef interface {
	constructs.IConstruct
	// A reference to a SecretTargetAttachment resource.
	// Experimental.
	SecretTargetAttachmentRef() *SecretTargetAttachmentReference
}

// The jsii proxy for ISecretTargetAttachmentRef
type jsiiProxy_ISecretTargetAttachmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISecretTargetAttachmentRef) SecretTargetAttachmentRef() *SecretTargetAttachmentReference {
	var returns *SecretTargetAttachmentReference
	_jsii_.Get(
		j,
		"secretTargetAttachmentRef",
		&returns,
	)
	return returns
}

