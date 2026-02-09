package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for CodeInterpreterCustom resources.
// Experimental.
type ICodeInterpreterCustom interface {
	interfacesawsbedrockagentcore.ICodeInterpreterCustomRef
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grants `Get` and `List` actions on the Code Interpreter.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants `Invoke`, `Start`, and `Stop` actions on the Code Interpreter.
	// Experimental.
	GrantUse(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this code interpreter.
	// Experimental.
	Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of errors for a specific API operation performed on this code interpreter.
	// Experimental.
	MetricErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the given named metric related to the API operation performed on this code interpreter.
	// Experimental.
	MetricForApiOperation(metricName *string, operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the total number of API requests made for a specific code interpreter operation.
	// Experimental.
	MetricInvocationsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the latency of a specific API operation performed on this code interpreter.
	// Experimental.
	MetricLatencyForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric measuring the duration of code interpreter sessions.
	// Experimental.
	MetricSessionDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of system errors for a specific API operation performed on this code interpreter.
	// Experimental.
	MetricSystemErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of throttled requests for a specific API operation performed on this code interpreter.
	// Experimental.
	MetricThrottlesForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return a metric containing the number of user errors for a specific API operation performed on this code interpreter.
	// Experimental.
	MetricUserErrorsForApiOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the code interpreter resource.
	// Experimental.
	CodeInterpreterArn() *string
	// The id of the code interpreter.
	// Experimental.
	CodeInterpreterId() *string
	// Timestamp when the code interpreter was created.
	// Experimental.
	CreatedAt() *string
	// The IAM role that provides permissions for the code interpreter to access AWS services.
	// Experimental.
	ExecutionRole() awsiam.IRole
	// Timestamp when the code interpreter was last updated.
	// Experimental.
	LastUpdatedAt() *string
	// The status of the code interpreter.
	// Experimental.
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

func (i *jsiiProxy_ICodeInterpreterCustom) Metric(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

