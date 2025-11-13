package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// A managed policy.
type IManagedPolicy interface {
	interfacesawsiam.IManagedPolicyRef
	// The ARN of the managed policy.
	ManagedPolicyArn() *string
}

// The jsii proxy for IManagedPolicy
type jsiiProxy_IManagedPolicy struct {
	internal.Type__interfacesawsiamIManagedPolicyRef
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

