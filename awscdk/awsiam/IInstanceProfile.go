package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// Represents an IAM Instance Profile.
type IInstanceProfile interface {
	awscdk.IResource
	// The InstanceProfile's ARN.
	InstanceProfileArn() *string
	// The InstanceProfile's name.
	InstanceProfileName() *string
	// The role associated with the InstanceProfile.
	Role() IRole
}

// The jsii proxy for IInstanceProfile
type jsiiProxy_IInstanceProfile struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IInstanceProfile) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceProfile) InstanceProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceProfile) Role() IRole {
	var returns IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

