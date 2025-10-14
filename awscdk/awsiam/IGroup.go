package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM Group.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups.html
//
type IGroup interface {
	IGroupRef
	IIdentity
	// Returns the IAM Group ARN.
	GroupArn() *string
	// Returns the IAM Group Name.
	GroupName() *string
}

// The jsii proxy for IGroup
type jsiiProxy_IGroup struct {
	jsiiProxy_IGroupRef
	jsiiProxy_IIdentity
}

func (i *jsiiProxy_IGroup) AddManagedPolicy(policy IManagedPolicy) {
	if err := i.validateAddManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addManagedPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IGroup) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
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

func (i *jsiiProxy_IGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IGroup) AttachInlinePolicy(policy Policy) {
	if err := i.validateAttachInlinePolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"attachInlinePolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IGroup) GroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) GroupRef() *GroupReference {
	var returns *GroupReference
	_jsii_.Get(
		j,
		"groupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

