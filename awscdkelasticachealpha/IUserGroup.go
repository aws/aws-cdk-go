package awscdkelasticachealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/internal"
)

// Represents an ElastiCache UserGroup.
// Experimental.
type IUserGroup interface {
	awscdk.IResource
	// Add a user to this user group.
	// Experimental.
	AddUser(user IUser)
	// The engine type for the user group.
	// Experimental.
	Engine() UserEngine
	// The ARN of the user group.
	// Experimental.
	UserGroupArn() *string
	// The name of the user group.
	// Experimental.
	UserGroupName() *string
	// List of users in the user group.
	// Experimental.
	Users() *[]IUser
}

// The jsii proxy for IUserGroup
type jsiiProxy_IUserGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IUserGroup) AddUser(user IUser) {
	if err := i.validateAddUserParameters(user); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addUser",
		[]interface{}{user},
	)
}

func (j *jsiiProxy_IUserGroup) Engine() UserEngine {
	var returns UserEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserGroup) UserGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserGroup) UserGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserGroup) Users() *[]IUser {
	var returns *[]IUser
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

