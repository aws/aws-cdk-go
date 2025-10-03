package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecretTargetAttachment.
// Experimental.
type ISecretTargetAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISecretTargetAttachmentRef
type jsiiProxy_ISecretTargetAttachmentRef struct {
	internal.Type__constructsIConstruct
}

