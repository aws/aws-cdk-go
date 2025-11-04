package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyPrincipalAttachment.
// Experimental.
type IPolicyPrincipalAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PolicyPrincipalAttachment resource.
	// Experimental.
	PolicyPrincipalAttachmentRef() *PolicyPrincipalAttachmentReference
}

// The jsii proxy for IPolicyPrincipalAttachmentRef
type jsiiProxy_IPolicyPrincipalAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPolicyPrincipalAttachmentRef) PolicyPrincipalAttachmentRef() *PolicyPrincipalAttachmentReference {
	var returns *PolicyPrincipalAttachmentReference
	_jsii_.Get(
		j,
		"policyPrincipalAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyPrincipalAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyPrincipalAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

