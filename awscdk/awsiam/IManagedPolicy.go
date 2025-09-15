package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A managed policy.
type IManagedPolicy interface {
	IManagedPolicyRef
	// The ARN of the managed policy.
	ManagedPolicyArn() *string
}

// The jsii proxy for IManagedPolicy
type jsiiProxy_IManagedPolicy struct {
	jsiiProxy_IManagedPolicyRef
}

func (j *jsiiProxy_IManagedPolicy) ManagedPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArn",
		&returns,
	)
	return returns
}

