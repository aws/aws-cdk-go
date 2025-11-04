package awssecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SecretTargetAttachment.
// Experimental.
type ISecretTargetAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SecretTargetAttachment resource.
	// Experimental.
	SecretTargetAttachmentRef() *SecretTargetAttachmentReference
}

// The jsii proxy for ISecretTargetAttachmentRef
type jsiiProxy_ISecretTargetAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISecretTargetAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISecretTargetAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

