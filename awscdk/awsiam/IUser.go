package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM user.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html
//
type IUser interface {
	IIdentity
	interfacesawsiam.IUserRef
	// Adds this user to a group.
	AddToGroup(group IGroup)
	// The user's ARN.
	UserArn() *string
	// The user's name.
	UserName() *string
}

// The jsii proxy for IUser
type jsiiProxy_IUser struct {
	jsiiProxy_IIdentity
	internal.Type__interfacesawsiamIUserRef
}

func (i *jsiiProxy_IUser) AddToGroup(group IGroup) {
	if err := i.validateAddToGroupParameters(group); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToGroup",
		[]interface{}{group},
	)
}

func (i *jsiiProxy_IUser) AddManagedPolicy(policy IManagedPolicy) {
	if err := i.validateAddManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addManagedPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IUser) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
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

func (i *jsiiProxy_IUser) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IUser) AttachInlinePolicy(policy Policy) {
	if err := i.validateAttachInlinePolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"attachInlinePolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IUser) UserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUser) UserRef() *interfacesawsiam.UserReference {
	var returns *interfacesawsiam.UserReference
	_jsii_.Get(
		j,
		"userRef",
		&returns,
	)
	return returns
}

