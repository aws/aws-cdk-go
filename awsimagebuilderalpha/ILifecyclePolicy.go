package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Lifecycle Policy.
// Experimental.
type ILifecyclePolicy interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the lifecycle policy.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions to the given grantee for the lifecycle policy.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the lifecycle policy.
	// Experimental.
	LifecyclePolicyArn() *string
	// The name of the lifecycle policy.
	// Experimental.
	LifecyclePolicyName() *string
}

// The jsii proxy for ILifecyclePolicy
type jsiiProxy_ILifecyclePolicy struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ILifecyclePolicy) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILifecyclePolicy) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILifecyclePolicy) LifecyclePolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecyclePolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecyclePolicy) LifecyclePolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecyclePolicyName",
		&returns,
	)
	return returns
}

