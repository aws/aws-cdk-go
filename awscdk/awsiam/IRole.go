package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Role object.
type IRole interface {
	IIdentity
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

