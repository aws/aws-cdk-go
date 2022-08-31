package awskinesisanalyticsflink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisanalyticsflink/internal"
)

// An interface expressing the public properties on both an imported and CDK-created Flink application.
// Experimental.
type IApplication interface {
	awsiam.IGrantable
	awscdk.IResource
	// Convenience method for adding a policy statement to the application role.
	// Experimental.
	AddToRolePolicy(policyStatement awsiam.PolicyStatement) *bool
	// Return a CloudWatch metric associated with this Flink application.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time (in milliseconds) this task or operator is back pressured per second.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Operator, Task, Parallelism.
	// Experimental.
	MetricBackPressuredTimeMsPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time (in milliseconds) this task or operator is busy (neither idle nor back pressured) per second.
	//
	// Can be NaN, if the value could not be
	// calculated.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Operator, Task, Parallelism.
	// Experimental.
	MetricBusyTimePerMsPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The overall percentage of CPU utilization across task managers.
	//
	// For
	// example, if there are five task managers, Kinesis Data Analytics publishes
	// five samples of this metric per reporting interval.
	//
	// Units: Percentage
	//
	// Reporting Level: Application.
	// Experimental.
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The last watermark this application/operator/task/thread has received.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricCurrentInputWatermark(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The last watermark this application/operator/task/thread has received.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricCurrentOutputWatermark(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time elapsed during an outage for failing/recovering jobs.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Application.
	// Experimental.
	MetricDowntime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of times this job has fully restarted since it was submitted.
	//
	// This metric does not measure fine-grained restarts.
	//
	// Units: Count
	//
	// Reporting Level: Application.
	// Experimental.
	MetricFullRestarts(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Overall heap memory utilization across task managers.
	//
	// For example, if there
	// are five task managers, Kinesis Data Analytics publishes five samples of
	// this metric per reporting interval.
	//
	// Units: Percentage
	//
	// Reporting Level: Application.
	// Experimental.
	MetricHeapMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time (in milliseconds) this task or operator is idle (has no data to process) per second.
	//
	// Idle time excludes back pressured time, so if the task
	// is back pressured it is not idle.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Operator, Task, Parallelism.
	// Experimental.
	MetricIdleTimeMsPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of Kinesis Processing Units that are used to run your stream processing application.
	//
	// The average number of KPUs used each hour
	// determines the billing for your application.
	//
	// Units: Count
	//
	// Reporting Level: Application.
	// Experimental.
	MetricKpus(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time it took to complete the last checkpoint.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Application.
	// Experimental.
	MetricLastCheckpointDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total size of the last checkpoint.
	//
	// Units: Bytes
	//
	// Reporting Level: Application.
	// Experimental.
	MetricLastCheckpointSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of managed memory.
	//
	// Units: Bytes
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricManagedMemoryTotal(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of managed memory currently used.
	//
	// Units: Bytes
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricManagedMemoryUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Derived from managedMemoryUsed/managedMemoryTotal.
	//
	// Units: Percentage
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricManagedMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of times checkpointing has failed.
	//
	// Units: Count
	//
	// Reporting Level: Application.
	// Experimental.
	MetricNumberOfFailedCheckpoints(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records this operator or task has dropped due to arriving late.
	//
	// Units: Count
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricNumLateRecordsDropped(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of records this application, operator, or task has received.
	//
	// Units: Count
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricNumRecordsIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of records this application, operator or task has received per second.
	//
	// Units: Count/Second
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricNumRecordsInPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of records this application, operator or task has emitted.
	//
	// Units: Count
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricNumRecordsOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of records this application, operator or task has emitted per second.
	//
	// Units: Count/Second
	//
	// Reporting Level: Application, Operator, Task, Parallelism.
	// Experimental.
	MetricNumRecordsOutPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of old garbage collection operations that have occurred across all task managers.
	//
	// Units: Count
	//
	// Reporting Level: Application.
	// Experimental.
	MetricOldGenerationGCCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total time spent performing old garbage collection operations.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Application.
	// Experimental.
	MetricOldGenerationGCTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of live threads used by the application.
	//
	// Units: Count
	//
	// Reporting Level: Application.
	// Experimental.
	MetricThreadsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time that the job has been running without interruption.
	//
	// Units: Milliseconds
	//
	// Reporting Level: Application.
	// Experimental.
	MetricUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The application ARN.
	// Experimental.
	ApplicationArn() *string
	// The name of the Flink application.
	// Experimental.
	ApplicationName() *string
	// The application IAM role.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IApplication
type jsiiProxy_IApplication struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplication) AddToRolePolicy(policyStatement awsiam.PolicyStatement) *bool {
	if err := i.validateAddToRolePolicyParameters(policyStatement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"addToRolePolicy",
		[]interface{}{policyStatement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricBackPressuredTimeMsPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBackPressuredTimeMsPerSecondParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBackPressuredTimeMsPerSecond",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricBusyTimePerMsPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBusyTimePerMsPerSecondParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBusyTimePerMsPerSecond",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricCurrentInputWatermark(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCurrentInputWatermarkParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCurrentInputWatermark",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricCurrentOutputWatermark(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCurrentOutputWatermarkParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCurrentOutputWatermark",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricDowntime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDowntimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDowntime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricFullRestarts(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFullRestartsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFullRestarts",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricHeapMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricHeapMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricHeapMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricIdleTimeMsPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIdleTimeMsPerSecondParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIdleTimeMsPerSecond",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricKpus(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricKpusParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricKpus",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricLastCheckpointDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLastCheckpointDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLastCheckpointDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricLastCheckpointSize(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLastCheckpointSizeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLastCheckpointSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricManagedMemoryTotal(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricManagedMemoryTotalParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricManagedMemoryTotal",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricManagedMemoryUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricManagedMemoryUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricManagedMemoryUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricManagedMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricManagedMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricManagedMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricNumberOfFailedCheckpoints(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumberOfFailedCheckpointsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumberOfFailedCheckpoints",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricNumLateRecordsDropped(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumLateRecordsDroppedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumLateRecordsDropped",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricNumRecordsIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumRecordsInParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumRecordsIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricNumRecordsInPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumRecordsInPerSecondParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumRecordsInPerSecond",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricNumRecordsOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumRecordsOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumRecordsOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricNumRecordsOutPerSecond(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumRecordsOutPerSecondParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumRecordsOutPerSecond",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricOldGenerationGCCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricOldGenerationGCCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricOldGenerationGCCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricOldGenerationGCTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricOldGenerationGCTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricOldGenerationGCTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricThreadsCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThreadsCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThreadsCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) MetricUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricUptimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

