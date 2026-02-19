package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Role object.
type IRole interface {
	IIdentity
	interfacesawsiam.IRoleRef
	// Grant the actions defined in actions to the identity Principal on this resource.
	Grant(grantee IPrincipal, actions ...*string) Grant
	// Grant permissions to the given principal to assume this role.
	GrantAssumeRole(grantee IPrincipal) Grant
	// Grant permissions to the given principal to pass this role.
	GrantPassRole(grantee IPrincipal) Grant
	// Returns the ARN of this role.
	RoleArn() *string
	// Returns the name of this role.
	RoleName() *string
}

// The jsii proxy for IRole
type jsiiProxy_IRole struct {
	jsiiProxy_IIdentity
	internal.Type__interfacesawsiamIRoleRef
}

func (i *jsiiProxy_IRole) Grant(grantee IPrincipal, actions ...*string) Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRole) GrantAssumeRole(grantee IPrincipal) Grant {
	if err := i.validateGrantAssumeRoleParameters(grantee); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		i,
		"grantAssumeRole",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRole) GrantPassRole(grantee IPrincipal) Grant {
	if err := i.validateGrantPassRoleParameters(grantee); err != nil {
		panic(err)
	}
	var returns Grant

	_jsii_.Invoke(
		i,
		"grantPassRole",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRole) AddManagedPolicy(policy IManagedPolicy) {
	if err := i.validateAddManagedPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addManagedPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRole) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
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

func (i *jsiiProxy_IRole) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRole) AttachInlinePolicy(policy Policy) {
	if err := i.validateAttachInlinePolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"attachInlinePolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRole) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRole) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) RoleRef() *interfacesawsiam.RoleReference {
	var returns *interfacesawsiam.RoleReference
	_jsii_.Get(
		j,
		"roleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRole) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

