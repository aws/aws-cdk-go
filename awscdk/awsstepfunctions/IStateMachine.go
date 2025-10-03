package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A State Machine.
type IStateMachine interface {
	awsiam.IGrantable
	awscdk.IResource
	// Grant the given identity custom permissions.
	Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions for all executions of a state machine.
	GrantExecution(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity read permissions for this state machine.
	GrantRead(identity awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to start an execution of this state machine.
	GrantStartExecution(identity awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to start a synchronous execution of this state machine.
	GrantStartSyncExecution(identity awsiam.IGrantable) awsiam.Grant
	// Grant the given identity read permissions for this state machine.
	GrantTaskResponse(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this State Machine's executions.
	// Default: - sum over 5 minutes.
	//
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that were aborted.
	// Default: - sum over 5 minutes.
	//
	MetricAborted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that failed.
	// Default: - sum over 5 minutes.
	//
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that were started.
	// Default: - sum over 5 minutes.
	//
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that succeeded.
	// Default: - sum over 5 minutes.
	//
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that were throttled.
	// Default: sum over 5 minutes.
	//
	MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the interval, in milliseconds, between the time the execution starts and the time it closes.
	// Default: - sum over 5 minutes.
	//
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that timed out.
	// Default: - sum over 5 minutes.
	//
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the state machine.
	StateMachineArn() *string
}

// The jsii proxy for IStateMachine
type jsiiProxy_IStateMachine struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IStateMachine) Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(identity); err != nil {
		panic(err)
	}
	args := []interface{}{identity}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) GrantExecution(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantExecutionParameters(identity); err != nil {
		panic(err)
	}
	args := []interface{}{identity}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantExecution",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) GrantRead(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) GrantStartExecution(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStartExecutionParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStartExecution",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) GrantStartSyncExecution(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStartSyncExecutionParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStartSyncExecution",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) GrantTaskResponse(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantTaskResponseParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantTaskResponse",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IStateMachine) MetricAborted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricAbortedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricAborted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStateMachine) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IStateMachine) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachine) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachine) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateMachine) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

