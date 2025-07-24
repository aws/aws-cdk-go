package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A construct that represents an IAM principal, such as a user, group or role.
type IIdentity interface {
	IPrincipal
	awscdk.IResource
	// Attaches a managed policy to this principal.
	AddManagedPolicy(policy IManagedPolicy)
	// Attaches an inline policy to this principal.
	//
	// This is the same as calling `policy.addToXxx(principal)`.
	AttachInlinePolicy(policy Policy)
}

// The jsii proxy for IIdentity
type jsiiProxy_IIdentity struct {
	jsiiProxy_IPrincipal
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IIdentity) AddManagedPolicy(policy IManagedPolicy) {
	if err := i.validateAddManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addManagedPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IIdentity) AttachInlinePolicy(policy Policy) {
	if err := i.validateAttachInlinePolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"attachInlinePolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IIdentity) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := i.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		i,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIdentity) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IIdentity) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentity) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentity) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentity) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentity) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentity) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

