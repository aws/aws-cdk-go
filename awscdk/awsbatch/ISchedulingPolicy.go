package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Scheduling Policy.
//
// Scheduling Policies tell the Batch
// Job Scheduler how to schedule incoming jobs.
type ISchedulingPolicy interface {
	awscdk.IResource
	interfacesawsbatch.ISchedulingPolicyRef
	// The arn of this scheduling policy.
	SchedulingPolicyArn() *string
	// The name of this scheduling policy.
	SchedulingPolicyName() *string
}

// The jsii proxy for ISchedulingPolicy
type jsiiProxy_ISchedulingPolicy struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsbatchISchedulingPolicyRef
}

func (i *jsiiProxy_ISchedulingPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_ISchedulingPolicy) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedulingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedulingPolicy) SchedulingPolicyRef() *interfacesawsbatch.SchedulingPolicyReference {
	var returns *interfacesawsbatch.SchedulingPolicyReference
	_jsii_.Get(
		j,
		"schedulingPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedulingPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

