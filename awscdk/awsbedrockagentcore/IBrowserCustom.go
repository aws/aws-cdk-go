package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Browser resources.
type IBrowserCustom interface {
	interfacesawsbedrockagentcore.IBrowserCustomRef
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants `Get` and `List` actions on the Browser.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants `Invoke`, `Start`, and `Update` actions on the Browser.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this browser.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of errors for a specific API operation performed on this browser.
	MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the given named metric related to the API operation performed on this browser.
	MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of API requests made for a specific browser operation.
	MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the latency of a specific API operation performed on this browser.
	MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the duration of browser sessions.
	MetricSessionDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of system errors for a specific API operation performed on this browser.
	MetricSystemErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user takeovers.
	MetricTakeOverCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the duration of user takeovers.
	MetricTakeOverDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user takeovers released.
	MetricTakeOverReleaseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of throttled requests for a specific API operation performed on this browser.
	MetricThrottlesForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user errors for a specific API operation performed on this browser.
	MetricUserErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the browser resource.
	BrowserArn() *string
	// The id of the browser.
	BrowserId() *string
	// Timestamp when the browser was created.
	CreatedAt() *string
	// The IAM role that provides permissions for the browser to access AWS services.
	ExecutionRole() awsiam.IRole
	// Timestamp when the browser was last updated.
	LastUpdatedAt() *string
	// The status of the browser.
	Status() *string
}

// The jsii proxy for IBrowserCustom
type jsiiProxy_IBrowserCustom struct {
	internal.Type__interfacesawsbedrockagentcoreIBrowserCustomRef
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBrowserCustom) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
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

func (i *jsiiProxy_IBrowserCustom) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) GrantUse(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantUseParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantUse",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IBrowserCustom) MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricErrorsForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricErrorsForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricForApiOperationParameters(metricName, operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricForApiOperation",
		[]interface{}{metricName, operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocationsForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLatencyForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatencyForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricSessionDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSessionDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSessionDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricSystemErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSystemErrorsForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSystemErrorsForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricTakeOverCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTakeOverCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTakeOverCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricTakeOverDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTakeOverDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTakeOverDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricTakeOverReleaseCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTakeOverReleaseCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTakeOverReleaseCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricThrottlesForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottlesForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottlesForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) MetricUserErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricUserErrorsForApiOperationParameters(operation, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricUserErrorsForApiOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBrowserCustom) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IBrowserCustom) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBrowserCustom) BrowserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) BrowserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) LastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) BrowserCustomRef() *interfacesawsbedrockagentcore.BrowserCustomReference {
	var returns *interfacesawsbedrockagentcore.BrowserCustomReference
	_jsii_.Get(
		j,
		"browserCustomRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBrowserCustom) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

