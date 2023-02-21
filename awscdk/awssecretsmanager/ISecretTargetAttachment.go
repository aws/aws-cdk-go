package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISecretTargetAttachment interface {
	ISecret
	// Same as `secretArn`.
	SecretTargetAttachmentSecretArn() *string
}

// The jsii proxy for ISecretTargetAttachment
type jsiiProxy_ISecretTargetAttachment struct {
	jsiiProxy_ISecret
}

func (j *jsiiProxy_ISecretTargetAttachment) SecretTargetAttachmentSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretTargetAttachmentSecretArn",
		&returns,
	)
	return returns
}

