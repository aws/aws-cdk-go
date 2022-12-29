package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A managed policy.
type IManagedPolicy interface {
	// The ARN of the managed policy.
	ManagedPolicyArn() *string
}

// The jsii proxy for IManagedPolicy
type jsiiProxy_IManagedPolicy struct {
	_ byte // padding
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

