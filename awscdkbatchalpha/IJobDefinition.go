// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkbatchalpha/v2/internal"
)

// Represents a JobDefinition.
// Experimental.
type IJobDefinition interface {
	awscdk.IResource
	// Add a RetryStrategy to this JobDefinition.
	// Experimental.
	AddRetryStrategy(strategy RetryStrategy)
	// The ARN of this job definition.
	// Experimental.
	JobDefinitionArn() *string
	// The name of this job definition.
	// Experimental.
	JobDefinitionName() *string
	// The default parameters passed to the container These parameters can be referenced in the `command` that you give to the container.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html#parameters
	//
	// Experimental.
	Parameters() *map[string]interface{}
	// The number of times to retry a job.
	//
	// The job is retried on failure the same number of attempts as the value.
	// Experimental.
	RetryAttempts() *float64
	// Defines the retry behavior for this job.
	// Experimental.
	RetryStrategies() *[]RetryStrategy
	// The priority of this Job.
	//
	// Only used in Fairshare Scheduling
	// to decide which job to run first when there are multiple jobs
	// with the same share identifier.
	// Experimental.
	SchedulingPriority() *float64
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes,
	// Batch terminates your jobs if they aren't finished.
	// Experimental.
	Timeout() awscdk.Duration
}

// The jsii proxy for IJobDefinition
type jsiiProxy_IJobDefinition struct {
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

