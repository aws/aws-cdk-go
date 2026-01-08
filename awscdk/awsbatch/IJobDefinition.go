package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a JobDefinition.
type IJobDefinition interface {
	interfacesawsbatch.IJobDefinitionRef
	awscdk.IResource
	// Add a RetryStrategy to this JobDefinition.
	AddRetryStrategy(strategy RetryStrategy)
	// The ARN of this job definition.
	JobDefinitionArn() *string
	// The name of this job definition.
	JobDefinitionName() *string
	// The default parameters passed to the container These parameters can be referenced in the `command` that you give to the container.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html#parameters
	//
	// Default: none.
	//
	Parameters() *map[string]interface{}
	// The number of times to retry a job.
	//
	// The job is retried on failure the same number of attempts as the value.
	// Default: 1.
	//
	RetryAttempts() *float64
	// Defines the retry behavior for this job.
	// Default: - no `RetryStrategy`.
	//
	RetryStrategies() *[]RetryStrategy
	// The priority of this Job.
	//
	// Only used in Fairshare Scheduling
	// to decide which job to run first when there are multiple jobs
	// with the same share identifier.
	// Default: none.
	//
	SchedulingPriority() *float64
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes,
	// Batch terminates your jobs if they aren't finished.
	// Default: - no timeout.
	//
	Timeout() awscdk.Duration
}

// The jsii proxy for IJobDefinition
type jsiiProxy_IJobDefinition struct {
	internal.Type__interfacesawsbatchIJobDefinitionRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IJobDefinition) AddRetryStrategy(strategy RetryStrategy) {
	if err := i.validateAddRetryStrategyParameters(strategy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addRetryStrategy",
		[]interface{}{strategy},
	)
}

func (i *jsiiProxy_IJobDefinition) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) RetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) RetryStrategies() *[]RetryStrategy {
	var returns *[]RetryStrategy
	_jsii_.Get(
		j,
		"retryStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) SchedulingPriority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) JobDefinitionRef() *interfacesawsbatch.JobDefinitionReference {
	var returns *interfacesawsbatch.JobDefinitionReference
	_jsii_.Get(
		j,
		"jobDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobDefinition) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

