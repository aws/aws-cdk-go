package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM Instance Profile.
type IInstanceProfile interface {
	IInstanceProfileRef
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
	jsiiProxy_IInstanceProfileRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInstanceProfile) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IInstanceProfile) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceProfile) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

