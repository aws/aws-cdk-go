package awscdkelasticachealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/internal"
)

// Represents an ElastiCache base user.
// Experimental.
type IUser interface {
	awscdk.IResource
	// The engine for the user.
	// Experimental.
	Engine() UserEngine
	// The user's ARN.
	// Experimental.
	UserArn() *string
	// The user's ID.
	// Experimental.
	UserId() *string
	// The user's name.
	// Experimental.
	UserName() *string
}

// The jsii proxy for IUser
type jsiiProxy_IUser struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IUser) Engine() UserEngine {
	var returns UserEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_IUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
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

