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

// Interface for CodeInterpreterCustom resources.
type ICodeInterpreterCustom interface {
	interfacesawsbedrockagentcore.ICodeInterpreterCustomRef
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants `Get` and `List` actions on the Code Interpreter.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants `Invoke`, `Start`, and `Stop` actions on the Code Interpreter.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this code interpreter.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of errors for a specific API operation performed on this code interpreter.
	MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the given named metric related to the API operation performed on this code interpreter.
	MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of API requests made for a specific code interpreter operation.
	MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the latency of a specific API operation performed on this code interpreter.
	MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the duration of code interpreter sessions.
	MetricSessionDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of system errors for a specific API operation performed on this code interpreter.
	MetricSystemErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of throttled requests for a specific API operation performed on this code interpreter.
	MetricThrottlesForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user errors for a specific API operation performed on this code interpreter.
	MetricUserErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the code interpreter resource.
	CodeInterpreterArn() *string
	// The id of the code interpreter.
	CodeInterpreterId() *string
	// Timestamp when the code interpreter was created.
	CreatedAt() *string
	// The IAM role that provides permissions for the code interpreter to access AWS services.
	ExecutionRole() awsiam.IRole
	// Timestamp when the code interpreter was last updated.
	LastUpdatedAt() *string
	// The status of the code interpreter.
	Status() *string
}

// The jsii proxy for ICodeInterpreterCustom
type jsiiProxy_ICodeInterpreterCustom struct {
	internal.Type__interfacesawsbedrockagentcoreICodeInterpreterCustomRef
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICodeInterpreterCustom) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_ICodeInterpreterCustom) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_ICodeInterpreterCustom) GrantUse(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_ICodeInterpreterCustom) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricSessionDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricSystemErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricThrottlesForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) MetricUserErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ICodeInterpreterCustom) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ICodeInterpreterCustom) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICodeInterpreterCustom) CodeInterpreterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInterpreterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) CodeInterpreterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInterpreterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) ExecutionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) LastUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) CodeInterpreterCustomRef() *interfacesawsbedrockagentcore.CodeInterpreterCustomReference {
	var returns *interfacesawsbedrockagentcore.CodeInterpreterCustomReference
	_jsii_.Get(
		j,
		"codeInterpreterCustomRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeInterpreterCustom) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

