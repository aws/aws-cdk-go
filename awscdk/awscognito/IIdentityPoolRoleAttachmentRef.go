package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPoolRoleAttachment.
// Experimental.
type IIdentityPoolRoleAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdentityPoolRoleAttachment resource.
	// Experimental.
	IdentityPoolRoleAttachmentRef() *IdentityPoolRoleAttachmentReference
}

// The jsii proxy for IIdentityPoolRoleAttachmentRef
type jsiiProxy_IIdentityPoolRoleAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IIdentityPoolRoleAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolRoleAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

