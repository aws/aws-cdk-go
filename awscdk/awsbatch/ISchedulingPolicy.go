package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
)

// Represents a Scheduling Policy.
//
// Scheduling Policies tell the Batch
// Job Scheduler how to schedule incoming jobs.
type ISchedulingPolicy interface {
	awscdk.IResource
	// The arn of this scheduling policy.
	SchedulingPolicyArn() *string
	// The name of this scheduling policy.
	SchedulingPolicyName() *string
}

// The jsii proxy for ISchedulingPolicy
type jsiiProxy_ISchedulingPolicy struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISchedulingPolicy) SchedulingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedulingPolicy) SchedulingPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyName",
		&returns,
	)
	return returns
}

