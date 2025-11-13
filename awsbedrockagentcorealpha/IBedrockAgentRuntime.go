package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Agent Runtime resources.
// Experimental.
type IBedrockAgentRuntime interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Adds a policy statement to the runtime's execution role.
	//
	// Returns: The runtime instance for chaining.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement) IBedrockAgentRuntime
	// Grant the runtime specific actions on AWS resources.
	//
	// Returns: The Grant object for chaining.
	// Experimental.
	Grant(actions *[]*string, resources *[]*string) awsiam.Grant
	// Permits an IAM principal to invoke this runtime both directly and on behalf of users Grants both bedrock-agentcore:InvokeAgentRuntime and bedrock-agentcore:InvokeAgentRuntimeForUser permissions.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal to invoke this runtime Grants the bedrock-agentcore:InvokeAgentRuntime permission.
	// Experimental.
	GrantInvokeRuntime(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal to invoke this runtime on behalf of a user Grants the bedrock-agentcore:InvokeAgentRuntimeForUser permission Required when using the X-Amzn-Bedrock-AgentCore-Runtime-User-Id header.
	// Experimental.
	GrantInvokeRuntimeForUser(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this agent runtime.
	// Experimental.
	Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of invocations for this agent runtime.
	// Experimental.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of invocations across all resources.
	// Experimental.
	MetricInvocationsAggregated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the latency of requests for this agent runtime.
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of agent sessions for this agent runtime.
	// Experimental.
	MetricSessionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of sessions across all resources.
	// Experimental.
	MetricSessionsAggregated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of system errors for this agent runtime.
	// Experimental.
	MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of throttled requests for this agent runtime.
	// Experimental.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of errors (system + user) for this agent runtime.
	// Experimental.
	MetricTotalErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user errors for this agent runtime.
	// Experimental.
	MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the agent runtime resource - Format `arn:${Partition}:bedrock-agentcore:${Region}:${Account}:runtime/${RuntimeId}`.
	//
	// Example:
	//   "arn:aws:bedrock-agentcore:us-west-2:123456789012:runtime/runtime-abc123"
	//
	// Experimental.
	AgentRuntimeArn() *string
	// The ID of the agent runtime.
	//
	// Example:
	//   "runtime-abc123"
	//
	// Experimental.
	AgentRuntimeId() *string
	// The name of the agent runtime.
	// Experimental.
	AgentRuntimeName() *string
	// The version of the agent runtime.
	// Experimental.
	AgentRuntimeVersion() *string
	// The current status of the agent runtime.
	// Experimental.
	AgentStatus() *string
	// The time at which the runtime was created.
	//
	// Example:
	//   "2024-01-15T10:30:00Z"
	//
	// Experimental.
	CreatedAt() *string
	// The time at which the runtime was last updated.
	//
	// Example:
	//   "2024-01-15T14:45:00Z"
	//
	// Experimental.
	LastUpdatedAt() *string
	// The IAM role that provides permissions for the agent runtime.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IBedrockAgentRuntime
type jsiiProxy_IBedrockAgentRuntime struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBedrockAgentRuntime) AddToRolePolicy(statement awsiam.PolicyStatement) IBedrockAgentRuntime {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns IBedrockAgentRuntime

	_jsii_.Invoke(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) Grant(actions *[]*string, resources *[]*string) awsiam.Grant {
	if err := i.validateGrantParameters(actions, resources); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		[]interface{}{actions, resources},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) GrantInvokeRuntime(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeRuntimeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeRuntime",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) GrantInvokeRuntimeForUser(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeRuntimeForUserParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeRuntimeForUser",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, dimensions, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, dimensions, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricInvocationsAggregated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsAggregatedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationsAggregated",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricSessionCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSessionCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSessionCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricSessionsAggregated(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSessionsAggregatedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSessionsAggregated",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSystemErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSystemErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricTotalErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricUserErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricUserErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBedrockAgentRuntime) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IBedrockAgentRuntime) AgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) AgentRuntimeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) AgentRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) AgentRuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) AgentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) LastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockAgentRuntime) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

