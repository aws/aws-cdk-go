package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions/internal"
)

// A State Machine.
// Experimental.
type IStateMachine interface {
	awsiam.IGrantable
	awscdk.IResource
	// Grant the given identity custom permissions.
	// Experimental.
	Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions for all executions of a state machine.
	// Experimental.
	GrantExecution(identity awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity read permissions for this state machine.
	// Experimental.
	GrantRead(identity awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to start an execution of this state machine.
	// Experimental.
	GrantStartExecution(identity awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to start a synchronous execution of this state machine.
	// Experimental.
	GrantStartSyncExecution(identity awsiam.IGrantable) awsiam.Grant
	// Grant the given identity read permissions for this state machine.
	// Experimental.
	GrantTaskResponse(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this State Machine's executions.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that were aborted.
	// Experimental.
	MetricAborted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that failed.
	// Experimental.
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that were started.
	// Experimental.
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that succeeded.
	// Experimental.
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that were throttled.
	// Experimental.
	MetricThrottled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the interval, in milliseconds, between the time the execution starts and the time it closes.
	// Experimental.
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of executions that timed out.
	// Experimental.
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the state machine.
	// Experimental.
	StateMachineArn() *string
}

// The jsii proxy for IStateMachine
type jsiiProxy_IStateMachine struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IStateMachine) Grant(identity awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (j *jsiiProxy_IStateMachine) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

