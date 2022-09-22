package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an IAM Group.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups.html
//
// Experimental.
type IGroup interface {
	IIdentity
	// Returns the IAM Group ARN.
	// Experimental.
	GroupArn() *string
	// Returns the IAM Group Name.
	// Experimental.
	GroupName() *string
}

// The jsii proxy for IGroup
type jsiiProxy_IGroup struct {
	jsiiProxy_IIdentity
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

