package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a ComputeEnvironment.
type IComputeEnvironment interface {
	interfacesawsbatch.IComputeEnvironmentRef
	awscdk.IResource
	// The ARN of this compute environment.
	ComputeEnvironmentArn() *string
	// The name of the ComputeEnvironment.
	ComputeEnvironmentName() *string
	// Whether or not this ComputeEnvironment can accept jobs from a Queue.
	//
	// Enabled ComputeEnvironments can accept jobs from a Queue and
	// can scale instances up or down.
	// Disabled ComputeEnvironments cannot accept jobs from a Queue or
	// scale instances up or down.
	//
	// If you change a ComputeEnvironment from enabled to disabled while it is executing jobs,
	// Jobs in the `STARTED` or `RUNNING` states will not
	// be interrupted. As jobs complete, the ComputeEnvironment will scale instances down to `minvCpus`.
	//
	// To ensure you aren't billed for unused capacity, set `minvCpus` to `0`.
	Enabled() *bool
	// The role Batch uses to perform actions on your behalf in your account, such as provision instances to run your jobs.
	// Default: - a serviceRole will be created for managed CEs, none for unmanaged CEs.
	//
	ServiceRole() awsiam.IRole
}

// The jsii proxy for IComputeEnvironment
type jsiiProxy_IComputeEnvironment struct {
	internal.Type__interfacesawsbatchIComputeEnvironmentRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IComputeEnvironment) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) ComputeEnvironmentRef() *interfacesawsbatch.ComputeEnvironmentReference {
	var returns *interfacesawsbatch.ComputeEnvironmentReference
	_jsii_.Get(
		j,
		"computeEnvironmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComputeEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

