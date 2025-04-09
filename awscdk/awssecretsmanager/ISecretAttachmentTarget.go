package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A secret attachment target.
type ISecretAttachmentTarget interface {
	// Renders the target specifications.
	AsSecretAttachmentTarget() *SecretAttachmentTargetProps
}

// The jsii proxy for ISecretAttachmentTarget
type jsiiProxy_ISecretAttachmentTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_ISecretAttachmentTarget) AsSecretAttachmentTarget() *SecretAttachmentTargetProps {
	var returns *SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

