package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an IAM Policy.
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage.html
//
type IPolicy interface {
	IPolicyRef
	awscdk.IResource
	// The name of this policy.
	PolicyName() *string
}

// The jsii proxy for IPolicy
type jsiiProxy_IPolicy struct {
	jsiiProxy_IPolicyRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) PolicyRef() *PolicyReference {
	var returns *PolicyReference
	_jsii_.Get(
		j,
		"policyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

