package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an IAM user.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html
//
type IUser interface {
	IIdentity
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

