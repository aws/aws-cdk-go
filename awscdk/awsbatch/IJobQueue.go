package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a JobQueue.
type IJobQueue interface {
	interfacesawsbatch.IJobQueueRef
	awscdk.IResource
	// Add a `ComputeEnvironment` to this Queue.
	//
	// The Queue will prefer lower-order `ComputeEnvironment`s.
	AddComputeEnvironment(computeEnvironment interfacesawsbatch.IComputeEnvironmentRef, order *float64)
	// The set of compute environments mapped to a job queue and their order relative to each other.
	//
	// The job scheduler uses this parameter to determine which compute environment runs a specific job.
	// Compute environments must be in the VALID state before you can associate them with a job queue.
	// You can associate up to three compute environments with a job queue.
	// All of the compute environments must be either EC2 (EC2 or SPOT) or Fargate (FARGATE or FARGATE_SPOT);
	// EC2 and Fargate compute environments can't be mixed.
	//
	// *Note*: All compute environments that are associated with a job queue must share the same architecture.
	// AWS Batch doesn't support mixing compute environment architecture types in a single job queue.
	ComputeEnvironments() *[]*OrderedComputeEnvironment
	// If the job queue is enabled, it is able to accept jobs.
	//
	// Otherwise, new jobs can't be added to the queue, but jobs already in the queue can finish.
	// Default: true.
	//
	Enabled() *bool
	// The ARN of this job queue.
	JobQueueArn() *string
	// The name of the job queue.
	//
	// It can be up to 128 letters long.
	// It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_).
	JobQueueName() *string
	// The priority of the job queue.
	//
	// Job queues with a higher priority are evaluated first when associated with the same compute environment.
	// Priority is determined in descending order.
	// For example, a job queue with a priority value of 10 is given scheduling preference over a job queue with a priority value of 1.
	Priority() *float64
	// The SchedulingPolicy for this JobQueue.
	//
	// Instructs the Scheduler how to schedule different jobs.
	// Default: - no scheduling policy.
	//
	SchedulingPolicy() ISchedulingPolicy
}

// The jsii proxy for IJobQueue
type jsiiProxy_IJobQueue struct {
	internal.Type__interfacesawsbatchIJobQueueRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IJobQueue) AddComputeEnvironment(computeEnvironment interfacesawsbatch.IComputeEnvironmentRef, order *float64) {
	if err := i.validateAddComputeEnvironmentParameters(computeEnvironment, order); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addComputeEnvironment",
		[]interface{}{computeEnvironment, order},
	)
}

func (i *jsiiProxy_IJobQueue) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IJobQueue) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IJobQueue) ComputeEnvironments() *[]*OrderedComputeEnvironment {
	var returns *[]*OrderedComputeEnvironment
	_jsii_.Get(
		j,
		"computeEnvironments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) JobQueueArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) JobQueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobQueueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) SchedulingPolicy() ISchedulingPolicy {
	var returns ISchedulingPolicy
	_jsii_.Get(
		j,
		"schedulingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) JobQueueRef() *interfacesawsbatch.JobQueueReference {
	var returns *interfacesawsbatch.JobQueueReference
	_jsii_.Get(
		j,
		"jobQueueRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueue) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

