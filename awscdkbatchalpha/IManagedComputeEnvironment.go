package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Managed ComputeEnvironment.
//
// Batch will provision EC2 Instances to
// meet the requirements of the jobs executing in this ComputeEnvironment.
// Experimental.
type IManagedComputeEnvironment interface {
	IComputeEnvironment
	awsec2.IConnectable
	// The maximum vCpus this `ManagedComputeEnvironment` can scale up to.
	//
	// *Note*: if this Compute Environment uses EC2 resources (not Fargate) with either `AllocationStrategy.BEST_FIT_PROGRESSIVE` or
	// `AllocationStrategy.SPOT_CAPACITY_OPTIMIZED`, or `AllocationStrategy.BEST_FIT` with Spot instances,
	// The scheduler may exceed this number by at most one of the instances specified in `instanceTypes`
	// or `instanceClasses`.
	// Experimental.
	MaxvCpus() *float64
	// Specifies whether this Compute Environment is replaced if an update is made that requires replacing its instances.
	//
	// To enable more properties to be updated,
	// set this property to `false`. When changing the value of this property to false,
	// do not change any other properties at the same time.
	// If other properties are changed at the same time,
	// and the change needs to be rolled back but it can't,
	// it's possible for the stack to go into the UPDATE_ROLLBACK_FAILED state.
	// You can't update a stack that is in the UPDATE_ROLLBACK_FAILED state.
	// However, if you can continue to roll it back,
	// you can return the stack to its original settings and then try to update it again.
	//
	// The properties which require a replacement of the Compute Environment are:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html
	//
	// Experimental.
	ReplaceComputeEnvironment() *bool
	// The security groups this Compute Environment will launch instances in.
	// Experimental.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// Whether or not to use spot instances.
	//
	// Spot instances are less expensive EC2 instances that can be
	// reclaimed by EC2 at any time; your job will be given two minutes
	// of notice before reclamation.
	// Experimental.
	Spot() *bool
	// Whether or not any running jobs will be immediately terminated when an infrastructure update occurs.
	//
	// If this is enabled, any terminated jobs may be retried, depending on the job's
	// retry policy.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
	//
	// Experimental.
	TerminateOnUpdate() *bool
	// Only meaningful if `terminateOnUpdate` is `false`.
	//
	// If so,
	// when an infrastructure update is triggered, any running jobs
	// will be allowed to run until `updateTimeout` has expired.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html
	//
	// Experimental.
	UpdateTimeout() awscdk.Duration
	// Whether or not the AMI is updated to the latest one supported by Batch when an infrastructure update occurs.
	//
	// If you specify a specific AMI, this property will be ignored.
	// Experimental.
	UpdateToLatestImageVersion() *bool
	// The VPC Subnets this Compute Environment will launch instances in.
	// Experimental.
	VpcSubnets() *awsec2.SubnetSelection
}

// The jsii proxy for IManagedComputeEnvironment
type jsiiProxy_IManagedComputeEnvironment struct {
	jsiiProxy_IComputeEnvironment
	internal.Type__awsec2IConnectable
}

func (i *jsiiProxy_IManagedComputeEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IManagedComputeEnvironment) MaxvCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxvCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) ReplaceComputeEnvironment() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"replaceComputeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) Spot() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) TerminateOnUpdate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminateOnUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) UpdateTimeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"updateTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) UpdateToLatestImageVersion() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"updateToLatestImageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) VpcSubnets() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"vpcSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) ComputeEnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IManagedComputeEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

